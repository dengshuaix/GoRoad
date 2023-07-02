package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "织染未来"
	result1 := strings.HasPrefix(name, "织染")
	result2 := strings.HasPrefix(name, "染")
	fmt.Println(result1, result2) // true false

}
