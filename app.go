package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis"
	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure"
	"go.uber.org/zap"
)

func main() {
	var (
		mg = flag.Bool("migration", false, "migration")
	)

	flag.Parse()

	zapLogger, err := zap.NewProduction(zap.AddCallerSkip(1))

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	logger := framework.NewLogger(zapLogger)

	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbConf := cmd.DBConfig{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	}

	ent, err := cmd.InitDB(dbConf)

	if err != nil {
		log.Fatal(err)
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

	gin, err := infrastructure.NewGinRouter(logger, ent, redisClient)

	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if *mg {
		err := cmd.Migrate(ctx, ent)

		if err != nil {
			log.Fatal(err)
		}

		return
	}

	err = gin.Serve(serverAddr)

	log.Fatal(err)

}
