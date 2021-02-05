package middleware

import (
	"github.com/docker/libkv/store/zookeeper"
	"github.com/docker/libkv/store"
)

func NewZkSimpleClient(urls []string) (store.Store, error) {
	return zookeeper.New(urls, nil)
}