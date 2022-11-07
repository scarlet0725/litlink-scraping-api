package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/scarlet0725/prism-api/cache"
	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/scraping"
	"github.com/scarlet0725/prism-api/serializer"
)

var supportedSites = map[string]string{
	"t.livepocket.jp": "livepocket",
	"lit.link":        "litlink",
}

func main() {

	port, flg := os.LookupEnv("PORT")
	if !flg {
		port = "8080"
	}

	cacheAddr, flg := os.LookupEnv("CACHE_ADDR")

	if !flg {
		cacheAddr = "localhost:6379"
	}

	reidsConfig := &redis.Options{
		Addr:     cacheAddr,
		Password: "",
		DB:       0,
	}

	redisClient := redis.NewClient(reidsConfig)

	cache := cache.CreateRedisManager(redisClient)

	sc := scraping.CreateClient()
	sl := serializer.CreateSerializer()

	c := controller.CreateContoroller(supportedSites, &sc, &sl, &cache)

	http.HandleFunc("/scraiping", c.ScrapingRequestHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
