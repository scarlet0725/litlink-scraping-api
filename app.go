package main

import (
	"log"
	"net/http"
	"os"

	"github.com/scarlet0725/litlink-scraping-api/controller"
	"github.com/scarlet0725/litlink-scraping-api/scraping"
	"github.com/scarlet0725/litlink-scraping-api/serializer"
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

	sc := scraping.CreateClient()
	sl := serializer.CreateSerializer()

	c := controller.CreateContoroller(supportedSites, &sc, &sl)

	http.HandleFunc("/scraiping", c.ScrapingRequestHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
