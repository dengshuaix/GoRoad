package main

import "fmt"

func main() {
	type Address struct {
		city, state string
	}

	type Person struct {
		name    string
		age     int
		address Address
	}

	p1 := Person{name: "二狗子", age: 19, address: Address{"北京", "BJ"}}
	p2 := p1

	fmt.Println(p1.address, &p1.address) // {{北京 BJ} &{北京 BJ}

	fmt.Println(p2.address, &p2.address) // {北京 BJ} &{北京 BJ}

	p1.address.city = "上海"

	fmt.Println(p1.address, &p1.address) // {上海 BJ} &{上海 BJ}

	fmt.Println(p2.address, &p2.address) // {北京 BJ} &{北京 BJ}

}
