package simpleExample

import (
	"fmt"
	"os"
)

func Say(name string) {
	fmt.Printf("Hello %s\n", name)
}

func SimpleExample() {

	if len(os.Args) > 1 {
		Say(os.Args[1])
	} else {
		Say("world")
	}

}
