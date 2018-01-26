# Cache Proxy
HTTP proxy with LRU cache

The redis-proxy service contains 3 parts

### LRU Cache 

For building an LRU cache, I thought of two data structures

First, we'll need a hash-table. Get (and set) should be O(1).

Next, we'll need a doubly linked list - O(n).  This is for implementing the eviction policy. Whenever we get or set an item from the cache, we'll put that same object to the front of our list. During eviction, we can trim the tail of our list and remove it from our hash table.

These read and write operation are done by using a read-write mutex for concurrency. LRU cache runs multiple workers to handle concurrent requests.

### HTTP server 

The get url is expected to be:
```console
/proxy?key=KEY
```
The HTTP request handler reads the request and adds it to backend queue (of size equal to max Connection)


### Redis and Config Handlers

Redis pool is configured to streamline the redis access. Several options are read from env variables and stored in config container.
```console
        environment:
            REDIS_URL: redis:6379
            PROXY_PORT: 8080
            CACHE_CAP: 1000
            CACHE_TTL_SEC: 600

```


## INSTRUCTION

### Prerequisite

The software build and runs on a modern Linux distribution or Mac OS installation, 

The system should have the following software installed:
```console
make
docker
docker-compose
Bash
```

### Steps
```console
# git clone https://github.com/amandhora/cache-proxy.git
# cd cache-proxy
# make test
```
