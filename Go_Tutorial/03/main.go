package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcom := "welcome to user inpt"
	fmt.Println(welcom)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter something")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
	if err != nil {

		fmt.Println(err)
	} else {
		fmt.Println(numrating + 1)
	}
}
