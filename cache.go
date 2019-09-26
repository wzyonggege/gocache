package gocache

import lru "github.com/hashicorp/golang-lru"

type Store interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	Close() error
}

type StoreCache struct {
	store Store
	cache *lru.Cache
}

func (s *StoreCache) Set(key string, value string) error {
	panic("i")
}

func (s *StoreCache) Get(key string) (string, error) {
	panic("i")
}

func (s *StoreCache) Delete(key string) error {
	panic("1")
}

func (s *StoreCache) Close() error {
	panic("1")
}