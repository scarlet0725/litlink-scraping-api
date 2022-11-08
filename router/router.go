package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Router interface {
	Respond(http.ResponseWriter, *http.Request, any) error
	RespondError(http.ResponseWriter, *http.Request, int, string) error
	ScrapingRequestHandler(w http.ResponseWriter, r *http.Request)
	HealthCheckHandler(http.ResponseWriter, *http.Request)
	DefaultHandler(http.ResponseWriter, *http.Request)
}

type router struct {
}

func NewRouter() Router {
	return &router{}
}

func (r *router) ScrapingRequestHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (r *router) HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"ok\": true}"))
}

func (r *router) DefaultHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	m := "{\"message\": \"This is prism-api\", \"github\": \"https://github.com/scarlet0725/prism-api\"}"

	w.Write([]byte(m))
}

func (r *router) Respond(w http.ResponseWriter, req *http.Request, v interface{}) error {
	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func (r *router) RespondError(w http.ResponseWriter, req *http.Request, code int, msg string) error {
	w.WriteHeader(code)
	_, err := w.Write([]byte(fmt.Sprintf("{\"ok\": false, \"error\": \"%s\"}", msg)))
	return err
}
