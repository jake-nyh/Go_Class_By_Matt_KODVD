package BasicTypes

import "fmt"

func BasicTypes() {
	a := 2
	b := 3.1

	fmt.Printf("a: %T %[1]v\n", a)
	fmt.Printf("b: %T %v\n", b, b)

	a = int(b)

	fmt.Printf("a: %T %[1]v\n", a, a)

}
