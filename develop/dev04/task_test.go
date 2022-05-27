package main

import (
	"reflect"
	"testing"
)

func TestSortAnagrams(t *testing.T) {
	words := []string{"столик", "пятак", "листок", "пятка", "слиток", "тяпка", "кот"}

	got := sortAnagrams(words)
	want := map[string][]string{
		"пятак":  {"пятак", "пятка", "тяпка"},
		"столик": {"листок", "слиток", "столик"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestSortAnagrams: want %v; got %v", want, got)
	}
}
