package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 真true 循环 ,死循环
	for {
		fmt.Println("1234")
		time.Sleep(time.Second * 1)
	}

	// 2. 条件循环，条件判断也是  布尔结果
	for 2 > 1 {
		fmt.Println("1234")
		time.Sleep(time.Second * 1)
	}
	// 3. 变量判断
	number := 1
	for number < 4 {
		fmt.Println("1234", &number)
		time.Sleep(time.Second * 1)
		number += 1         // 重新赋值
		number = number + 1 // 重新赋值
	}
}
