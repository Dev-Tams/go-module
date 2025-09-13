package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Dev-Tams/go-module/concepts"
)


func main(){

	t := concepts.User{
		Person: concepts.Person{
			Name:  "Tams",
			Email: "tams@example.com",
		},
		AccountId: "A123",
	}
	m := concepts.SayHi()

	fmt.Println(m)


	p := 6.0
	concepts.ChangePointerValue(&p)
	fmt.Println(p)

	a := 4
	b := 5
	concepts.Swap(&a, &b)
	fmt.Println(a, b)


	my := concepts.Car{Color: "red" , Brand: "Toyota",  Mile: 0}

	y := concepts.CarInfo(&my)
	x := concepts.CarInfo(&my)
	fmt.Println(y, x)


	cc := concepts.Counter{ Value: 5}
	cc.Decrement()
	fmt.Println(cc.Show())
	cc.Multiply(2)
	fmt.Println(cc.Show())
	cc.Reset()
	fmt.Println(cc.Show())
	
	yy:= concepts.Person{
		Name: "T",
		Career: "Trainer",
		Age: 20,
	}

	fmt.Println(yy)
	
	fmt.Println(yy.Hbd())

	shape := []concepts.Shape{concepts.Circle{Radius: 5}, 
		concepts.Rectangle{Width: 40, Height: 20},
	}

	for _, sh := range shape{
		concepts.PrintArea(sh)
	}

	noti := []concepts.Notifier{
		concepts.Email{Address: "t@mail.com"}, 
		concepts.Sms{Phone_number: "12536"},
	}
	
	for _, i := range noti{
		concepts.SendNoti(i)
	}


	pay := []concepts.PaymentMethods{
		{Method: concepts.CreditCard{Name: "mastercard"}, Amount: 00},
		{Method: concepts.Paypal{Email: "t@paypal.com"}, Amount: 40},
		{Method: concepts.CryptoWallet{Address: "0xxxxxxxxxxxx"}, Amount: 700},
	}

	for _, j := range pay{
		errpay, err := concepts.ProcessPayment(j.Method, j.Amount)
		if err != nil{
			var tErr concepts.TransactionError
			if errors.As(err, &tErr){
				fmt.Println(tErr)
			}else{
				println(err)
			}
		}
		fmt.Println(errpay)
	}

	

	concepts.CheckType(pay)
	concepts.Process(pay)
	mm := concepts.LogAnything(concepts.CarInfo(&my))
	fmt.Println(mm)

	var logbook []string 

	logbook = append(logbook, concepts.LogAnything(mm))

	filename := "logbook.txt"
	
	for i := range logbook{
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil{
			fmt.Println("error with file", err)
		}
		defer f.Close()

		_, err = f.WriteString(logbook[i] + "\n")
		if err != nil {
			fmt.Println("error writing to file")
		}
		fmt.Println("file written to succesfully")
	}

	//write to json
	userjson := "usersjson.json"
	us, err := os.Create(userjson)
	if err != nil{
		fmt.Println("error creating file")
	}
	defer us.Close()
	encoder := json.NewEncoder(us)
	err = encoder.Encode(t)
	encoder.SetIndent("", "  ")
	if err != nil{
		fmt.Println("error encoding struct to json file")
	}else {
		fmt.Println("written to json")
	}
	
	//read from json
	us, err = os.OpenFile(userjson, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("error reading from file")
	}

	decoder := json.NewDecoder(us)	
	err = decoder.Decode(&t)
	if err != nil{
		fmt.Println("error decoding struct")
		return
	}
	fmt.Printf("decode successfully %v\n", t)
}
