package main

import "fmt"

func main() {

	// 一维数组
	oneArray := [...]int{1, 2, 3, 4}
	fmt.Println("一维数组：", oneArray)
	// 二维数组
	twoArray := [2][3]int{{1, 2, 3}, {11, 22, 33}}
	fmt.Println("二维数组：", twoArray)
	// 三维数组
	threeArray := [2][2][3]int{{{1, 2, 3}, {11, 22, 33}}, {{1, 2, 3}, {11, 22, 33}}}
	fmt.Println("三维数组：", threeArray)
}
