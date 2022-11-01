package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scarlet0725/litlink-scraping-api/model"
	"github.com/scarlet0725/litlink-scraping-api/scraping"
	"github.com/scarlet0725/litlink-scraping-api/serializer"
)

type Controller interface {
	ScrapingRequestHandler(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	SupportedSites map[string]string
	ScrapingClient scraping.Client
	Serializer     serializer.Serializer
}

func (c *controller) ScrapingRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	u := r.URL.Query().Get("target_url")

	host, ok := c.validationURL(u)

	if ok != nil {
		resopondError(w, r, http.StatusBadRequest, "invalid_url")
		return
	}

	s, err := c.ScrapingClient.Execute(u)
	if err != nil {
		resopondError(w, r, http.StatusBadRequest, "scraping_error")
	}

	b := bytes.NewReader(s.Data)

	var res model.ApiResponse

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

func (c *controller) validationURL(u string) (string, error) {
	result, ok := url.Parse(u)

	if ok != nil {
		return "", errors.New("invalid_url")
	}

	if _, ok := c.SupportedSites[result.Host]; !ok {
		return "", errors.New("unsupported_site")
	}

	return result.Host, nil
}

func CreateContoroller(m map[string]string, c *scraping.Client, s *serializer.Serializer) Controller {
	con := &controller{
		SupportedSites: m,
		ScrapingClient: *c,
		Serializer:     *s,
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
