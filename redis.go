package gocache

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type RedisStore struct {
	pool *redis.Client
}

func NewRedisStore() (Store, error) {
	pool := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "", // no password set
		DB:       0,       // use default DB
	})
	r := &RedisStore{
		pool: pool,
	}
	return r, nil
}

func (s *RedisStore) Set(key string, value string) error {
	val := s.pool.Set(key, value, time.Second * 1000)
	if err := val.Err(); err != nil {
		return err
	}
	return nil
}

func (s *RedisStore) Get(key string) (string, error) {
	val := s.pool.Get(key)
	if err := val.Err(); err != nil {
		return "", err
	}
	return val.String(), nil
}

func (s *RedisStore) Delete(key string) error {
	val := s.pool.Del(key)
	if err := val.Err(); err != nil {
		return err
	}
	return nil
}

func (s *RedisStore) Close() error {
	return s.pool.Close()
}