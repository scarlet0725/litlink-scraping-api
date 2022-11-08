package router

import (
	"net/http"
)

type Router interface {
	ScrapingRequestHandler(http.ResponseWriter, *http.Request)
	HealthCheckHandler(http.ResponseWriter, *http.Request)
}

type router struct {
}

func InitRouter() Router {
	return &router{}
}

func (r *router) ScrapingRequestHandler(w http.ResponseWriter, req *http.Request) {
}

func (r *router) HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"ok\": true}"))
}
