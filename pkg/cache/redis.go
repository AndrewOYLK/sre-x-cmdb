package cache

type RedisCache struct{}

func (r *RedisCache) Save(cv CacheValue) error {
	return nil
}

func (r *RedisCache) Query(key string) string {
	return ""
}
