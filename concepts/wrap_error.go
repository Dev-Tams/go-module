package concepts

import "fmt"

type InventoryError struct {
	Item, Qty  int
	Reason string
}

type EmailValidationError struct{
	Email string
}

func (i InventoryError) Error() string {
	return fmt.Sprintf("failed to add item: %v", i.Reason)
}
func (e EmailValidationError) Error() string {
	return fmt.Sprintf("Error with mail: %s", e.Email)
}

func AddToCart(email string, item, qty int) (string, error){

	cem, err := CheckEmail(email)
	if err != nil{
		return "", EmailValidationError{Email: "invalid address"} 
	}

	if item == 0 {
		return "", InventoryError{Reason: "Out of stock"}
	}
	if qty > 35 {
		return "", InventoryError{Reason: "Bulk purchase detected"}
	}
	return fmt.Sprintln("Item added to cart", cem), nil
}