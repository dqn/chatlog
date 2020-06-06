# chatlog

Fetch chats from YouTube Live archive.

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
	c := chatlog.New("VIDEO_ID")

	// handle simply
	err := c.HandleChatItem(func(item *chatlog.ChatItem) error {
		switch {
		case item.LiveChatViewerEngagementMessageRenderer.ID != "":
			// ...
		case item.LiveChatTextMessageRenderer.ID != "":
			// ...
		case item.LiveChatMembershipItemRenderer.ID != "":
			// ...
		case item.LiveChatMembershipItemRenderer.ID != "":
			// ...
		case item.LiveChatPaidMessageRenderer.ID != "":
			// ...
		case item.LiveChatPlaceholderItemRenderer.ID != "":
			// ...
		}

		return nil
	})

	if err != nil {
		// handle error
	}
}
```

Also, can handle manually.

```go
package main

import (
	"github.com/dqn/chatlog"
)

func main() {
	c := chatlog.New("VIDEO_ID")

	cont, err := c.GetInitialContinuation()
	if err != nil {
		// handle error
	}

	for cont != "" {
		r, err := c.FecthChats(cont)
		if err != nil {
			// handle error
		}
		cont = r.Continuation

		for _, continuationAction := range r.Action {
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

See [payload.go](payload.go)

## License

MIT
