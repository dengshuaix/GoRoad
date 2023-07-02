package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	r1 := 10
	// 十进制 整型 转换成 字符串类型
	strR1 := strconv.FormatInt(int64(r1), 2)
	fmt.Println(strR1)

	// 字符串 转换成 int类型，可以指定进制级别，但结果是以 int64 表达
	r2 := "1001000101"
	// base : 2 表示以 二进制 形式转换成 十进制
	// base : 8 表示以 二进制 形式转换成 十进制
	// base : 16 表示以 二进制 形式转换成 十进制
	// bitSize 表示，转换的过程中对结果进行约束
	intR2, err := strconv.ParseInt(r2, 8, 0)
	if err == nil {
		fmt.Printf("字符串 r2 转换 为2进制：%d\n", intR2, "结果类型：", reflect.TypeOf(intR2))
	} else {
		fmt.Printf("转换失败，%s\n", err)
	}

}
