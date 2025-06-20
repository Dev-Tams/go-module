package main

import (
	"fmt"

	"github.com/Dev-Tams/go-module/concepts"
)

func main() {

	// Methods
	c := concepts.SayHi()
	fmt.Println(c)

	m := concepts.Calc{X: 3, Y:4}
	fmt.Println(m.Example())


	// Pointers
	// Pointers are used to reference a memory address of a variable.
	cp := 6.3
	concepts.ChangePointerValue(&cp)
	fmt.Println("Multiplied pointer value by 2:", cp)

	x, y := 5, 10
	fmt.Println("Before swapped:", x, y)	
	concepts.Swap(&x, &y)
	fmt.Println("Swapped values:", x, y)	


	car := concepts.Car{Mile: 2400, Brand: "Toyota Avalon", Color: "Red"}

	ecar := concepts.ElectricCar{
		Car: concepts.Car{Mile: 2400, Brand: "Tesla", Color: "Red"},
		BatteryHealth: 85,
	}
	fmt.Println(ecar)
	fmt.Println(car)

	fmt.Println("for every test ride, a mile is added to", car.Mile)
	concepts.CarInfo(&car)

	fmt.Println("for every test ride, a mile is added to", car.Mile)
	concepts.CarInfo(&ecar.Car)
	fmt.Println(car.Repaint("Black"))
	fmt.Println(ecar.Repaint("Grey"))

	B := concepts.Person{Name: "Tami", Age: 23, Career: "Programming"}
	fmt.Println(B)
	fmt.Println(B.Hbd())
	fmt.Println(B.Hbd())
	fmt.Println(B.Hbd())

	cv := concepts.Counter{Value: 5}
	fmt.Println(cv.Show()) 
    cv.Increment()
    cv.Increment()
    fmt.Println(cv.Show()) 
    cv.Decrement()
    fmt.Println(cv.Show()) 
	fmt.Println(cv.Multiply(3))
	fmt.Println(cv.Show()) 
    fmt.Println(cv.Reset())
    fmt.Println(cv.Show()) 
	}