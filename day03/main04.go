package main

import "fmt"

func main() {

	// go 循环之打标签 break/continue
f1:
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j == 4 {
				fmt.Println("跳转到指定标签")
				continue f1 // 跳出当前循环，执行标签位置语句
			}
			fmt.Printf("i:%d,j:%d\n", i, j)
		}
	}
fmt.Println("")
f2:
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j == 4 {
				fmt.Println("跳转到指定标签")
				break f2 // 直接中断
			}
			fmt.Printf("i:%d,j:%d\n", i, j)
		}
	}
}
