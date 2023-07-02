package main

import "fmt"

func other(a1 int, a2 int) {
	fmt.Println("defer函数被执行了")
}

func do() int {
	fmt.Println("风吹")
	defer fmt.Println("函数执行完毕了")
	defer other(1, 22)
	fmt.Println("屁屁凉")
	return 666
}


func main() {
	ret := do()
	fmt.Println(ret)
	// 风吹
	// 屁屁凉
	// defer函数被执行了
	// 函数执行完毕了
}