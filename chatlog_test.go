package chatlog

import (
	"fmt"
	"testing"
)

func TestExampleSuccess(t *testing.T) {
	chatlog, err := New("fzd9nzDpjh0")
	if err != nil {
		t.Fatal("Should be succeeded", err)
	}
	fmt.Println(chatlog.VideoId)
	fmt.Println(chatlog.Continuation)
}
