package concepts

import "slices"

import "fmt"


type RoleChecker interface{
	isAdmin() bool
}


type User struct {
	 AcountId string
	 Person
}

type Admin struct {
	Privileges []string
	User
}



// Create a method like func (a Admin) HasPrivilege(p string) bool

// Create another method like func (u User) FullContact() string returning name + email.


func (u User) Greet() string {
	return fmt.Sprintf("welcome, %s!", u.Name)
}

func (a Admin) HasPrivi(priv string) bool{
	return slices.Contains(a.Privileges, priv)
	// for _, p := range a.Privileges{
	// 	if p == priv{
	// 		return true
	// 	}
	// }
	// return false
}

func (u User) FullContact() string{
	return fmt.Sprintf("%v : %v", u.Name, u.Email)
}

func (a Admin) isAdmin() bool{
	return true
}

func (u User) isAdmin() bool{
	return false
}

func CheckAdmin(r RoleChecker){
	if r.isAdmin(){
		fmt.Println("This user is admin")
	}else{
		fmt.Println("Not an admin")
	}
}