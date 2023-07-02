package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode/utf8"
)

func main() {

	// unicode字符集：文字 -> 码点（ucs4, 4个字节表示）
	// utf-8编码，对unicode字符集的码点进行编码最终得到：1000100001

	//1。 本质是utf-8编码的序列
	var name string = "指染未来"
	fmt.Println(name[0], "字符串转2进制：", strconv.FormatInt(int64(name[0]), 2), "数据类型：", reflect.TypeOf(name[0]))
	fmt.Println(name[1], "字符串转2进制：", strconv.FormatInt(int64(name[2]), 2), "数据类型：", reflect.TypeOf(name[1]))
	fmt.Println(name[2], "字符串转2进制：", strconv.FormatInt(int64(name[2]), 2), "数据类型：", reflect.TypeOf(name[2]))
	fmt.Println(name[3], "字符串转2进制：", strconv.FormatInt(int64(name[3]), 2), "数据类型：", reflect.TypeOf(name[3]))
	fmt.Println(name[4], "字符串转2进制：", strconv.FormatInt(int64(name[4]), 2), "数据类型：", reflect.TypeOf(name[4]))
	fmt.Println(name[5], "字符串转2进制：", strconv.FormatInt(int64(name[5]), 2), "数据类型：", reflect.TypeOf(name[5]))
	fmt.Println(name[6], "字符串转2进制：", strconv.FormatInt(int64(name[6]), 2), "数据类型：", reflect.TypeOf(name[6]))
	fmt.Println(name[7], "字符串转2进制：", strconv.FormatInt(int64(name[7]), 2), "数据类型：", reflect.TypeOf(name[7]))
	fmt.Println(name[8], "字符串转2进制：", strconv.FormatInt(int64(name[8]), 2), "数据类型：", reflect.TypeOf(name[8]))
	fmt.Println(name[9], "字符串转2进制：", strconv.FormatInt(int64(name[9]), 2), "数据类型：", reflect.TypeOf(name[9]))
	fmt.Println(name[10], "字符串转2进制：", strconv.FormatInt(int64(name[10]), 2), "数据类型：", reflect.TypeOf(name[10]))
	fmt.Println(name[11], "字符串转2进制：", strconv.FormatInt(int64(name[11]), 2), "数据类型：", reflect.TypeOf(name[11]))
	/*
		// 上面的结果
		230 字符串转2进制： 11100110 数据类型： uint8
		140 字符串转2进制： 10000111 数据类型： uint8
		135 字符串转2进制： 10000111 数据类型： uint8
		230 字符串转2进制： 11100110 数据类型： uint8
		159 字符串转2进制： 10011111 数据类型： uint8
		147 字符串转2进制： 10010011 数据类型： uint8
		230 字符串转2进制： 11100110 数据类型： uint8
		156 字符串转2进制： 10011100 数据类型： uint8
		170 字符串转2进制： 10101010 数据类型： uint8
		230 字符串转2进制： 11100110 数据类型： uint8
		157 字符串转2进制： 10011101 数据类型： uint8
		165 字符串转2进制： 10100101 数据类型： uint8
	*/

	// 2。 获取字符串长度 (字节长度)
	fmt.Println(len(name)) // 12

	//3. 将字符串转换成 一个 "字节集合"
	byteList := []byte(name)
	fmt.Println(byteList) // [230 140 135 230 159 147 230 156 170 230 157 165]

	//4。将`字节集合`转换成字符串
	byteList2 := []byte{230, 140, 135, 230, 159, 147, 230, 156, 170, 230, 157, 165}
	byteToStr1 := string(byteList)
	byteToStr2 := string(byteList2)
	fmt.Println(byteToStr1, byteToStr2) // 指染未来 指染未来

	// 5. 将  字符串 转换为 unicode 字符集码典 集合, 使用rune函数
	unicodeList := []rune(name)
	fmt.Println(unicodeList)                                                 // [25351 26579 26410 26469] 这是 十进制的数据
	fmt.Println(unicodeList[0], strconv.FormatInt(int64(unicodeList[0]), 2)) //25351 110001100000111
	fmt.Println(unicodeList[1], strconv.FormatInt(int64(unicodeList[1]), 2)) //26579 110011111010011
	fmt.Println(unicodeList[2], strconv.FormatInt(int64(unicodeList[2]), 2)) //26410 110011100101010
	fmt.Println(unicodeList[3], strconv.FormatInt(int64(unicodeList[3]), 2)) //26469 110011101100101

	// 6. rune 集合 转换成字符串
	runeList := []rune{25351, 26579, 26410, 26469}
	runeToStr := string(runeList)
	fmt.Println(runeToStr) //指染未来

	// 7. 字符串，以字符形式获取长度
	stringLength := utf8.RuneCountInString(name)
	fmt.Println(stringLength) //4
}
