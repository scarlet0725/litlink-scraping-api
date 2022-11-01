package scraping

import (
	"bytes"
	"io"
	"net/http"

	"github.com/scarlet0725/litlink-scraping-api/model"
)

func GetSiteContent(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res.Body, err
}

type Client interface {
	Execute(url string) (model.ScrapingResult, error)
}

type client struct {
}

func (c *client) Execute(url string) (model.ScrapingResult, error) {
	r, err := GetSiteContent(url)
	if err != nil {
		return model.ScrapingResult{}, err
	}
	defer r.Close()

	buf := new(bytes.Buffer)

	io.Copy(buf, r)
	b := buf.Bytes()
	s := model.ScrapingResult{
		Data: b,
	}

	return s, nil
}

func CreateClient() *client {
	return &client{}
}
