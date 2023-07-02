package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	v1 := string(65)
	fmt.Println("字符：", v1)

	v2 := string(37001)
	fmt.Println("字符", v2)

	// 字符串 转数字
	v3, size := utf8.DecodeRuneInString("邓")
	fmt.Println("v3：", v3, "长度：", size) //v3： 37011 长度： 3


}
