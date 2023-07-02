package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}

	// 初始化结构体
	p1 := Person{"shdeng", 18}
	fmt.Println(p1.name, p1.age)

	// 初始化结构体指针
	p2 := &Person{"shdeng", 1} // p2 指向 结构体的内存地址
	fmt.Println(p2.name, p2.age)

}
