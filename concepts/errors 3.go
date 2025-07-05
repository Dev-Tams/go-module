package concepts

import "fmt"

type OrderError struct {
	Field, Message string
}

type DeliveryError struct {
	DeliveryInfo string
}

func (o OrderError) Error() string {
	return fmt.Sprintf("Invalid Order [%s]: %s", o.Field, o.Message)
}

func (d DeliveryError) Error() string {
	return fmt.Sprintf("Error with delivery %v", d.DeliveryInfo)
}

//handle error

func HandleError(err error) {
	switch e := err.(type){
	case DeliveryError:
		fmt.Println("Delivery issue:", e)
	case OrderError:
		fmt.Println("Order Issue", e.Error())
	default:
		fmt.Println("Generic Issue", err)
	}
}


func PlaceOrder(qty int, location, item, time string) (string, error){
	if item == "" {
		return "", OrderError{"item", "Please select an item"}
	}
	if qty == 0 {
		return "", OrderError{"Quantity", "Please specify your quantity"}
	}
	if location == "" || time == ""{
		return "", DeliveryError{"Error: please check delivery information"}
	}
	return "Order #201 confirmed", nil
}