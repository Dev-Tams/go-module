package concepts

import "fmt"


//this func changes the value pointed to by the pointer p
func ChangePointerValue(p *float64) float64{
	fmt.Println("Before change:", *p) 
	*p = *p * 2
	return *p // returns the new value pointed to by p
}

func Swap(x, y *int){
	temp := *x // dereference x
	*x = *y
	*y = temp
}