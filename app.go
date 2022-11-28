package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/gateway"
	"github.com/scarlet0725/prism-api/infrastructure"
	"github.com/scarlet0725/prism-api/parser"
	"github.com/scarlet0725/prism-api/selializer"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	var (
		mg = flag.Bool("migration", false, "migration")
	)

	flag.Parse()

	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.Debug()

	if err != nil {
		log.Fatal(err)
	}

	if *mg {
		cmd.MigrationDB(db)
		return
	}

	orm := infrastructure.NewGORMClient(db)

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

	gin := infrastructure.NewGinRouter(fetchController, parser, serializer, orm)

	err = gin.Serve(serverAddr)

	log.Fatal(err)

}
