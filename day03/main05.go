package main

import "fmt"

func main() {
	var name string
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)

	if name == "shdeng" {
		goto SVIP
	} else if name == "deng" {
		goto VIP
	}
// goto 特性，不return 会将一下的 fmt.Println(xxx) 全部执行，增加return 以main函数形式返回结果
	fmt.Println("需要预约。。。")
	return
VIP:
	fmt.Println("须等待。。。。")
	return

SVIP:
	fmt.Println("无须等待，直接进入")
	return

}
