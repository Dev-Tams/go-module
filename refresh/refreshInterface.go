package refresh

import "fmt"

////in this file we take a little bit of refresher for Interface

type Logger interface{
	Log(message string)
}

type ConsoleLogger struct{

}

type FileLogger struct{
	Filename string
}

type RemoteLogger struct{
	Endpoint string
}

type LogJob struct{
	Logger Logger
	Message string
}


func (c ConsoleLogger) Log(message string) {
	fmt.Printf("Console: %v\n", message)

}

func (f FileLogger) Log(message string) {
	fmt.Printf("%v %v\n", message, f.Filename)
}

func (r RemoteLogger) Log(message string) {
	fmt.Printf("%v %v\n", message, r.Endpoint)
}


func SendLog(l Logger, message string) {
	l.Log(message)
}
