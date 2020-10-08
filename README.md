# chatlog

[![build status](https://github.com/dqn/chatlog/workflows/build/badge.svg)](https://github.com/dqn/chatlog/actions)

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
    switch r := renderer.(type) {
    // System message
    case *LiveChatViewerEngagementMessageRenderer:
      // Print formatted message.
      // e.g. "[Live chat replay is on. Messages that appeared when the stream was live will show up here.]"
      fmt.Println(renderer.ChatMessage())
      return nil

    // Chat message
    case *LiveChatTextMessageRenderer:
      fmt.Println(r.AuthorName.SimpleText)         // Author name
      fmt.Println(r.AuthorExternalChannelID)       // Channel ID
      fmt.Println(r.AuthorPhoto.Thumbnails[0].URL) // Icon URL

      // Print formatted message.
      // e.g. "Alice: hello!"
      fmt.Println(renderer.ChatMessage())
      return nil

    // Membership joining
    case *LiveChatMembershipItemRenderer:
      fmt.Println(r.AuthorName.SimpleText)         // Author name
      fmt.Println(r.AuthorExternalChannelID)       // Channel ID
      fmt.Println(r.AuthorPhoto.Thumbnails[0].URL) // Icon URL

      // Print formatted message.
      // e.g. "[Welcome to Membership!] Bob"
      fmt.Println(renderer.ChatMessage())
      return nil

    // Super Chat
    case *LiveChatPaidMessageRenderer:
      fmt.Println(r.AuthorName.SimpleText)         // Author name
      fmt.Println(r.AuthorExternalChannelID)       // Channel ID
      fmt.Println(r.AuthorPhoto.Thumbnails[0].URL) // Icon URL
      fmt.Println(r.PurchaseAmountText.SimpleText) // Super Chat Amount

      // Print formatted message.
      // e.g. "[$10.00] Carol: bye!"
      fmt.Println(renderer.ChatMessage())
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
// Example for LiveChatTextMessageRenderer.

r, _ := renderer.(*LiveChatTextMessageRenderer)

var buf bytes.Buffer

buf.WriteString(r.AuthorName.SimpleText + "> ")

for _, run := range r.Message.Runs {
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
