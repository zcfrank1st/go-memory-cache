package memcache

import (
    "testing"
    "time"
    "fmt"
)

func TestGoMemCache_Get_Set(t *testing.T) {
    goCache := &GoMemCache{cache:make(map[string]string)}

    goCache.Set("hello", "world")

    if v, ok := goCache.Get("hello"); ok {
        if v != "world" {
            t.Fatal("value not the same")
        }
    } else {
        t.Fatal("get value wrong")
    }
}

func TestGoMemCache_SetWithTTL(t *testing.T) {
    goCache := &GoMemCache{cache:make(map[string]string)}


    goCache.SetWithTTL("hello", "world", time.Second * 10)

    v, ok := goCache.Get("hello")
    fmt.Println(v, ok)

    if ok {
        if v != "world" {
            t.Fatal("value not the same")
        }
    } else {
        t.Fatal("get value wrong")
    }

    time.Sleep(time.Second * 15)

    v1, ok1 := goCache.Get("hello")

    fmt.Println(v1, ok1)
    if !ok1 {
        if v1 != "" {
            t.Fatal("ttl not work")
        }
    }
}