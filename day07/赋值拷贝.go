package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Person struct {
		name string
		age  int
	}

	p1 := Person{name: "shdeng", age: 18}
	p2 := p1 // 内部将p1重新拷贝一份

	fmt.Println(p1, reflect.TypeOf(p1), &p1) // {shdeng 18} main.Person &{shdeng 18}

	fmt.Println(p2, reflect.TypeOf(p2), &p2) //{shdeng 18} main.Person &{shdeng 18}

	p1.name = "zhou"

	fmt.Println(p1, reflect.TypeOf(p1), &p1) // {zhou 18} main.Person &{zhou 18}

	fmt.Println(p2, reflect.TypeOf(p2), &p2) // {shdeng 18} main.Person &{shdeng 18}

}
