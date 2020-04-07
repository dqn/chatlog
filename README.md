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
	"github.com/dqn/chatlog"
)

func main() {
	c, err := chatlog.New("VIDEO_ID")
	if err != nil {
		panic(err)
	}

	for c.Continuation != "" {
		resp, err := c.Fecth()
		if err != nil {
			panic(err)
		}
		for _, continuationAction := range resp {
			for _, chatAction := range continuationAction.ReplayChatItemAction.Actions {
				chatItem := chatAction.AddChatItemAction.Item
				liveChatTickerItem := chatAction.AddLiveChatTickerItemAction.Item

				switch {
				case chatItem.LiveChatViewerEngagementMessageRenderer.ID != "":
					// ...
				case chatItem.LiveChatTextMessageRenderer.ID != "":
					// ...
				case chatItem.LiveChatMembershipItemRenderer.ID != "":
					// ...
				case chatItem.LiveChatMembershipItemRenderer.ID != "":
					// ...
				case chatItem.LiveChatPaidMessageRenderer.ID != "":
					// ...
				case chatItem.LiveChatPlaceholderItemRenderer.ID != "":
					// ...
				case liveChatTickerItem.LiveChatTickerSponsorItemRenderer.ID != "":
					// ...
				case liveChatTickerItem.LiveChatTickerPaidMessageItemRenderer.ID != "":
					// ...
				}
			}
		}
	}
}
```

## More details

See [chat/chat_response.go](chat/chat_response.go)

## License

MIT
