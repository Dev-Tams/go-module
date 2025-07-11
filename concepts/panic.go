package concepts

import "fmt"

func Negative(n int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if n < 0 {
		panic("Negative value not allowed!")
	}

	fmt.Println("processing:", n)
}
