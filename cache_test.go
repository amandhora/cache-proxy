package main

import (
	"container/list"
	"strconv"
	"testing"
)

func TestLruEvict(t *testing.T) {

	nums := []string{"zero", "one", "two", "three", "four"}
	capacity := len(nums)

	c := &Cache{
		maxCapacity:  capacity,
		expiryFactor: 40, // Expire 40%
		expiryPeriod: DEFAULT_CACHE_ENTRY_TTL,
		data:         make(map[string]Element, capacity),
		lruList:      list.New(),
	}

	for idx, v := range nums {
		c.set(strconv.Itoa(idx), v)
	}
	if c.getCapacity() != 5 {
		t.Errorf("Cache size not as expected : %d", c.getCapacity())
	}

	// Test Evict
	c.set("5", "five")
	if c.getCapacity() != 4 {
		t.Errorf("Cache size not as expected : %d", c.getCapacity())
	}

	c.set("6", "six")
	if c.getCapacity() != 5 {
		t.Errorf("Cache size not as expected : %d", c.getCapacity())
	}

	c.set("7", "seven")
	if c.getCapacity() != 4 {
		t.Errorf("Cache size not as expected : %d", c.getCapacity())
	}
}

func TestGetSetRemove(t *testing.T) {

	nums := []string{"zero", "one", "two", "three", "four"}
	capacity := len(nums)

	c := &Cache{
		maxCapacity:  capacity,
		expiryFactor: 40, // Expire 40%
		expiryPeriod: DEFAULT_CACHE_ENTRY_TTL,
		data:         make(map[string]Element, capacity),
		lruList:      list.New(),
	}

	// Set entries in cache
	for idx, v := range nums {
		c.set(strconv.Itoa(idx), v)
	}
	if c.getCapacity() != 5 {
		t.Errorf("Cache size not as expected : %d", c.getCapacity())
	}

	// Read the entries back from cache
	for idx, v := range nums {
		key := strconv.Itoa(idx)
		value, exist := c.getValidElement(key)
		if !exist {
			t.Errorf("Cache value not found")
		}
		if v != value {
			t.Errorf("Cache value not as expected : %-6s , found: %-6s", v, value)
		}
	}

	c.remove("1")
	if c.getCapacity() != 4 {
		t.Errorf("Cache size not as expected : %d", c.getCapacity())
	}

	c.remove("unknown")
	if c.getCapacity() != 4 {
		t.Errorf("Cache size not as expected : %d", c.getCapacity())
	}
}
