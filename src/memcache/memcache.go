package memcache

import (
    "time"
)

type GoMemCache struct {
    cache map[string]string
}

type CacheOption interface {
    Set(key string, value string)
    Get(key string) (string, bool)
    SetWithTTL(key string, value string, ttl time.Duration)
}

func (m *GoMemCache) Set(key string, value string) {
    m.cache[key] = value
}

func (m *GoMemCache) Get(key string) (v string, ok bool) {
    v, ok = m.cache[key]
    return
}

func (m *GoMemCache) SetWithTTL(key string, value string, ttl time.Duration) {
    m.cache[key] = value
    go cacheCleaner(m, key, ttl)
}

func cacheCleaner(m *GoMemCache, key string, ttl time.Duration) {
    timer := time.NewTimer(ttl)
    go func() {
        <- timer.C
        delete(m.cache, key)
    }()
}


