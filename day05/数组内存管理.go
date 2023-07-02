package main

import "fmt"

func main() {
	var numbers1 [3]int
	numbers1[1] = 1
	numbers1[2] = 2
	numbers1[0] = 3
	fmt.Printf("数组内存地址：%p \n", &numbers1)       // 数组内存地址：0xc00001e090
	fmt.Printf("数组第一个元素的内存地址0：%p \n", &numbers1[0]) //数组第一个元素的内存地址：0xc00001e090
	fmt.Printf("数组第一个元素的内存地址1：%p \n", &numbers1[1]) //数组第一个元素的内存地址1：0xc00001e098
	fmt.Printf("数组第一个元素的内存地址2：%p \n", &numbers1[2]) //数组第一个元素的内存地址2：0xc00001e0a0


}
