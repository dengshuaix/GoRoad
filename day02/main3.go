package main

import "fmt"

//const v4 int  // 错误
const v2 = "123" // 正确

// 因式分解
const (
	v1 = 1
	v3 = 3
)

func main() {
	fmt.Println(v1, v2, v3)
}
