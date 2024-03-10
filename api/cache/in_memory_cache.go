package cache

import (
	"crypto/sha256"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

var AppCache *cache.Cache

var AppCacheExpiration = 45 * time.Second

func NewCache() {
	AppCache = cache.New(AppCacheExpiration, 5*time.Minute)
}

func GenerateKey(metricName string, d time.Duration) string {
	inputData := metricName + d.String()

	hash := sha256.Sum256([]byte(inputData))

	return fmt.Sprintf("%x", hash)
}
