package main

import "fmt"

func main() {

	User := User{"ajay", "email"}
	defer fmt.Println(User.Name, User.Email)
	User.changeName()
}

type User struct {
	Name  string
	Email string
}

func (obj User) changeName() {
	obj.Email = "raj@mai.com"
	fmt.Println("value of email", obj.Email)
}
