package main

import "fmt"

func main() {

	// 1. switch 条件表达式

	switch 1 + 2 {
	case 1:
		fmt.Println(1)

	case 2:
		fmt.Println(2)

	case 3:
		fmt.Println(3)
	default:
		fmt.Println(4)
	}

	// 2. switch 变量
	var page int
	fmt.Scan(&page)
	switch page {
	case 1:
		fmt.Println(1)

	case 2:
		fmt.Println(2)

	case 3:
		fmt.Println(3)
	default:
		fmt.Println(4)
	}

}
