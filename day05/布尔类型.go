package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 字符串转换布尔类型
	// true: "1", "t", "T", "true", "TRUE", "True"
	// false:" 0", "f", "F", "false", "FALSE", "False"
	v1, err := strconv.ParseBool("1")
	fmt.Println(v1, "数据类型:", reflect.TypeOf(v1), "错误信息：", err) // true 数据类型: bool 错误信息： <nil>

	// 布尔类型转换字符串
	v2 := strconv.FormatBool(false)
	fmt.Println(v2, "数据类型:", reflect.TypeOf(v2)) // false 数据类型: string

}
