package main

import "fmt"

type Person struct {
	name string
	age  int
}

var P Person = Person{"shdeng", 18}

func doSomething() *Person {
	fmt.Printf("%p\n", &P) // 0x113d230
	return &P              // 返回结构体地址
}

func main() {

	data := doSomething()
	P.name = "zhou"
	fmt.Println(data)         // &{zhou 18}
	fmt.Printf("%p\n", &data) // 0xc000012028

	fmt.Println(P)         // {zhou 18}
	fmt.Printf("%p\n", &P) // 0x113d230

}
