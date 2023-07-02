package main

import "fmt"

/*
	type 结构体名称 struct {
		字段 类型
		...
}
*/

type Address struct {
	city, state string
}

// 定义结构体
type Person struct {
	name        string
	age         int
	email       string
	sex, idCard string // 同一类型的字段，可以这个写
	hobby       []string
	Address     // 匿名结构体 ， 等同于 Address Address
}

func main() {
	// 初始化结构体
	var p1 = Person{"deng", 18, "dengshuai37@163.com", "nan", "000",
		[]string{"1", "2"}, Address{"bj", "cp"}}

	// 获取结构体中的数据
	fmt.Println(p1.name)          // deng
	fmt.Println(p1.age)           // 18
	fmt.Println(p1.email)         // dengshuai37@163.com
	fmt.Println(p1.sex)           // nan
	fmt.Println(p1.idCard)        // 000
	fmt.Println(p1.city)          // bj 等同于 p1.Address.city
	fmt.Println(p1.Address.city)  // bj
	fmt.Println(p1.state)         // cp
	fmt.Println(p1.Address.state) // cp 等同于 p1.Address.state

	// 修改 结构体数据
	p1.age = 20
	fmt.Println(p1.age) // 20
}
