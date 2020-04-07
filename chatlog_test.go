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
	c, err := chatlog.Fecth()
	if err != nil {
		t.Fatal("Should be succeeded", err)
	}

	for _, v := range c.Response.ContinuationContents.LiveChatContinuation.Actions {
		for _, w := range v.ReplayChatItemAction.Actions {
			l := w.AddChatItemAction.Item.LiveChatTextMessageRenderer
			fmt.Printf("%s %s: ", l.TimestampText.SimpleText, l.AuthorName.SimpleText)
			for _, x := range l.Message.Runs {
				if x.Text != "" {
					fmt.Print(x.Text)
				} else {
					fmt.Print(x.Emoji.Shortcuts[0])
				}
			}
			fmt.Println()
		}
	}
}
