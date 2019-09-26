package gocache

import "github.com/garyburd/redigo/redis"

type RdsStore struct {
	rds *redis.Pool
}

func NewRdsStore() (Store, error) {
	r := &RdsStore{}
	return r, nil
}

func (s *RdsStore) Set(key string, value string) error {
	panic("i")
}

func (s *RdsStore) Get(key string) (string, error) {
	panic("i")
}

func (s *RdsStore) Delete(key string) error {
	panic("1")
}

func (s *RdsStore) Close() error {
	panic("1")
}