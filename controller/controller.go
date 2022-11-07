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
	"github.com/scarlet0725/prism-api/scraping"
	"github.com/scarlet0725/prism-api/serializer"
)

type Controller interface {
	ScrapingRequestHandler(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	SupportedSites map[string]string
	ScrapingClient scraping.Client
	Serializer     serializer.Serializer
	cache          cache.Cache
}

func (c *controller) ScrapingRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error

	u := r.URL.Query().Get("target_url")

	host, ok := c.validateURL(u)

	if ok != nil {
		resopondError(w, r, http.StatusBadRequest, "invalid_url")
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
		resopondError(w, r, http.StatusBadRequest, "scraping_error")
	}

	b := bytes.NewReader(s.Data)

	var res model.APIResponse

	switch host {
	case "t.livepocket.jp":
		res, err = c.Serializer.Livepocket(b)
	case "lit.link":
		res, err = c.Serializer.Litlink(b)
	default:
		resopondError(w, r, http.StatusBadRequest, "unsupported_site")
	}

	if err != nil {
		resopondError(w, r, http.StatusBadRequest, "scraping_error")
	}

	respondToClient(w, r, &res)

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

func CreateContoroller(m map[string]string, c *scraping.Client, s *serializer.Serializer, cache *cache.Cache) Controller {
	con := &controller{
		SupportedSites: m,
		ScrapingClient: *c,
		Serializer:     *s,
		cache:          *cache,
	}
	return con
}

func respondToClient(w http.ResponseWriter, r *http.Request, res any) error {
	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(res)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func resopondError(w http.ResponseWriter, r *http.Request, errorCode int, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf("{\"ok\": false, \"error\": \"%s\"}", message)))
}
