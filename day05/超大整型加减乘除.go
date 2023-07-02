package main

import (
	"fmt"
	"math/big"
	"reflect"
)

func main() {

	n1 := new(big.Int)
	n1.SetInt64(1000)
	n2 := new(big.Int)
	n2.SetInt64(33)

	result1 := new(big.Int)
	result1.Add(n1, n2)

	fmt.Println(result1, reflect.TypeOf(result1)) // 3102 *big.Int

	//  NewInt 是big模块对指针类型的缩写
	n3 := big.NewInt(29)

	n4 := big.NewInt(221)
	result2 := new(big.Int)
	result2.Add(n3, n4)
	fmt.Println(result2)

	// 超大整型，加减乘除

	result := new(big.Int)
	result.Add(n1, n2) // 加法,仅限两个元素
	fmt.Println("超大整型加法：", result)

	result.Sub(n1, n2) // 减法,仅限两个元素
	fmt.Println("超大整型减法：", result)

	result.Mul(n1, n2) // 乘法,仅限两个元素
	fmt.Println("超大整型乘法：", result)

	result.Div(n1, n2) // 除法,仅限两个元素
	fmt.Println("超大整型除法：", result)

	minder := new(big.Int) // 余数
	result.DivMod(n1, n2, minder) // 除法求余数,仅限两个元素
	fmt.Println("超大整型取余数：", result,minder)

}
