package concepts

import "fmt"

type Logger interface {
	Log() string
}

type UserLog struct {
	Name string
}

type ErrorLog struct {
	Code    int
	Message string
}

func (u UserLog) Log() string {
	return fmt.Sprintf("User: %v", u.Name)
}

func (e ErrorLog) Log() string{
	return fmt.Sprintf("Error %d: %v", e.Code, e.Message)

}

func LogAnything(data any) {
	switch v := data.(type) {
	case string:
		fmt.Println("[STRING]", v)
	case int:
		fmt.Println("[NUMBER]", v)
	case Logger:
		fmt.Printf("[%T] %v\n", v, v.Log())
	default: 
		fmt.Sprintln("[UNKNOWN TYPE]", "{%v}", v)
	}
}
