package chatlog

import (
	"fmt"
	"os"
	"testing"
)

func TestGetInitialContinuation(t *testing.T) {
	c := New(os.Getenv("VIDEO_ID"))
	_, err := c.GetInitialContinuation()
	if err != nil {
		t.Fatal(err)
	}
}

func TestFetch(t *testing.T) {
	c := New(os.Getenv("VIDEO_ID"))
	cont, err := c.GetInitialContinuation()
	if err != nil {
		t.Fatal(err)
	}

	for cont != "" {
		r, err := c.FecthChats(cont)
		if err != nil {
			t.Fatal(err)
		}
		cont = r.Continuation

		for i, continuationAction := range r.Action {
			for j, chatAction := range continuationAction.ReplayChatItemAction.Actions {
				chatItem := chatAction.AddChatItemAction.Item
				liveChatTickerItem := chatAction.AddLiveChatTickerItemAction.Item
				switch {
				case chatItem.LiveChatViewerEngagementMessageRenderer.ID != "":
					continue

				case chatItem.LiveChatTextMessageRenderer.ID != "":
					r := chatItem.LiveChatTextMessageRenderer
					m := ""
					for _, v := range r.Message.Runs {
						if v.Text != "" {
							m += v.Text
						} else {
							m += v.Emoji.Image.Accessibility.AccessibilityData.Label
						}
					}
					fmt.Printf("%6s [%s]: %s\n", r.TimestampText.SimpleText, r.AuthorName.SimpleText, m)

				case chatItem.LiveChatMembershipItemRenderer.ID != "":
					r := chatItem.LiveChatMembershipItemRenderer
					m := ""
					for _, v := range r.HeaderSubtext.Runs {
						m += v.Text
					}
					fmt.Printf("       [%s]: %s\n", r.AuthorName.SimpleText, m)

				case chatItem.LiveChatPaidMessageRenderer.ID != "":
					r := chatItem.LiveChatPaidMessageRenderer
					m := ""
					for _, v := range r.Message.Runs {
						if v.Text != "" {
							m += v.Text
						} else {
							m += v.Emoji.Image.Accessibility.AccessibilityData.Label
						}
					}
					fmt.Printf("%6s [%s]: <%s> %s\n", r.TimestampText.SimpleText, r.AuthorName.SimpleText, r.PurchaseAmountText.SimpleText, m)

				case chatItem.LiveChatPlaceholderItemRenderer.ID != "":
					continue

				case liveChatTickerItem.LiveChatTickerSponsorItemRenderer.ID != "":
					r := liveChatTickerItem.LiveChatTickerSponsorItemRenderer
					rr := r.ShowItemEndpoint.ShowLiveChatItemEndpoint.Renderer.LiveChatMembershipItemRenderer
					m := ""
					for _, v := range rr.HeaderSubtext.Runs {
						m += v.Text
					}
					fmt.Printf("       [%s]: %s\n", rr.AuthorName.SimpleText, m)

				case liveChatTickerItem.LiveChatTickerPaidMessageItemRenderer.ID != "":
					r := liveChatTickerItem.LiveChatTickerPaidMessageItemRenderer
					rr := r.ShowItemEndpoint.ShowLiveChatItemEndpoint.Renderer.LiveChatPaidMessageRenderer
					m := ""
					for _, v := range rr.Message.Runs {
						if v.Text != "" {
							m += v.Text
						} else {
							m += v.Emoji.Image.Accessibility.AccessibilityData.Label
						}
					}
					fmt.Printf("%6s [%s]: <%s> %s\n", rr.TimestampText.SimpleText, rr.AuthorName.SimpleText, rr.PurchaseAmountText.SimpleText, m)

				default:
					fmt.Println(c, i, j, continuationAction.ReplayChatItemAction.VideoOffsetTimeMsec)
					t.Fatal("Should be succeeded")
				}
			}
		}
	}
}
