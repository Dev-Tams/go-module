package concepts

import "fmt"


func CheckPositive(x int) (int, error) {
	if x <= 0 {
		return 0, fmt.Errorf("input %v: Error: negative number not allowed", x)
	}else{
		return x, nil
	}
}



