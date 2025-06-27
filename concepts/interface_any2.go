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

func LogAnything(data any) string{
	switch v := data.(type) {
	case string:
		return fmt.Sprintf("[STRING] %v", v)
	case int:
		return fmt.Sprintf("[NUMBER] %v", v)
	case Logger:
		return fmt.Sprintf("[%T] %v", v, v.Log())
	default: 
		return fmt.Sprintf("[UNKNOWN TYPE] %v", v)
	}
}
