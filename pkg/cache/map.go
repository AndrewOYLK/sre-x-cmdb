package cache

import "sync"

var mc *MapCache

type MapCache struct {
	data map[string]string
}

func (r *MapCache) Save(cv CacheValue) error {
	r.data[cv.Key] = cv.Value
	return nil
}

func (r *MapCache) Query(key string) string {
	return r.data[key]
}

func NewMapCache() Cache {
	once := sync.Once{}
	once.Do(func() {
		if mc == nil {
			mc = &MapCache{}
		}
	})
	return mc
}
