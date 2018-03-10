package interfaces

type Cacher interface {
    Get(key string) string
    Set(key, value string) bool
    IsExists(key string) bool
    Delete(key string) bool
}
