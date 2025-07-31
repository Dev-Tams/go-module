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


type Car struct{
	Color, Brand string
	Mile float64
}

type Person struct{
	Name string 	`json:"name"`
	Career string 	`json:"career,omitempty"`
	Email string	`json:"email"`
	Age int
}


type Counter struct{
	Value int
}

func(c *Counter) Increment(){
	c.Value++
}
func(c *Counter) Decrement(){
	c.Value--
}
func(c *Counter) Reset() string{
	c.Value = 0
	return fmt.Sprintf("Current count has been reset to: %v", c.Value)
}
func(c *Counter) Show() string{
	return fmt.Sprintf("Current count is:%v", c.Value)
} 

func(c *Counter) Multiply(n int) (string, error){
	c.Value *= n
	if n == 0 {
		return "", fmt.Errorf("cannot pass O as value")
	}
	return fmt.Sprintf("Current count multiplied by %v is: %v", n, c.Value), nil
} 




func CarInfo(c *Car) string{
	c.Mile++
	return fmt.Sprintf("for %v, with color %v, it has a mile of %0.2f\n", c.Brand, c.Color, c.Mile )
}

func (c *Car) Repaint(newColor string) string{
		c.Color = newColor
		return fmt.Sprintf("for %v, with mile of %0.2f, repainted color to %v \n", c.Brand, c.Mile, c.Color )
}

func (c *ElectricCar) Charge() string{
		c.BatteryHealth ++
		return fmt.Sprintf("Charging %v..., Battery at %v", c.Brand, c.BatteryHealth) 
}


func(h *Person) Hbd() string {
	h.Age++
	return  fmt.Sprintf("Happy birthday %v, you are now %v year's old \n Whats the next intent on your %v career", h.Name, h.Age, h.Career)
	
}
