package memcache

import "sync"

type CacheImpl struct {
	sync.Map
}

func New() *CacheImpl {
	return &CacheImpl{
		Map: sync.Map{},
	}
}

func (c *CacheImpl) Set(key string, val any) {
	c.Store(key, val)
}

func (c *CacheImpl) Get(key string) (any, bool) {
	return c.Load(key)
}
