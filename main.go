package main

import (
	// "bufio"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"

	// "go/scanner"
	"os"

	"github.com/Dev-Tams/go-module/concepts"
	"github.com/Dev-Tams/go-module/refresh"
)


var records = [][]string{
	{"Name", "Sex", "Email"},
	{"Tami", "f", "tami@mail.com"},
	{"Tonye", "f", "tonye@mail.com"},
	{"Tony", "m", "tony@mail.com"},
	{"Tommy", "m", "tommy@mail.com"},
	{"T", "f", "t@mail.com"},
}

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


	//interface{} i.e any

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

	//appending any
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


	//errors

	num, err := concepts.CheckPositive(-5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid num:", num)
	}

	email, err := concepts.CheckEmail("tami@mail.com")
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
			fmt.Println("Handle specifically:", tErr)
		} else {
			fmt.Println("Generic error:", err)
		}
	}else{
		fmt.Println(result)
	}

	cart, err := concepts.AddToCart("tamiil.com", 0, 25)
		if err != nil{
			var (
				invErr concepts.InventoryError
				emailErr concepts.EmailValidationError
			)
			if errors.As(err, &invErr) {
				fmt.Println("Custom Inventory Error:", invErr)
			}else if errors.As(err, &emailErr){
				fmt.Println("Email Error:", emailErr)
			}else{
				fmt.Println("Generic error:", err)
			}
		}else{
			fmt.Print(cart)
		}

		//panic
		concepts.Negative(10)
		concepts.Negative(-5)
		concepts.Negative(10)

		admin := concepts.Admin{
		User: concepts.User{
			AcountId: "2",
			Person: concepts.Person{
				Name: "Tami",
				Email: "tami@mail.com",
			},
		},
		Privileges: []string{"read", "write", "delete"},
		}


		user := concepts.User{
			AcountId: "4",
			Person: concepts.Person{
				Name: "Tammy",
				Email: "tammy@mail.com",
			},
		}
		

		
		fmt.Println("Admin Name:", admin.Name)           
		fmt.Println("Admin Email:", admin.Email)
		fmt.Println("Admin ID:", admin.AcountId)
		fmt.Println("Privileges:", admin.Privileges)

		fmt.Println(admin.Greet())
		fmt.Println(user.FullContact())
		fmt.Println(admin.HasPrivi("delete"))
		concepts.CheckAdmin(user)
		concepts.CheckAdmin(admin)


		//opens a file
		file, err := os.Open("logbook.txt")
		if err != nil{
			fmt.Println("error opening file", err)
		}else{
			fmt.Println("File opened successfully!")
		}
		defer file.Close()

		//reads the file by line
		// scanner := bufio.NewScanner(file)

		// for scanner.Scan(){
		// 	line := scanner.Text()
		// 	fmt.Println("line", line)
		// }

		// if  err := scanner.Err(); err != nil{
		// 	fmt.Println("error with reading", err)
		// }

		//create a new file

		newfile, err := os.Create("notes/newfile")
		if err != nil{
			fmt.Println("Error creating new file", err)
		}else{
			fmt.Println("file created")
		}

		defer newfile.Close()

		//write to a file

		
		_, err = newfile.WriteString("Hello, there!\n writing to a new file ")
		if err != nil{
			fmt.Println("error writing to a new file", err)
				return
			}else{
				fmt.Println("file written successfully")
			}
			
			
		//Additionally you cant write to an existing file due to permission errors, you can add by appending

		//appending to the existing file tied to the file var
		file, _ = os.OpenFile("logbook.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer file.Close()

		_, err = file.WriteString("Hello, there!\n Appending to an existing file ")
		if err != nil{
			fmt.Println("error appending to file", err)
			return
		}else{
			fmt.Println("Line appended successfully.")
		}

		rfile, err := os.Create("notes/refresh")
		if err != nil{
			fmt.Println(" error creating file", err)
			return 
		}else{
			fmt.Println("file created" , rfile)
		}

		defer rfile.Close()


		//writing to a json file
		
		jfile, err := os.Create("user.json")

		if err == nil{
			fmt.Println("new json file created")
		}else{
			fmt.Println("error creating json file", err)
			return
		}

		defer jfile.Close()

		encoder := json.NewEncoder(jfile)

		err = encoder.Encode(user)
		if err != nil{
			fmt.Println(" error econding json", err)
			return
		}else{
			fmt.Println("Json written to file")
		}

		//reading from a json file
		jfile, err = os.Open("user.json")
		if err != nil{
			fmt.Println("Error opening json file")
			return
		}else{
			fmt.Println(" reading json file")
		}
		defer jfile.Close()

		//decode

		decode := json.NewDecoder(jfile)
		err = decode.Decode(&user)
		if err == nil{
			fmt.Printf("User loaded from JSON: %+v\n, user", user)
		}else{
			fmt.Println(" Error loading json file")
		}


		//writing a csv

		cfile, err := os.Create("Cfile.csv")
		if err != nil{
			fmt.Println(" Errror creating csv file", err)
			return
		}else{
			fmt.Println(" CSV file created")
		}
		defer cfile.Close()

		csvwrite := csv.NewWriter(cfile)
		
		for _, i := range records{
			err := csvwrite.Write(i)
			if err != nil{
				fmt.Println(" Failed to write csv record")
				return
			}
		}
		defer csvwrite.Flush()
		cfile.Close()
		fmt.Println("csv written ")

		cfile, err = os.OpenFile("Cfile.csv", os.O_RDONLY| os.O_WRONLY, 0644)
		if err != nil{
			fmt.Println("Error opening csv file")
			return
		}else{
			fmt.Println(" reading csv file")
		}
		defer cfile.Close()
		
		reader := csv.NewReader(cfile)
		records, err := reader.ReadAll()
		for i, record := range records {
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			return
		}
		fmt.Printf("Row %d: %v\n", i, record)
		
	}
}
