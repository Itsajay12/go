package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Print(presentTime.Format("01-02-2006 15:04"))

}
