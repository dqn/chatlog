package chatlog

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/dqn/chatlog/chat"
)

type Chatlog struct {
	VideoId      string
	continuation string
	client       *chatlogClient
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
	return &Chatlog{videoId, continuation, client}, nil
}

func (c *Chatlog) Fecth() ([]chat.ContinuationAction, error) {
	v := &url.Values{}
	v.Add("pbj", "1")
	v.Add("continuation", c.continuation)

	body, err := c.client.Get("/live_chat_replay/get_live_chat_replay", v)
	if err != nil {
		return nil, err
	}

	chat := &chat.ChatResponse{}
	if err := json.Unmarshal(body, chat); err != nil {
		return nil, err
	}
	cont := chat.Response.ContinuationContents.LiveChatContinuation
	c.continuation = cont.Continuations[0].LiveChatReplayContinuationData.Continuation
	return cont.Actions, nil
}
