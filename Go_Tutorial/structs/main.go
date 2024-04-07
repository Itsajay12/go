package main

import "fmt"

func main() {
	fmt.Println("structure")
	ajay := User{"ajay", "aj@gmail.com"}
	fmt.Println(ajay)
}

type User struct {
	Uname string
	Email string
}
