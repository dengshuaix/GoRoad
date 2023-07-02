package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 定义字符串，字符串是以什么形式存在于Go编译器（utf-8编码）
	name := "邓周"

	// 其他语言：
	// 0 邓
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2)) // 233 11101001
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2)) // 130 10000010
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2)) // 147 10010011

	// 1 周
	fmt.Println(name[3])
	fmt.Println(name[4])
	fmt.Println(name[5])

}
