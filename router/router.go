package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/usecase"
)

type Router interface {
	Respond(http.ResponseWriter, *http.Request, any) error
	RespondError(http.ResponseWriter, *http.Request, int, string) error
	ScrapingRequestHandler(w http.ResponseWriter, r *http.Request)
	HealthCheckHandler(http.ResponseWriter, *http.Request)
	DefaultHandler(http.ResponseWriter, *http.Request)
}

type router struct {
	Scraping usecase.ScrapingApplication
}

func NewRouter(s *usecase.ScrapingApplication) Router {
	return &router{
		Scraping: *s,
	}
}

func (r *router) ScrapingRequestHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	hasRequiredParams := req.URL.Query().Has("url")

	if !hasRequiredParams {
		r.RespondError(w, req, http.StatusBadRequest, "url is required")
		return
	}

	u := req.URL.Query().Get("url")

	host, ok := r.validateURL(u)

	if ok != nil {
		r.RespondError(w, req, http.StatusBadRequest, "invalid_url")
	}

	s := model.ScrapingRequest{
		URL:  u,
		Host: host,
	}

	result, err := r.Scraping.Execute(&s)

	if err != nil {
		var e *model.AppError
		if errors.As(err, &e) {
			r.RespondError(w, req, e.Code, e.Msg)
			return
		}
		r.RespondError(w, req, http.StatusInternalServerError, "internal_server_error")
		return
	}

	r.Respond(w, req, result)

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

func (r *router) validateURL(u string) (string, error) {
	result, ok := url.Parse(u)

	if ok != nil {
		return "", errors.New("invalid_url")
	}

	return result.Host, nil
}
