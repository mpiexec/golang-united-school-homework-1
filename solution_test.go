package solution

import (
	"github.com/kyokomi/emoji"
	"testing"
)

func TestGetMessage(t *testing.T) {
	result := GetMessage()
	expect := emoji.Sprint("Hello :world_map:!")
	if result != expect {
		t.Errorf("expecting %s, got %s", expect, result)
	}
}
