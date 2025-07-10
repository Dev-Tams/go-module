package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Dev-Tams/go-module/concepts"
	"github.com/Dev-Tams/go-module/refresh"
)

func main() {

	// Methods
	c := concepts.SayHi()
	fmt.Println(c)

	// m := concepts.Calc{X: 3, Y:4}
	// fmt.Println(m.Example())

	// // Pointers
	// // Pointers are used to reference a memory address of a variable.
	// cp := 6.3
	// concepts.ChangePointerValue(&cp)
	// fmt.Println("Multiplied pointer value by 2:", cp)

	// x, y := 5, 10
	// fmt.Println("Before swapped:", x, y)
	// concepts.Swap(&x, &y)
	// fmt.Println("Swapped values:", x, y)

	// car := concepts.Car{Mile: 2400, Brand: "Toyota Avalon", Color: "Red"}

	// ecar := concepts.ElectricCar{
	// 	Car: concepts.Car{Mile: 2400, Brand: "Tesla", Color: "Red"},
	// 	BatteryHealth: 85,
	// }
	// fmt.Println(ecar)
	// fmt.Println(car)

	// fmt.Println("for every test ride, a mile is added to", car.Mile)
	// concepts.CarInfo(&car)

	// fmt.Println("for every test ride, a mile is added to", car.Mile)
	// concepts.CarInfo(&ecar.Car)
	// fmt.Println(car.Repaint("Black"))
	// fmt.Println(ecar.Repaint("Grey"))
	// fmt.Println(ecar.Charge())

	// B := concepts.Person{Name: "Tami", Age: 23, Career: "Programming"}
	// fmt.Println(B)
	// fmt.Println(B.Hbd())
	// fmt.Println(B.Hbd())
	// fmt.Println(B.Hbd())

	cv := concepts.Counter{Value: 5}
	// fmt.Println(cv.Show())
	// cv.Increment()
	// cv.Increment()
	// fmt.Println(cv.Show())
	// cv.Decrement()
	// fmt.Println(cv.Show())
	res, err := cv.Multiply(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	// fmt.Println(cv.Multiply(-1))
	// fmt.Println(cv.Show())
	// fmt.Println(cv.Reset())
	// fmt.Println(cv.Show())

	// Interfaces
	// circle := concepts.Circle{Radius: 5}
	// rectangle := concepts.Rectangle{Width: 4, Height: 6}
	// concepts.PrintArea(circle) //Interface call is abstract and reusable
	// concepts.PrintArea(rectangle)
	// fmt.Println("Area of circle:", circle.Area()) //Direct call is explicit and flexible
	// fmt.Println("Area of rectangle:", rectangle.Area())

	//instead of // calling the Notify method directly, we can use the Notifier interface
	// to send notifications through different channels like Email and SMS.

	// email := concepts.Email{Address: "test@mail.com"}
	// concepts.SendNoti(email)
	// call := concepts.Sms{Phone_number: "+234-00-00-1"}
	// concepts.SendNoti(call)

	// Using a slice of Notifier interface to send notifications
	// This allows us to handle multiple notification types in a single loop.
	// 	notifiers := []concepts.Notifier{
	// 	concepts.Email{Address: "test@mail.com"},
	// 	concepts.Sms{Phone_number: "+234-00-00-1"},
	// }

	// for _, n := range notifiers {
	// 	concepts.SendNoti(n)
	// }

	payments := []concepts.PaymentMethods{
		{Method: concepts.CreditCard{Name: "Tami"}, Amount: 300},
		{Method: concepts.Paypal{Email: "tami@mail.com"}, Amount: 150},
		{Method: concepts.CryptoWallet{Address: "0x123abc456"}, Amount: 820},
	}
	for _, p := range payments {
		concepts.ProcessPayment(p.Method, p.Amount) // Process payment using the interface method
	}

	// //refresher
	// u := refresh.UserSub{Name: "Tami", Amount: 0}
	// u.CheckUSerSub()
	// u.SetUserSub()
	// fmt.Print(u.PayUserSub(30))
	// u.SetUserSub()
	// u.CheckUSerSub()

	syslogger := []refresh.LogJob{
		{Logger: refresh.ConsoleLogger{}, Message: "prints to screen"},
		{Logger: refresh.FileLogger{Filename: "sys"}, Message: "writing to file"},
		{Logger: refresh.RemoteLogger{Endpoint: "sys123"}, Message: "sending log to server"},
	}
	for _, p := range syslogger {
		refresh.SendLog(p.Logger, p.Message)
	}

	var a any = "Tami"

	s, ok := a.(string)
	if ok {
		fmt.Println("Value:", s)
	} else {
		fmt.Println("interface conversion: interface {} is not string")
	}
	var x any = "Tami"
	concepts.CheckType(x)
	concepts.CheckType(true)
	concepts.Process("Test Type")
	concepts.Process(10)
	concepts.Process(true)
	concepts.Process(3.24)
	concepts.Process(c)
	concepts.Process(false)

	var logbook []string

	logbook = append(logbook, concepts.LogAnything("Hello there!"))
	logbook = append(logbook, concepts.LogAnything(43))
	logbook = append(logbook, concepts.LogAnything(concepts.UserLog{Name: "Tami"}))
	logbook = append(logbook, concepts.LogAnything(concepts.ErrorLog{Code: 404, Message: "Not found"}))
	logbook = append(logbook, concepts.LogAnything(123.456))

	filename := "logbook.txt"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	defer f.Close()

	for _, entry := range logbook {
		_, err := f.WriteString(entry + "\n")
		if err != nil {
			fmt.Println("Error writing entry:", err)
			return
		}
	}

	fmt.Println("Log entries appended to", filename)

	num, err := concepts.CheckPositive(-5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid num:", num)
	}

	email, err := concepts.CheckEmail("tamialemu@mail.com")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("%v is a valid email\n", email)
	}

	age, err := concepts.CheckAge(12)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("age %v is allowed\n", age)
	}

	username, err := concepts.CheckLogin("Tami", "password123")
	if err != nil {
		if loginErr, ok := err.(concepts.LoginError); ok {
			fmt.Println("Custom error caught:", loginErr)
		} else {
			fmt.Println("Generic error:", err)
		}
	} else {
		fmt.Println("Successfully logged in as", username)
	}

	usernames, err := concepts.Inactive("Tami", "12")
	if err != nil {
		if inactiveErr, ok := err.(concepts.LoginError); ok {
			fmt.Println("Error:", inactiveErr)
		} else {
			fmt.Println("Generic error:", err)
		}
	} else {
		fmt.Println("user", usernames, "is active")
	}

	status, err := concepts.PlaceOrder(4, "Chita Avanue", "Bread", "18:30")
	if err != nil {
		concepts.HandleError(err)
	} else {
		fmt.Println(status)
	}

	result, err := concepts.ProcessPayment(concepts.CreditCard{Name: "Tami"}, 0)
	if err != nil {
		var tErr concepts.TransactionError
		if errors.As(err, &tErr) {
			fmt.Println("Handle specifically:", tErr.Reason)
		} else {
			fmt.Println("Generic error:", err)
		}
	}else{
		fmt.Println(result)
	}

}
