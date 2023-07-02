package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var name string = "邓周"
	fmt.Println(len(name)) // 获取字节长度 ， utf8 一个字符占3个字节
	fmt.Println(utf8.RuneCountInString(name))  // 获取字符长度 ， 2个字符
}
