package main

import "fmt"

func main() {
	// 方式 1
	for i := 1; i < 10; {
		fmt.Println(i)
		i += 1
	}

	//方式 2
	// i++  等同于 i=i+1
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

}
