package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	num1 := bufio.NewReader(os.Stdin)
	newnum1, _ := num1.ReadString('\n')
	num2 := bufio.NewReader(os.Stdin)
	newnum2, _ := num2.ReadString('\n')
	nnewnum1, _ := strconv.ParseFloat(strings.TrimSpace(newnum1), 64)
	nnewnum2, _ := strconv.ParseFloat(strings.TrimSpace(newnum2), 64)
	fmt.Println(nnewnum1 + nnewnum2)
}
