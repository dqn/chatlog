package chatlog

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	chatlog, err := New("fzd9nzDpjh0")
	if err != nil {
		t.Fatal("Should be succeeded", err)
	}
	fmt.Println(chatlog.VideoId)
	fmt.Println(chatlog.Continuation)
}

func TestFetch(t *testing.T) {
	chatlog, _ := New("fzd9nzDpjh0")
	resp, err := chatlog.Fecth()
	if err != nil {
		t.Fatal("Should be succeeded", err)
	}

	for _, continuationAction := range resp {
		for _, chatAction := range continuationAction.ReplayChatItemAction.Actions {
			l := chatAction.AddChatItemAction.Item.LiveChatTextMessageRenderer
			fmt.Printf("%s %s: ", l.TimestampText.SimpleText, l.AuthorName.SimpleText)
			for _, run := range l.Message.Runs {
				if run.Text != "" {
					fmt.Print(run.Text)
				} else if s := run.Emoji.Shortcuts[0]; s != "" {
					fmt.Print(s)
				}
			}
			fmt.Println()
		}
	}
}
