package main

import (
	"container/list"
	"sync"
	"time"
)

const (
	EVICT_ENTRIES_PERCENT int = 5 // Evict 5% entries
)

var (
	proxyCache *Cache
	Wrk        []Worker
)

//
// Cache Element
//
type Element struct {
	key        string        // Key
	value      string        // Value type can be further generalized by using interface
	accessTime time.Time     // Last time the cache entry was accessed (read/write)
	listElm    *list.Element // Pointer to element in LRU list
}

func (elm Element) IsValidEntry(tSec time.Duration) bool {

	if time.Now().Sub(elm.accessTime) > tSec {
		return false
	}
	return true
}

//
// Cache container
//
type Cache struct {
	sync.RWMutex
	maxCapacity  int                // Max number of entries in cache
	expiryFactor int                // Fraction of LRU cache to clear during eviction
	expiryPeriod time.Duration      // TTL for each cache entry
	data         map[string]Element // Map for key-value lookup
	lruList      *list.List         // Double link list maintaining least recently used order
}

func NewCache(capacity int, expPeriod time.Duration) *Cache {

	return &Cache{
		maxCapacity:  capacity,
		expiryFactor: EVICT_ENTRIES_PERCENT,
		expiryPeriod: expPeriod,
		data:         make(map[string]Element, capacity),
		lruList:      list.New(),
	}
}

func (c *Cache) getCapacity() int {
	return len(c.data)
}

func (c *Cache) IsFull() bool {
	if c.maxCapacity-len(c.data) <= 0 {
		return true
	}
	return false
}

func (c *Cache) LruEvict() {

	c.Lock()

	// Number of entries at one time
	entriesToEvict := (c.maxCapacity * c.expiryFactor) / 100

	for i := 0; i < entriesToEvict; i++ {

		// Get the least recently used element
		back := c.lruList.Back()
		if back == nil {
			break
		}

		// Remove last entry
		elm := c.lruList.Remove(back).(Element)
		delete(c.data, elm.key)
	}

	c.Unlock()
}

func (c *Cache) remove(key string) {

	c.Lock()

	elm, exist := c.data[key]
	if exist {
		c.lruList.Remove(elm.listElm)
		delete(c.data, key)
	}
	c.Unlock()
}

func (c *Cache) set(key string, value string) {

	if c.IsFull() {
		// Make room for new entries
		c.LruEvict()
	}

	c.Lock()

	elm, exist := c.data[key]
	if exist {

		// Move the entry to front
		c.lruList.MoveToFront(elm.listElm)
		elm.accessTime = time.Now()

	} else {

		// Create new entry and insert in front
		elm := Element{
			key:        key,
			value:      value,
			accessTime: time.Now(),
		}

		elm.listElm = c.lruList.PushFront(elm)
		c.data[key] = elm
	}

	c.Unlock()
}

func (c *Cache) lookup(key string) (string, bool, bool) {

	elm, exist := c.data[key]
	if exist != true {
		return "", false, false
	}

	if elm.IsValidEntry(c.expiryPeriod) != true {
		return "", false, true
	}

	// Update access time
	elm.accessTime = time.Now()
	return elm.value, true, true
}

func (c *Cache) getValidElement(key string) (string, bool) {

	c.RLock()
	value, valid, exist := c.lookup(key)
	c.RUnlock()

	// Entry not found
	if exist != true {
		return "", exist
	}

	// Expired cached entry
	if valid != true {
		//log.Println("Expiring Cache entry: ", key)
		c.remove(key)
		return "", valid
	}

	// Found a valid cache entry
	return value, true
}

//
// Get data from Redis Proxy
//
func (c *Cache) Get(key string) string {

	// Lookup data in proxy cache
	value, exist := c.getValidElement(key)
	if exist {
		return value
	}

	// Element not found in cache, lookup backend
	data, _ := RedisGet(key)
	if data == "" {
		return ""
	}

	// Found data in backend, save in cache and return
	c.set(key, data)

	return data
}

//
// Workers to support concurrent processing of http requests
//
type Worker struct {
	jobQueue chan Job
	quit     chan bool
}

func NewWorker(queue chan Job) Worker {

	return Worker{
		jobQueue: queue,
		quit:     make(chan bool),
	}

}

func (w *Worker) Start() {

	for {
		select {
		case job := <-w.jobQueue:
			value := proxyCache.Get(job.k)
			SendGetRsp(job, value)

		case <-w.quit:
			return
		}
	}
}

func (w *Worker) Stop() {
	w.quit <- true
}

func StartLruCacheHandlers(conf *Config, jobQueue chan Job) {

	proxyCache = NewCache(conf.cacheCapacity, conf.cacheExpiry)

	Wrk = make([]Worker, conf.parallelReqCnt)

	// Number of job that are being concurrently procesed
	for i := 0; i < conf.parallelReqCnt; i++ {
		Wrk[i] = NewWorker(jobQueue)
		go Wrk[i].Start()
	}

}

func StopLruCacheHandlers() {
	for i := 0; i < len(Wrk); i++ {
		Wrk[i].Stop()
	}
}
