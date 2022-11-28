package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	_, err := http.Get(fmt.Sprintf("http://localhost:%s/health", port))

	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
