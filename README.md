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
	c, err := chatlog.New("RVP-epwlIAk")
	if err != nil {
		panic(err)
	}
	resp, err := c.Fecth()
	if err != nil {
		panic(err)
	}
	for _, continuationAction := range resp {
		for _, chatAction := range continuationAction.ReplayChatItemAction.Actions {
			if l := chatAction.AddChatItemAction.Item.LiveChatTextMessageRenderer; &l != nil {
				for _, run := range l.Message.Runs {
					if run.Text != "" {
						fmt.Println(run.Text)
					} else if e := run.Emoji; &e != nil {
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
