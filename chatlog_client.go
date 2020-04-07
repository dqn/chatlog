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

type ChatlogClient struct {
	httpClient *http.Client
}

func newDefaultRequest(method string) (*http.Request, error) {
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", userAgent)
	return req, nil
}

func newClient() *ChatlogClient {
	return &ChatlogClient{&http.Client{}}
}

func (c *ChatlogClient) Do(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c *ChatlogClient) Get(path string, values *url.Values) ([]byte, error) {
	req, err := newDefaultRequest("GET")
	if err != nil {
		return nil, err
	}
	req.URL.Path = path
	req.URL.RawQuery = values.Encode()
	return c.Do(req)
}
