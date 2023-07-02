package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串 和 “rune 集合”
	var name string = "邓周"
	// 将字符串 转换为unicode编码
	unicodeList := []rune(name)
	fmt.Println("Unicode编码集合：", unicodeList)                                             // Unicode编码集合： [37011 21608]
	fmt.Println(unicodeList[0], "转换为16进制", strconv.FormatInt(int64(unicodeList[0]), 16)) // 37011 转换为16进制 9093
	fmt.Println(unicodeList[1], "转换为16进制", strconv.FormatInt(int64(unicodeList[1]), 16)) // 21608 转换为16进制 5468

	// rune 集合转换为字符串
	unicodeList2 := []rune{37011, 21608}
	runeToStrResult := string(unicodeList2)
	fmt.Println(runeToStrResult) //邓周
}
