package simpleExample

import (
	"fmt"
	"os"
)

func Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func SimpleExample() {

	if len(os.Args) > 1 {
		Say(os.Args[1])
	} else {
		Say("world")
	}

}
