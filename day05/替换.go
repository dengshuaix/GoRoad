package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "dengzhouzhouzhou"

	result1 := strings.Replace(name, "zhou", "yx", 1)  // 找到zhou替换为yx，从左到右找第一个替换
	result2 := strings.Replace(name, "zhou", "yx", 2)  // 找到zhou替换为yx，从左到右找前两个替换
	result3 := strings.Replace(name, "zhou", "yx", -1) // 找到zhou替换为yx，替换所有

	fmt.Println(result1, result2, result3) //dengyxzhouzhou dengyxyxzhou dengyxyxyx

}