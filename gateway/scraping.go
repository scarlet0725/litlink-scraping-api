package gateway

import (
	"bytes"
	"io"
	"net/http"

	"github.com/scarlet0725/prism-api/model"
)

type Client interface {
	Execute(url string) (model.ScrapingResult, error)
}

type client struct {
}

func (c *client) Execute(url string) (model.ScrapingResult, error) {
	res, err := http.Get(url)

	if err != nil {
		return model.ScrapingResult{}, err
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)

	io.Copy(buf, res.Body)
	b := buf.Bytes()
	s := model.ScrapingResult{
		Data: b,
	}

	return s, nil
}

func NewScrapingClient() Client {
	return &client{}
}
