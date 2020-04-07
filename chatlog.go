package chatlog

import (
	"fmt"
	"net/url"
	"strings"
)

type Chatlog struct {
	VideoId      string
	Continuation string
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
	client := newClient()

	v := &url.Values{}
	v.Add("v", videoId)

	body, err := client.Get("/watch", v)
	if err != nil {
		return nil, err
	}

	continuation, err := extractContinuation(body)
	if err != nil {
		return nil, err
	}

	return &Chatlog{videoId, continuation}, nil
}
