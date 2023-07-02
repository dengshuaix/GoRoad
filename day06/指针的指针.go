package main

import "fmt"

func main() {
	// 指针嵌套
	name := "shdeng"

	// 声明一个指针类型变量p1，内部存储name的内存地址
	var p1 *string = &name

	// 声明一个指针的指针类型变量p2，内部存储指针p1的内存地址
	var p2 **string = &p1

	// 声明一个指针的指针的指针类型变量p3，内部存储指针p2的内存地址
	var p3 ***string = &p2

	//改变name值
	***p3="zhou"
	fmt.Println(name)
}
