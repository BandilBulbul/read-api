package main

import (
	"testing"
)

//iteration return "aaaaa"
func TestRepeat(t *testing.T) {
	got := Repeat("a")
	expected := "aaaaa"
	if got != expected {
		t.Errorf(" %q got %q expected", got, expected)
	}
}
