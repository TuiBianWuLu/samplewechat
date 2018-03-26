package cache

import (
    "fmt"
    "time"

    "github.com/go-redis/redis"
)

type Redis struct {
    Client *redis.Client
}

var r = new(Redis)

func NewRedis(options *redis.Options) *Redis {
    r.Client = redis.NewClient(options)
    return r
}

func (r *Redis) Get(key string) (string, error) {
    return r.Client.Get(key).Result()
}

func (r *Redis) Set(key string, value interface{}, ttl time.Duration) (result bool, err error) {
    res, err := r.Client.Set(key, value, ttl).Result()
    if res == "OK" {
        result = true
    }
    return
}

func (r *Redis) IsExists(key string) (result bool, err error) {
    res, err := r.Client.Exists(key).Result()
    if res == 1 {
        result = true
    }
    return
}

func (r *Redis) Del(key string) (result bool, err error) {
    res, err := r.Client.Del(key).Result()
    if res == 1 {
        result = true
    }
    return
}
