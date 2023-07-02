package main

import "fmt"

func main() {
	result := func(arg int) int {
		return arg + 100
	}(123)

	fmt.Println(result) // 223
}
