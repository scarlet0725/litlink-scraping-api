package main

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/scarlet0725/prism-api/cache"
	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/gateway"
	"github.com/scarlet0725/prism-api/parser"
	"github.com/scarlet0725/prism-api/router"
	"github.com/scarlet0725/prism-api/scraping"
)

var supportedSites = map[string]string{
	"t.livepocket.jp": "livepocket",
	"lit.link":        "litlink",
}

func main() {

	serverAddr := cmd.ConfigureHTTPServer()
	cacheAddr := cmd.ConfigureCacheServer()

	redisPassword := cmd.GetRedisPassword()

	reidsConfig := &redis.Options{
		Addr:     cacheAddr,
		Password: redisPassword,
		DB:       0,
	}

	redisClient := redis.NewClient(reidsConfig)

	cache := cache.CreateRedisManager(redisClient)

	sc := scraping.CreateClient()
	sl := parser.CreateSerializer()

	rt := router.NewRouter()
	c := controller.CreateContoroller(supportedSites, &sc, &sl, &cache)

	server := gateway.InitAPIServer(serverAddr)
	server.AddRoute("/scraping", c.ScrapingRequestHandler)
	server.AddRoute("/health", rt.HealthCheckHandler)
	server.AddRoute("/", rt.DefaultHandler)

	err := server.Serve()

	log.Fatal(err)
}
