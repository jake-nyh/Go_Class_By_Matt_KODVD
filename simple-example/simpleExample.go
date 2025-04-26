package simpleExample

import (
	"fmt"
	"os"
	"strings"
)

func Say(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}

	return "Hello, " + strings.Join(names, ", ")
}

func SimpleExample() {
	fmt.Println(Say(os.Args[1:]))
}
