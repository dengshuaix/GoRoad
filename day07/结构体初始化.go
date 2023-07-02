package main

import "fmt"

func main() {
	// 定义一个结构体（类型），每个结构体包含 name、age、hobby 三个元素
	type Address struct {
		city, state string
	}
	type Person struct {
		name    string
		age     int
		Address // 匿名字段，那么默认Person就包含了Address的所有字段
	}

	//方式1：先后顺序
	p1 := Person{"shdeng", 19, Address{"北京", "中国"}}
	fmt.Println(p1.name, p1.age, p1.city, p1.state) // shdeng 19 北京 中国

	//方式2：关键字
	p2 := Person{name: "shdeng", age: 19, Address: Address{city: "北京", state: "中国"}}
	fmt.Println(p2.name, p2.age, p2.city, p2.state, p2.Address.city, p2.Address.state) // shdeng 19 北京 中国 北京 中国

	//方式3：先声明再赋值
	var p3 Person
	p3.name = "shdeng"
	p3.age = 50
	p3.Address = Address{
		city:  "北京",
		state: "BJ",
	}
	fmt.Println(p3.name, p3.age, p3.Address.city, p3.Address.state) // shdeng 50 北京 BJ
	// 或
	var p4 Person
	p4.name = "shdeng"
	p4.age = 50
	p4.city = "北京"
	p4.state = "BJ"
	fmt.Println(p3.name, p3.age, p3.Address.city, p3.Address.state) // shdeng 50 北京 BJ

}
