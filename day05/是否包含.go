package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "测试"
	result := strings.Contains(name, "测")
	fmt.Println(result) //true
}
