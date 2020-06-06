package chatlog

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type Chatlog struct {
	videoID      string
	continuation string
	client       *chatlogClient
}

func New(videoID string) (*Chatlog, error) {
	client := newClient()
	var v url.Values
	v.Add("v", videoID)

	body, err := client.Get("/watch", &v)
	if err != nil {
		return nil, err
	}

	continuation, err := retrieveContinuation(body)
	if err != nil {
		return nil, err
	}

	return &Chatlog{videoID, continuation, client}, nil
}

func retrieveContinuation(body []byte) (string, error) {
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

func (c *Chatlog) Fecth() ([]ContinuationAction, error) {
	v := &url.Values{
		"pbj":          {"1"},
		"continuation": {c.continuation},
	}

	body, err := c.client.Get("/live_chat_replay/get_live_chat_replay", v)
	if err != nil {
		return nil, err
	}

	var chat ChatResponse
	if err := json.Unmarshal(body, &chat); err != nil {
		return nil, err
	}

	if errors := chat.Response.ResponseContext.Errors.Error; errors != nil {
		err = fmt.Errorf("an error occurred: %v", errors[0].DebugInfo)
		return nil, err
	}

	cont := chat.Response.ContinuationContents.LiveChatContinuation
	c.continuation = cont.Continuations[0].LiveChatReplayContinuationData.Continuation

	return cont.Actions, nil
}
