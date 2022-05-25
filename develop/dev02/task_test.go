package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	got, err := unpack("a4bc2d5e")
	want := "aaaabccddddde"
	if err != nil {
		t.Errorf("TestUnpack: want %v; got %v", want, err)
	}
	if got != want {
		t.Errorf("TestUnpack: want %v; got %v", want, got)
	}

	got, err = unpack("abcd")
	want = "abcd"
	if err != nil {
		t.Errorf("TestUnpack: want %v; got %v", want, err)
	}
	if got != want {
		t.Errorf("TestUnpack: want %v; got %v", want, got)
	}

	got, err = unpack("45")
	if err == nil {
		t.Errorf("TestUnpack: want error; got %v", got)
	}

	got, err = unpack("")
	want = ""
	if err != nil {
		t.Errorf("TestUnpack: want %v; got %v", want, err)
	}
	if got != want {
		t.Errorf("TestUnpack: want %v; got %v", want, got)
	}

	got, err = unpack("qwe\\4\\5")
	want = "qwe45"
	if err != nil {
		t.Errorf("TestUnpack: want %v; got %v", want, err)
	}
	if got != want {
		t.Errorf("TestUnpack: want %v; got %v", want, got)
	}

	got, err = unpack("qwe\\45")
	want = "qwe44444"
	if err != nil {
		t.Errorf("TestUnpack: want %v; got %v", want, err)
	}
	if got != want {
		t.Errorf("TestUnpack: want %v; got %v", want, got)
	}

	got, err = unpack("qwe\\\\5")
	want = "qwe\\\\\\\\\\"
	if err != nil {
		t.Errorf("TestUnpack: want %v; got %v", want, err)
	}
	if got != want {
		t.Errorf("TestUnpack: want %v; got %v", want, got)
	}
}
