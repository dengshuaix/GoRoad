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

	p1 := &Person{"shdeng", 18} // 指针类型结构体
	p2 := p1                    // 指针类型结构体

	fmt.Println(p1, reflect.TypeOf(p1), &p1) // &{shdeng 18} *main.Person 0xc000012028

	fmt.Println(p2, reflect.TypeOf(p2), &p2) // &{shdeng 18} *main.Person 0xc000012030

	p1.name = "zhou"
	fmt.Println(p1, reflect.TypeOf(p1), &p1) // &{zhou 18} *main.Person 0xc000012028

	fmt.Println(p2, reflect.TypeOf(p2), &p2) // &{zhou 18} *main.Person 0xc000012030

}
