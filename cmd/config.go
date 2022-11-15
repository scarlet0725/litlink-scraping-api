package cmd

import (
	"os"
)

func ConfigureHTTPServer() string {
	port, flg := os.LookupEnv("PORT")
	if !flg {
		port = "8080"
	}

	addr := ":"

	conStr := addr + port

	return conStr
}

func ConfigureCacheServer() string {
	port, flg := os.LookupEnv("REDIS_PORT")
	if !flg {
		port = "6379"
	}

	addr, flg := os.LookupEnv("REDIS_ADDR")

	if !flg {
		addr = "localhost"
	}

	conStr := addr + ":" + port

	return conStr
}

func GetRedisPassword() string {
	secret, flg := os.LookupEnv("REDIS_PASSWORD")

	if !flg {
		secret = ""
	}

	return secret
}

func MigrateDB() {

}
