package concepts

import "fmt"

//in this file we take a little bit of refresher after learning some topics



type UserSub struct{
	Name, Age, Sub string
	Amount float64
}

//method
func (u UserSub) CheckUSerSub() {
	if u.Amount == 0 {
	 fmt.Printf("for user %v, he is a Normal subscriber\n", u.Name)
		
	}else{
		u.Sub = "Premium"
		fmt.Printf("for user %v, he is a %v subscriber\n", u.Name, u.Sub)
	}
}

//pointer
func (u *UserSub) ChangeUserSub(n float64) string {
	u.Amount += n
	return fmt.Sprintf("user %v added $%v to subscription plan\n", u.Name, u.Amount )	
}