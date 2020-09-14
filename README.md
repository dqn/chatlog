# chatlog

YouTube Live archive chats fecher. Premiered videos are also supported.

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
  c := chatlog.New("VIDEO_ID")

  err := c.HandleChat(func(renderer ChatRenderer) error {
    switch renderer.(type) {
    case *LiveChatViewerEngagementMessageRenderer:
      fmt.Println(renderer.ChatMessage())
      // e.g. "[Live chat replay is on. Messages that appeared when the stream was live will show up here.]"
      return nil

    case *LiveChatTextMessageRenderer:
      fmt.Println(renderer.ChatMessage())
      // e.g. "Alice: hello!"
      return nil

    case *LiveChatMembershipItemRenderer:
      fmt.Println(renderer.ChatMessage())
      // e.g. "[Welcome to Membership!] Bob"
      return nil

    case *LiveChatPaidMessageRenderer:
      fmt.Println(renderer.ChatMessage())
      // e.g. "[$10.00] Carol: bye!"
      return nil

    case *LiveChatPlaceholderItemRenderer:
      fmt.Println(renderer.ChatMessage())
      // (empty)
      return nil
    }
  })

  if err != nil {
    // Handle error.
  }
}
```

Also can custom message.

```go
var buf bytes.Buffer

buf.WriteString(renderer.AuthorName.SimpleText + "> ")

for _, run := range renderer.Message.Runs {
  if run.Text != "" {
    buf.WriteString(run.Text)
  } else {
    buf.WriteString(run.Emoji.Image.Accessibility.AccessibilityData.Label)
  }
}

fmt.Println(buf.String())
```

## Other

`ChatRenderer` is actual YouTube private API response structure. See payloads for more details.

## License

MIT
