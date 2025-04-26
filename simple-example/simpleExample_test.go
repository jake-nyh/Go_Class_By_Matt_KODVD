package simpleExample

import (
	"fmt"
	"testing"
)

func TestSay(t *testing.T) {
	want := "Hello, Test"
	got := Say("Test")

	fmt.Println(want)
	fmt.Println(got)

	if want != got {
		t.Errorf("wanted %s / got %s", want, got)
	}
}
