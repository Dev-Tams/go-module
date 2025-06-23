package concepts

import "fmt"

//in this file we take a little bit of refresher after learning some topics


type UserSub struct{
	Name, Age string
	Sub bool
	Amount float64
}

//method
func (u UserSub) CheckUSerSub() {
	if !u.Sub {
	 fmt.Printf("user %v, is a Normal subscriber\n", u.Name)
		
	}else{
		fmt.Printf("user %v, is a Premium subscriber\n", u.Name)
	}
}

//pointer
func (u *UserSub) PayUserSub(n float64) string {
	u.Amount += n
	return fmt.Sprintf("user %v added $%v to subscription plan\n", u.Name, u.Amount )	
}

func (u *UserSub) SetUserSub(){
	if u.Amount > 0 {
		u.Sub = true
	}else{
		fmt.Println("Issue with subscription, please try again!")
	}
}