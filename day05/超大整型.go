package main

import (
	"fmt"
	"math/big"
)

func main() {

	// 变量创建超大整型
	// 第一步创建 超大整型对象
	var v1 big.Int
	var v2 big.Int

	// 第二部，给超大整型赋值
	v1.SetInt64(9223372036854775807)
	v1.SetInt64(92233720368547758089223372036854775808) // 容易产生溢出问题
	fmt.Println(v1)

	v2.SetString("92233720368547758089223372036854775808", 10) // 不会产生溢出问题
	fmt.Println(v2)

	// 指针创建超大整型
	var v3 = new(big.Int)
	var v4 = new(big.Int)

	// 第二部，给超大整型赋值
	v3.SetInt64(9223372036854775807)
	fmt.Println(v1)

	v4.SetString("92233720368547758089223372036854775808", 10) // 不会产生溢出问题
	fmt.Println(v4)

}
