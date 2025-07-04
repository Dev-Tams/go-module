package concepts

import (
	"fmt"
	"strconv"
)

// Create a LoginError struct with fields like Username, Reason

// Implement Error() on it

// Use it in a CheckLogin(username, password string) function that returns an error if something is wrong

// In main, call CheckLogin(...) and handle the custom error

type LoginError struct{
	Username, Reason string
}

func (l LoginError) Error() string {
	 return fmt.Sprintf("Login failed for user '%s': %s", l.Username, l.Reason)
}

func CheckLogin(username, password string) (string, error){
	if username == "" || password == "" {
		return "", LoginError{Username: username, Reason: "Fields cannot be empty, try again"}
	}
	return username, nil
}

func Inactive(username, days string) (string, error){
	  x, _ := strconv.Atoi(days)
	 if x > 30 {
		return "", LoginError{Username: username, Reason: "User has been inactive, please contact support"}
	 }
	 return username, nil
}




		// +-----------------------------+
		// |        CheckLogin()        |
		// +-----------------------------+
		// 		|
		// 		v
		// +-----------------------------+
		// | Is username or password    |
		// | empty?                     |
		// +-----------------------------+
		// 	| Yes                  | No
		// 	v                      v
		// +-------------------+   +------------------------+
		// | Return LoginError |   | Return username, nil  |
		// | with Username &   |   +------------------------+
		// | Reason fields     |
		// +-------------------+

		// ↓
		// Handled in caller:
		// -----------------------
		// if err != nil {
		// 	if loginErr, ok := err.(LoginError) {
		// 		// Handle custom LoginError
		// 	} else {
		// 		// Generic fallback
		// 	}
		// }
