package chatlog

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/dqn/chatlog/chat"
)

type Chatlog struct {
	VideoID      string
	Continuation string
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

func New(videoID string) (*Chatlog, error) {
	client := newClient()
	v := &url.Values{}
	v.Add("v", videoID)
	body, err := client.Get("/watch", v)
	if err != nil {
		return nil, err
	}
	continuation, err := extractContinuation(body)
	if err != nil {
		return nil, err
	}
	return &Chatlog{videoID, continuation, client}, nil
}

func (c *Chatlog) Fecth() ([]chat.ContinuationAction, error) {
	v := &url.Values{}
	v.Add("pbj", "1")
	v.Add("continuation", c.Continuation)

	body, err := c.client.Get("/live_chat_replay/get_live_chat_replay", v)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(body))
	chat := &chat.ChatResponse{}
	if err := json.Unmarshal(body, chat); err != nil {
		return nil, err
	}
	if errors := chat.Response.ResponseContext.Errors.Error; errors != nil {
		return nil, fmt.Errorf("an error occurred: %v", errors[0].DebugInfo)
	}
	cont := chat.Response.ContinuationContents.LiveChatContinuation
	c.Continuation = cont.Continuations[0].LiveChatReplayContinuationData.Continuation
	return cont.Actions, nil
}
