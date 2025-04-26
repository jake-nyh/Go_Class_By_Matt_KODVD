package simpleExample

import (
	"testing"
)

func TestSay(t *testing.T) {

	type Subtest struct {
		items  []string
		result string
	}

	subtests := []Subtest{
		{
			items:  []string{"shwe", "li"},
			result: "Hello, shwe, li",
		},
		{
			result: "Hello, world",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s, got %s", s, st.result)
		}
	}

}
