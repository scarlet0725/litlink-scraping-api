package main

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/gateway"
	"github.com/scarlet0725/prism-api/parser"
	"github.com/scarlet0725/prism-api/router"
	"github.com/scarlet0725/prism-api/selializer"
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

	cache := gateway.NewRedisManager(redisClient)
	httpClient := gateway.NewHTTPClient()
	fetchController := controller.NewFetchController(httpClient, cache)

	parser := parser.NewParser()
	serializer := selializer.NewResponseSerializer()

	gin := router.NewGinRouter(fetchController, parser, serializer)

	err := gin.Serve(serverAddr)

	log.Fatal(err)
}
