package gocache

import (
	lru "github.com/hashicorp/golang-lru"
	"sync"
)

type Store interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	Close() error
}

type StoreCache struct {
	store Store
	cache *lru.Cache
	lock  sync.RWMutex
}

func NewCache(store Store, cacheSize int) (*StoreCache, error) {
	r := &StoreCache{}
	r.cache, _ = lru.New(cacheSize)
	r.store = store
	return r, nil
}

func (s *StoreCache) Set(key string, value string) error {
	old, exist := s.cache.Get(key)
	s.cache.Add(key, value)
	s.lock.Lock()
	err := s.store.Set(key, value)
	s.lock.Unlock()
	if err != nil && exist {
		s.cache.Add(key, old)
	}
	return err
}

func (s *StoreCache) Get(key string) (string, error) {
	// local cache get
	c, exist := s.cache.Get(key)
	if exist {
		return c.(string), nil
	}
	// redis cache get
	s.lock.Lock()
	v, err := s.store.Get(key)
	s.lock.Unlock()
	if v != "" {
		// grab data from redis and write into local cache
		s.cache.Add(key, v)
	}
	return v, err
}

func (s *StoreCache) Delete(key string) error {
	old, exist := s.cache.Get(key)

	s.cache.Remove(key)
	s.lock.Lock()
	err := s.store.Delete(key)
	s.lock.Unlock()

	if err != nil && exist {
		s.cache.Add(key, old)
	}
	return err
}

func (s *StoreCache) Close() error {
	return s.store.Close()
}