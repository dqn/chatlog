package chatlog

import (
	"fmt"
	"os"
	"testing"
)

func TestHandleChatItem(t *testing.T) {
	c := New(os.Getenv("VIDEO_ID"))

	err := c.HandleChat(func(renderer ChatRenderer) error {
		switch renderer.(type) {
		case *LiveChatViewerEngagementMessageRenderer:
			fmt.Println(renderer.ChatMessage())
			return nil

		case *LiveChatTextMessageRenderer:
			fmt.Println(renderer.ChatMessage())
			return nil

		case *LiveChatMembershipItemRenderer:
			fmt.Println(renderer.ChatMessage())
			return nil

		case *LiveChatPaidMessageRenderer:
			fmt.Println(renderer.ChatMessage())
			return nil

		case *LiveChatPlaceholderItemRenderer:
			fmt.Println(renderer.ChatMessage())
			return nil

		default:
			return fmt.Errorf("unknown renderer type: %v", renderer)
		}
	})

	if err != nil {
		t.Fatal(err)
	}
}
