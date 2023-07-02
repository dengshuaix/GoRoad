package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "织染未来"
	result1 := strings.HasSuffix(name, "来")
	result2 := strings.HasSuffix(name, "来1")
	fmt.Println(result1, result2) // true false

}
