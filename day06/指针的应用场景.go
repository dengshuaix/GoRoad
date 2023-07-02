package main

import "fmt"

func main() {
	var username string
	fmt.Printf("请输入用户名：")

	// 提示用户输入，当用户输入之后，将输入的值赋值给内存地址对应的区域中。
	fmt.Scanf("%s", &username)

	if username == "shdeng" {
		fmt.Println("登录成功")
	} else {
		fmt.Println("登录失败")
	}
}