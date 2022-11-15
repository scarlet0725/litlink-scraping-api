package gateway

import (
	"bytes"
	"io"
	"net/http"

	"github.com/scarlet0725/prism-api/repository"
)

type HttpClient struct {
	client *http.Client
}

func NewHTTPClient() repository.HTTPRepository {
	client := &http.Client{}
	return &HttpClient{
		client: client,
	}
}

func (c *HttpClient) Get(url string, header map[string]string, param map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	if len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	if len((param)) > 0 {
		q := req.URL.Query()
		for k, v := range param {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)
	b := buf.Bytes()

	return b, nil
}
