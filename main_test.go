package main

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	result := GetMessage()
	if result != "" {
		t.Errorf("expecting \"\", got %s", result)
	}
}
