package main

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	got := format(time.Date(0, 0, 0, 12, 34, 56, 0, time.UTC))
	if got != "12:34:56" {
		t.Errorf("TestFormat: want 12:34:56; got %v", got)
	}
}
