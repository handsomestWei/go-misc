package cache

import (
	"github.com/astaxie/beego/cache"
	"fmt"
)

func NewBeegoCache(interval int) (adapter cache.Cache, err error) {
	return cache.NewCache("memory", fmt.Sprintf(`{"interval":%s}`, string(interval)))
}