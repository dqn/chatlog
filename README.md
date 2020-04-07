# chatlog

Fetch YouTube Live archive chats

## Installation

```bash
$ go get github.com/dqn/chatlog
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/dqn/chatlog"
)

func main() {
	c, err := chatlog.New("VIDEO_ID")
	if err != nil {
		panic(err)
	}
	resp, err := c.Fecth()
	if err != nil {
		panic(err)
	}
	for _, continuationAction := range resp {
		for _, chatAction := range continuationAction.ReplayChatItemAction.Actions {
			if r := chatAction.AddChatItemAction.Item.LiveChatTextMessageRenderer; r.ID != "" {
				for _, run := range r.Message.Runs {
					if run.Text != "" {
						fmt.Println(run.Text)
					} else if e := run.Emoji; e.EmojiId != "" {
						fmt.Print(e.Image.Accessibility.AccessibilityData.Label)
					}
				}
			}
		}
	}
	c.Fecth() // next chats
}
```

## More details

See [chat/chat_response.go](chat/chat_response.go)

## License

MIT
