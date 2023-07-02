package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//数组的地址 == 数组的第一个元素的地址。
	dataList := [3]int8{1, 2, 3}
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &dataList, dataList)       // 内存地址：0xc00001c062 - 数据类型：[3]int8 是数组类型
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &dataList[0], dataList[0]) // 内存地址：0xc00001c062 - 数据类型：int8 元素类型
	fmt.Println()
	// 指针的计算
	// 1. 获取数组第一个元素的地址
	var firstDataPtr *int8 = &dataList[0]
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &firstDataPtr, firstDataPtr)

	// 2. 转换成 Pointer 类型
	ptr := unsafe.Pointer(firstDataPtr)
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &ptr, ptr)

	// 3. 转换成 uinptr ， 然后进行地址计算（地址加上1个字节，等同于取出第二个索引位置的数据）
	targetAddress := uintptr(ptr) + 1
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &targetAddress, targetAddress)

	// 4. 根据新地址，重新转换成Pointer类型
	newPtr := unsafe.Pointer(targetAddress)
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &newPtr, newPtr)

	// 5. Pointer对象转换成 int8 指针类型
	newValue := (*int8)(newPtr) // 类型强制转换
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &newValue, newValue)

	// 6. 根据 指针获取值
	fmt.Println(*newValue)
	/*
		// 指针计算示意图
		内存地址：0xc000012030 - 数据类型：*int8
		内存地址：0xc000012038 - 数据类型：unsafe.Pointer
		内存地址：0xc00001c068 - 数据类型：uintptr
		内存地址：0xc000012040 - 数据类型：unsafe.Pointer
		内存地址：0xc000012048 - 数据类型：*int8
		2
	*/
}
