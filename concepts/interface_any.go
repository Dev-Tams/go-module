package concepts

import (
	"fmt"
	"strings"
)

//type switch
func CheckType(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println("It's an int:", v)
	case string:
		fmt.Println("It's a string:", v)
	case bool:
		fmt.Println("It's a bool:", v)
	default:
		fmt.Println("Unknown type")
	}
}

// Create a function Process that accepts interface{} and:

// If the input is a string, print it in uppercase.

// If it’s an int, print it multiplied by 10.

// If it’s a bool, print “yes” or “no”.

// Otherwise, print “unsupported type”.

// Test it by passing different values like:

func Process(p any){
	switch v := p.(type){
	case string:
		fmt.Println(strings.ToUpper(v))
	case int:
		fmt.Println(v * 10)
	case bool:
		if v{
			fmt.Println("Yes")
		}else{
			fmt.Println("No")
		}
	default :
		fmt.Println("Unsupported type")
	}
}
