package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//整型之间的转换
	var v1 int8 = 10
	var v2 int16 = 120
	v3 := int16(v1) + v2
	fmt.Println(v3, reflect.TypeOf(v3)) // 130 int16

	var v4 int8 = 10
	var v5 int64 = 2022
	v6 := int64(v4)
	v7 := int8(v5)
	fmt.Println("低位转向高位：", v6) // 低位转向高位： 10
	fmt.Println("高位转向低位：", v7) // 高位转向低位：  -26 , 超过127 就会重新开始从-128开始循环

	//整型与字符串的转换 Itoa 和  Atoi
	a1 := 19
	strRes := strconv.Itoa(a1)
	fmt.Println("数字转字符串：", strRes) //数字转字符串： 19

	a2 := "666"
	intRes, _ := strconv.Atoi(a2)
	fmt.Println("字符串转数字", intRes) //字符串转数字 666

	// 整型进制转换 （默认整型就是10进制）
	// - go语言中，整型就是十进制
	// - 其他进制就是 以字符串的形式存在
	q1 := 99
	// base 是进制种类：2，8，10，16
	hexResult := strconv.FormatInt(int64(q1), 2)
	fmt.Println("整型进制转换：", hexResult, reflect.TypeOf(hexResult)) // 整型进制转换： 1100011 string

	// 字符串转数值类型
	q2 := "111001"
	strResHex, err := strconv.ParseInt(q2, 10, 0)
	fmt.Println("2进制转10进制", strResHex, "错误：", err, "类型：", reflect.TypeOf(strResHex)) //2进制转10进制 111001 错误： <nil> 类型： int64

	// 转其他进制 ，只能低转高，不能高转低。
	q3 := "270f" // 16进制
	strResHexQ3, err := strconv.ParseInt(q3, 16, 0)
	fmt.Println("2进制转16进制", strResHexQ3, "错误：", err, "类型：", reflect.TypeOf(strResHexQ3)) // 16进制字符串转10进制 9999 错误： <nil> 类型： int64

}
