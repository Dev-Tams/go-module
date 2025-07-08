package concepts

import (
	"fmt"
	"time"
)

type OrderError struct {
	Field, Message string
}

type DeliveryError struct {
	DeliveryInfo string
}

type TimeSlotError struct {
	Time string
}

func (o OrderError) Error() string {
	return fmt.Sprintf("Invalid Order [%s]: %s", o.Field, o.Message)
}

func (d DeliveryError) Error() string {
	return fmt.Sprintf("Error with delivery %v", d.DeliveryInfo)
}

func (t TimeSlotError) Error() string {
	return fmt.Sprintf("Error with Time %s", t.Time)
}

//handle error

func HandleError(err error) {
	switch e := err.(type) {
	case DeliveryError:
		fmt.Println("Delivery issue:", e)
	case OrderError:
		fmt.Println("Order Issue", e.Error())
	case TimeSlotError:
		fmt.Println("Time:", e)
	default:
		fmt.Println("Generic Issue", err)
	}
}

func PlaceOrder(qty int, location, item, timeStr string) (string, error) {
	
	parsedTime, _ := time.Parse("15:04", timeStr)
	hour := parsedTime.Hour()
	min := parsedTime.Minute()

	if item == "" {
		return "", OrderError{"item", "Please select an item"}
	}
	if qty == 0 {
		return "", OrderError{"Quantity", "Please specify your quantity"}
	}
	if location == "" || timeStr == "" {
		return "", DeliveryError{"Error: please check delivery information"}
	}
	if hour > 18 || (hour == 18 && min > 0) {
		return "", TimeSlotError{"Orders must be placed before 6:00 PM"}
	}
	return "Order #201 confirmed", nil
}