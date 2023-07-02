package main

import "fmt"

func main() {
	// 声明一个 字符串类型 的变量（默认初始化值为空字符串）。
	var v1 string
	// 声明一个 字符串的指针类型 的变量（默认初始化值为nil）。
	var v2 *string
	var v3 int
	var v4 *int
	fmt.Println(v1) // 空值
	fmt.Println(v2) // <nil>
	fmt.Println(v3) // 0
	fmt.Println(v4) // <nil>

	// 声明一个 字符串类型 的变量，值为 shdeng。
	var name string = "shdeng"
	// 声明一个 字符串的指针类型 的变量，值为 name 对应的内存地址。
	var pointer *string = &name
	fmt.Println(name)    // shdeng
	fmt.Println(pointer) // 0xc000014270
	var age int = 18
	var x1 *int = &age
	fmt.Println(age, x1) // 18 0xc0000ac020

}
