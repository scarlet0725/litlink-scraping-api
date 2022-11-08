package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scarlet0725/prism-api/cache"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/parser"
	"github.com/scarlet0725/prism-api/scraping"
)

type Controller interface {
	ScrapingRequestHandler(w http.ResponseWriter, r *http.Request)
	HealthCheckHandler(w http.ResponseWriter, r *http.Request) error
	Respond(http.ResponseWriter, *http.Request, any) error
	RespondError(http.ResponseWriter, *http.Request, int, string) error
}

type controller struct {
	SupportedSites map[string]string
	ScrapingClient scraping.Client
	Parser         parser.Serializer
	cache          cache.Cache
}

func (c *controller) ScrapingRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error

	u := r.URL.Query().Get("target_url")

	host, ok := c.validateURL(u)

	if ok != nil {
		c.RespondError(w, r, http.StatusBadRequest, "invalid_url")
		return
	}

	var s model.ScrapingResult

	cache, ok := c.cache.GetByKey(u)

	switch ok {
	case nil:
		s = model.ScrapingResult{
			Data: cache.Value,
		}
	default:
		s, err = c.ScrapingClient.Execute(u)
		if err != nil {
			break
		}
		cacheData := model.CacheData{
			Key:   u,
			Value: s.Data,
		}

		c.cache.Set(&cacheData, 600)

	}

	if err != nil {
		c.RespondError(w, r, http.StatusBadRequest, "scraping_error")
	}

	b := bytes.NewReader(s.Data)

	var res model.APIResponse

	switch host {
	case "t.livepocket.jp":
		res, err = c.Parser.Livepocket(b)
	case "lit.link":
		res, err = c.Parser.Litlink(b)
	default:
		c.RespondError(w, r, http.StatusBadRequest, "unsupported_site")
	}

	if err != nil {
		c.RespondError(w, r, http.StatusBadRequest, "scraping_error")
	}

	c.Respond(w, r, &res)

}

func (c *controller) Respond(w http.ResponseWriter, r *http.Request, v interface{}) error {
	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func (c *controller) RespondError(w http.ResponseWriter, r *http.Request, code int, msg string) error {
	w.WriteHeader(code)
	_, err := w.Write([]byte(fmt.Sprintf("{\"ok\": false, \"error\": \"%s\"}", msg)))
	return err
}

func (c *controller) validateURL(u string) (string, error) {
	result, ok := url.Parse(u)

	if ok != nil {
		return "", errors.New("invalid_url")
	}

	if _, ok := c.SupportedSites[result.Host]; !ok {
		return "", errors.New("unsupported_site")
	}

	return result.Host, nil
}

func (c *controller) HealthCheckHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("{\"ok\": true}"))
	return err
}

func CreateContoroller(m map[string]string, c *scraping.Client, s *parser.Serializer, cache *cache.Cache) Controller {
	con := &controller{
		SupportedSites: m,
		ScrapingClient: *c,
		Parser:         *s,
		cache:          *cache,
	}
	return con
}
