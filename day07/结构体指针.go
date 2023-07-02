package main

import "fmt"

func main() {

	type Person struct {
		name string
		age  int
	}
	// 初始化指针类型结构体
	// 方式一
	people := Person{"shdeng", 18}
	fmt.Println(people.name, people.age)
	p2 := &people
	fmt.Printf("p2类型：%T , p2地址：%p\n", p2, &p2) // p2类型：*main.Person , p2地址：0xc000012030

	// 方式二
	p3 := &Person{"zhou", 25}
	fmt.Printf("p3类型：%T , p3地址：%p\n", p3, &p3) // p3类型：*main.Person , p3地址：0xc000012038

	// 方式三
	var p4 *Person = new(Person)
	p4.name = "dz"
	p4.age = 5

	fmt.Println(p4.name, p4.age)
	fmt.Printf("p4类型：%T , p4地址：%p\n", p4, &p4) // p3类型：*main.Person ,p4地址：0xc000012040

}
