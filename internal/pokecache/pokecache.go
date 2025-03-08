package pokecache

import(
	"time"
	"sync"
)

type Cache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	_cache := Cache{
		cache:make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}

	// call reaploop in new job
	go _cache.reapLoop(interval)

	return _cache
}

func (_cache *Cache) Add(key string, value []byte) {
	_cache.mux.Lock()
	defer _cache.mux.Unlock()
	_cache.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: value,
	}
}

func (_cache *Cache) Get(key string) ([]byte, bool) {
	_cache.mux.Lock()
	defer _cache.mux.Unlock()
	val, ok := _cache.cache[key]
	return val.val, ok
}


func (_cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		_cache.reap(time.Now().UTC(), interval)
	}

}

func (_cache *Cache) reap(now time.Time, last time.Duration) {

	_cache.mux.Lock()
	defer _cache.mux.Unlock()
	for k, v := range _cache.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(_cache.cache, k)
		}
	}
}