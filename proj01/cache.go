package main

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pmylund/go-cache"
	"log"
	"os"
	"strings"
	"time"
)

var mc *memcache.Client
var lc *cache.Cache

var cacher CacheManager = CacheManager{
	Hosts:  os.Getenv("MEMCACHE_HOSTS"),
	Engine: "mc",
}

type CacheManager struct {
	Hosts  string
	Engine string
}

func (m *CacheManager) InitCache() {
	if m.Hosts == "" {
		m.Engine = "lc"
	}

	switch m.Engine {
	case "mc":
		hosts := strings.Split(m.Hosts, ",")
		mc = memcache.New(hosts...)
	case "lc":
		lc = cache.New(-1, 24*time.Hour)
	}
}

func (m *CacheManager) Set(key, value string) {
	switch m.Engine {
	case "lc":
		lcSet(key, value)
	case "mc":
		mcSet(key, value)
	}
}

func (m *CacheManager) Get(key string) string {
	switch m.Engine {
	case "mc":
		return mcGet(key)
	case "lc":
		return lcGet(key)
	}

	return ""
}

func (m *CacheManager) Delete(key string) {
	switch m.Engine {
	case "mc":
		mc.Delete(key)
	case "lc":
		lc.Delete(key)
	}
}

func mcSet(key, value string) {
	if err := mc.Add(&memcache.Item{Key: key, Value: []byte(value)}); err != nil {
		log.Println(err.Error())
	}
}

func mcGet(key string) string {
	item, err := mc.Get(key)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return string(item.Value)
}

func lcSet(key, value string) {
	if err := lc.Add(key, value, 24*time.Hour); err != nil {
		log.Println(err.Error())
	}
}

func lcGet(key string) string {
	item, found := lc.Get(key)

	if !found {
		log.Println("LC Cache miss")
		return ""
	}

	return item.(string)
}
