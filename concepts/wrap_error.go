package concepts

import "fmt"

type InventoryError struct {
	Item int
	Reason string
}

func (i InventoryError) Error() string {
	return fmt.Sprintf("failed to add item: %v", i.Reason)
}

func AddToCart(item int) (string, error){
	if item == 0 {
		return "", InventoryError{Reason: "Out of stock"}
	}
	return fmt.Sprintln("Item added to cart"), nil
}