package chatlog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL   = "https://www.youtube.com"
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"
)

type Chatlog struct {
	videoID string
	client  *http.Client
}

type FetchResult struct {
	Action       []ContinuationAction
	Continuation string
}

func New(videoID string) *Chatlog {
	return &Chatlog{
		videoID: videoID,
		client:  &http.Client{},
	}
}

func (c *Chatlog) GetInitialContinuation() (string, error) {
	v := url.Values{"v": {c.videoID}}
	body, err := c.fetch("/watch", &v)
	if err != nil {
		return "", err
	}

	cont, err := retrieveContinuation(body)
	if err != nil {
		return "", err
	}

	return cont, nil
}

func (c *Chatlog) FecthChats(continuation string) (*FetchResult, error) {
	v := &url.Values{
		"pbj":          {"1"},
		"continuation": {continuation},
	}

	body, err := c.fetch("/live_chat_replay/get_live_chat_replay", v)
	if err != nil {
		return nil, err
	}

	var chat ChatResponse
	if err := json.Unmarshal(body, &chat); err != nil {
		return nil, err
	}

	if errors := chat.Response.ResponseContext.Errors.Error; errors != nil {
		err = fmt.Errorf(errors[0].ExternalErrorMessage)
		return nil, err
	}

	cont := chat.Response.ContinuationContents.LiveChatContinuation
	r := FetchResult{
		Action:       cont.Actions,
		Continuation: cont.Continuations[0].LiveChatReplayContinuationData.Continuation,
	}

	return &r, nil
}

func (c *Chatlog) HandleChatItem(handler func(item *ChatItem) error) error {
	cont, err := c.GetInitialContinuation()
	if err != nil {
		return err
	}

	for cont != "" {
		r, err := c.FecthChats(cont)
		if err != nil {
			return err
		}
		cont = r.Continuation

		for _, continuationAction := range r.Action {
			for _, chatAction := range continuationAction.ReplayChatItemAction.Actions {
				if err = handler(&chatAction.AddChatItemAction.Item); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (c *Chatlog) fetch(path string, values *url.Values) ([]byte, error) {
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
