package main

import "fmt"

func main() {
	num := 10
	if num%2 == 0 {
		fmt.Println("Number is even ")
	} else {
		fmt.Println("number is odd")
	}

	switch num {
	case 1:
		fmt.Println("num is 1")
	case 10:
		fmt.Println("num is 10")
	default:
		fmt.Println("error")

	}
	// //loops
	numbers := []int{1, 2, 3, 4}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	for _, num := range numbers {
		fmt.Println(num)
	}
	days := []string{"sunday", "monday"}
	for index, day := range days {
		fmt.Println(index, day)
	}
}
