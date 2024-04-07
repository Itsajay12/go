package main

import "fmt"

func main() {
	fmt.Println(sum(1, 3))
	proAdd(1, 2, 3, 1, 5, 4)
}
func sum(num1 int, num2 int) int {
	fmt.Println("hi", num1)
	return num1 + num2
}
func proAdd(values ...int) {
	var total int = 0
	for _, value := range values {
		total += value
	}
	fmt.Println(total)
}
