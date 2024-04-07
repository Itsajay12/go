package main

import (
	"fmt"
)

func main() {
	var num1 int = 1
	var ptr = &num1
	fmt.Println(*ptr, ptr)
	var name = [3]string{"a", "b"}
	name[2] = "d"
	fmt.Println(name)
	fmt.Printf("type of %T\n", name)
	languauges := make(map[string]string)
	languauges["PY"] = "python"
	languauges["js"] = "javscript"
	fmt.Println(languauges)
	for key, value := range languauges {
		fmt.Printf("key %v value %v \n", key, value)
	}
	delete(languauges, "PY")
	fmt.Print(languauges)
}
