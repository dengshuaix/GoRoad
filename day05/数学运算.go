package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(math.Abs(-19))      // 绝对值
	fmt.Println(math.Floor(3.14))   // 向下取整
	fmt.Println(math.Ceil(3.14))    // 向上取整
	fmt.Println(math.Round(3.3478)) // 就近取整
	fmt.Println(math.Mod(11, 3))    // 取余数，同 11 % 3 =1
	fmt.Println(math.Pow(2, 5))     // 幂次计算
	fmt.Println(math.Pow10(2))      // 计算10次方，2的十次方
	fmt.Println(math.Max(1, 2))     // 比较两个大小，取大
	fmt.Println(math.Min(1, 2))     // 比较两个大小，取小
}
