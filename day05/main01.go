package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	r1 := 100

	// 数字转字符串
	strR1 := strconv.Itoa(r1)
	fmt.Printf("数字转字符串%s\n", strR1, "结果类型：", reflect.TypeOf(strR1), reflect.TypeOf(reflect.TypeOf(strR1)))
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	r2 := "100"
	intR2, err := strconv.Atoi(r2)
	if err == nil {
		fmt.Printf("转换成功：%d \n", intR2, "结果类型：", reflect.TypeOf(intR2))
	} else {
		fmt.Println("转换失败！")
	}

}
