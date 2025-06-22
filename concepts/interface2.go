package concepts

import "fmt"

type Notifier interface {
	Notify() string
}

type Email struct {
	Address string
}

type Sms struct {
	Phone_number int
}

func (e Email) Notify() string {
	return fmt.Sprintf("Sending mail to %v", e.Address)
}


func (s Sms) Notify() string {
	return fmt.Sprintf("calling %v", s.Phone_number)
}


func SendNoti(n Notifier){
	fmt.Println(n.Notify())
}