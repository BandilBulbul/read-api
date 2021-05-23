package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetOneLine(t *testing.T) {
	got := GetOneLine()
	want := 2
	if got != want {
		t.Errorf("%q got %q want", got, want)
	}

}
func TestGetOddIndex(t *testing.T) {
	data := []string{"Hacker", "Rank"}
	got := GetOddIndex(data)
	fmt.Print(got)
	want := []string{"akr", "ak"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%q got %q want", got, want)
	}

}

func TestGetEvenIndex(t *testing.T) {
	data := []string{"Hacker", "Rank"}
	got := GetEvenIndex(data)
	fmt.Print(got)
	want := []string{"Hce", "Rn"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%q got %q want", got, want)
	}

}

func TestGetIndex(t *testing.T) {
	//var length int
	//length = GetOneLine()
	data := [2]string{"Hacker", "Rank"}
	got := GetIndex(data)
	want := [2][2]string{{"Hce", "akr"}, {"Rn", "ak"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%q got %q want", got, want)

	}
}
