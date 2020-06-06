package chatlog

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURL   = "https://www.youtube.com"
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"
)

type chatlogClient struct {
	client *http.Client
}

func newClient() *chatlogClient {
	return &chatlogClient{&http.Client{}}
}

func (c *chatlogClient) Get(path string, values *url.Values) ([]byte, error) {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.URL.Path = path
	req.URL.RawQuery = values.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
