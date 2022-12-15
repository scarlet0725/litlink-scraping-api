package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure"
	"go.uber.org/zap"
)

func main() {

	zap, err := zap.NewProduction()
	logger := framework.NewLogger(zap)

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

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

	db, err := cmd.ConnectDB(dsn)

	if err != nil {
		log.Fatal(err)
	}

	if *mg {
		cmd.MigrationDB(db)
		return
	}

	serverAddr := cmd.ConfigureHTTPServer()
	cacheAddr := cmd.ConfigureCacheServer()

	redisPassword := cmd.GetRedisPassword()

	reidsConfig := &redis.Options{
		Addr:     cacheAddr,
		Password: redisPassword,
		DB:       0,
	}

	redisClient := redis.NewClient(reidsConfig)

	gin, err := infrastructure.NewGinRouter(logger, db, redisClient)

	if err != nil {
		log.Fatal(err)
	}

	err = gin.Serve(serverAddr)

	log.Fatal(err)

}
