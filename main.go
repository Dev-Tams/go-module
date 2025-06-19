package main

import (
	"fmt"

	"github.com/Dev-Tams/go-module/concepts"
)

func main() {

	// Methods
	c := concepts.SayHi()
	fmt.Println(c)

	m := concepts.Calc{X: 3, Y:4}
	fmt.Println(m.Example())


	// Pointers
	// Pointers are used to reference a memory address of a variable.
	cp := 6.3
	concepts.ChangePointerValue(&cp)
	fmt.Println("Multiplied pointer value by 2:", cp)

	x, y := 5, 10
	fmt.Println("Before swapped:", x, y)	
	concepts.Swap(&x, &y)
	fmt.Println("Swapped values:", x, y)	

	}