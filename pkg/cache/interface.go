package cache

// Cache 缓存接口
type Cache interface {
	Save(CacheValue) error
	Query(string) string
}

// Cache 缓存数据
type CacheValue struct {
	Key   string
	Value string
}
