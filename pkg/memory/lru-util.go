package memory

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"os"
	"strconv"
)

var LruCache *lru.Cache

func init() {
	lruCacheSize := 8
	CacheSize := os.Getenv("CACHE_SIZE")
	if CacheSize != "" {
		fmt.Println("CACHE-SIZE:", CacheSize)
		lruCacheNum, err := strconv.Atoi(CacheSize)
		if err == nil {
			lruCacheSize = lruCacheNum
		}
	}
	LruCache, _ = lru.New(lruCacheSize)
}
