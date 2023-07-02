package main

import "fmt"

func main() {

	var (
		name           = "shdeng"
		age            = 30
		hobby          = "jiuba"
		gender         string // 只声明不赋值，默认值是："" 空字符串
		lenth          int    // 只声明不赋值，默认值是：0
		has_girlfriend bool   //  只声明但不赋值， 默认值是：false
	)
	fmt.Println(name, age, hobby, gender, lenth, has_girlfriend)

}
