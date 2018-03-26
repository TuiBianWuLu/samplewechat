package cache

import "time"

type Cacher interface {
    Get(key string) (string, error)
    Set(key string, value interface{}, ttl time.Duration) (bool, error)
    IsExists(key string) (bool, error)
    Del(key string) (bool, error)
}
