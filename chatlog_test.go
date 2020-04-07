package chatlog

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New("fzd9nzDpjh0")
	if err != nil {
		t.Fatal("Should be succeeded", err)
	}
}

func TestFetch(t *testing.T) {
	chatlog, _ := New("wEPuZFo1qNc")
	c := 0

	for chatlog.Continuation != "" {
		resp, err := chatlog.Fecth()
		if err != nil {
			t.Fatal("Should be succeeded", err)
		}
		for i, continuationAction := range resp {
			for j, chatAction := range continuationAction.ReplayChatItemAction.Actions {
				chatItem := chatAction.AddChatItemAction.Item
				liveChatTickerItem := chatAction.AddLiveChatTickerItemAction.Item
				if chatItem.LiveChatViewerEngagementMessageRenderer.ID != "" {
					continue
				}
				if chatItem.LiveChatTextMessageRenderer.ID != "" {
					continue
				}
				if chatItem.LiveChatMembershipItemRenderer.ID != "" {
					continue
				}
				if chatItem.LiveChatMembershipItemRenderer.ID != "" {
					continue
				}
				if chatItem.LiveChatPaidMessageRenderer.ID != "" {
					continue
				}
				if chatItem.LiveChatPlaceholderItemRenderer.ID != "" {
					continue
				}
				if liveChatTickerItem.LiveChatTickerSponsorItemRenderer.ID != "" {
					continue
				}
				if liveChatTickerItem.LiveChatTickerPaidMessageItemRenderer.ID != "" {
					continue
				}

				fmt.Println(c, i, j, continuationAction.ReplayChatItemAction.VideoOffsetTimeMsec)
				t.Fatal("Should be succeeded")
			}
		}
		c++
	}
}
