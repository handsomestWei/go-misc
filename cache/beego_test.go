package cache

import (
	"testing"
	"time"
)

func TestNewBeegoCache(t *testing.T) {
	cache, err := NewBeegoCache(10)
	if err != nil {
		t.Error(err)
	}

	key := "testKey"
	val := "testVal"
	cache.Put(key, val, 2*time.Second)
	if cache.Get(key) != val {
		t.Fail()
	}
	time.Sleep(3 * time.Second)
	if cache.IsExist(key) {
		t.Fail()
	}
}
