package concepts

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) string
}

type CreditCard struct {
	Name string
}
type Paypal struct {
	Email string
}

type CryptoWallet struct {
	Address string
}

type PaymentMethods struct {
	Method PaymentMethod
	Amount float64
}

type TransactionError  struct{
	Reason string
}


func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("$%v was paid by %v with credit card ", amount, cc.Name)
}

func (pp Paypal) Pay(amount float64) string {
	return fmt.Sprintf("$%v was paid by %v via Paypal ", amount, pp.Email)

}

func (cw CryptoWallet) Pay(amount float64) string {
	return fmt.Sprintf("$%v was paid through %v address ", amount, cw.Address)
}

func (t TransactionError) Error() string{
	return fmt.Sprintf("card declined: %v", t.Reason)
}

func ProcessPayment(p PaymentMethod, amount float64) (string, error) {
	if amount == 0.0{
		return "",   TransactionError{Reason: "Insufficient funds"}

	}else{
		return fmt.Sprintln(p.Pay(amount)), nil
	}
}





