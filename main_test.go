package main

import (
	"github.com/kyokomi/emoji"
	"testing"
)

func TestGetMessage(t *testing.T) {
	result := GetMessage()
	emoji_msg := emoji.Sprint("Hello :world_map:!")
	if result != emoji_msg {
		t.Errorf("expecting %s, got %s", emoji_msg, result)
	}
}
