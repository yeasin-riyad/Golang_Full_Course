package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	u := User{Name: "Sangam", Age: 20}
	fmt.Println(u.Intro())

}

// val receiver means this method receives a copy of the User struct, 
// so any changes made to the User struct inside the method will not affect 
// the original struct outside the method.
func (u User) Intro() string {

	return fmt.Sprintf("Hi, I am %s", u.Name)
}
