package chatlog

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const baseURL = "https://www.youtube.com"

type Chatlog struct {
	VideoId      string
	Continuation string
}

func makeVideoPageRequest(videoId string) (*http.Request, error) {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	v := &url.Values{}
	v.Add("v", videoId)
	req.URL.RawQuery = v.Encode()
	req.URL.Path = "/watch"

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")

	return req, nil
}

func extractContinuation(body []byte) (string, error) {
	s := string(body)
	query := `"continuation":"`
	index := strings.LastIndex(s, query)
	if index == -1 {
		return "", fmt.Errorf("cannot find continuation")
	}

	b := make([]byte, 256)
	for i := index + len(query); s[i] != byte('"'); i++ {
		b = append(b, s[i])
	}

	return string(b), nil
}

func New(videoId string) (*Chatlog, error) {
	req, err := makeVideoPageRequest(videoId)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	continuation, err := extractContinuation(body)
	if err != nil {
		return nil, err
	}

	return &Chatlog{videoId, continuation}, nil
}
