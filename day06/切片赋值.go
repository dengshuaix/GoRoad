package main

import "fmt"

func main() {
	v1 := []int{6, 9}
	v2 := v1                                  // 出现拷贝现象，但是内部的 ptr （unsafe.Pointer）指向同一块内存地址
	fmt.Println(v1, v2)                       // [6 9] [6 9]
	fmt.Printf("v1的内存地址：%p \n", &v1)          // v1的内存地址：0xc000010030
	fmt.Printf("v2的内存地址：%p \n", &v2)          // v2的内存地址：0xc000010048
	fmt.Printf("v1下标为0数据的内存地址：%p \n", &v1[0]) // v1下标为0数据的内存地址：0xc00001c070
	fmt.Printf("v2下标为0数据的内存地址：%p \n", &v2[0]) // v2下标为0数据的内存地址：0xc00001c070
	fmt.Printf("v1下标为1数据的内存地址：%p \n", &v1[1]) // v1下标为1数据的内存地址：0xc00001c078
	fmt.Printf("v2下标为1数据的内存地址：%p \n", &v2[1]) // v2下标为1数据的内存地址：0xc00001c078

	v1[0] = 11111
	fmt.Println(v1, v2) // [11111 9] [11111 9]
}
