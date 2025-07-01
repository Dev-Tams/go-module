package concepts

import (
	"fmt"
	"strings"
)

func CheckPositive(x int) (int, error) {
	if x <= 0 {
		return 0, fmt.Errorf("input %v: Error: negative number not allowed", x)
	} else {
		return x, nil
	}
}

func CheckEmail(e string) (string, error) {

	x := len(e)
	if !strings.Contains(e, "@") || !strings.Contains(e, ".") {
		return "", fmt.Errorf("email must contain '@' and '.'")
	}
	if x < 5 {
		return "", fmt.Errorf("email is not valid")
	} else {
		return e, nil
	}

}


func CheckAge(age int) (int, error){

	if age < 13{
		return 0, fmt.Errorf("Age %d is not eligible. Must be at least 13", age)
	}else{
		return age, nil
	}
}
