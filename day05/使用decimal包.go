package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	var v1 = decimal.NewFromFloat(222.321321)
	var v2 = decimal.NewFromFloat(0.321321)

	var v3 = v1.Add(v2)
	var v4 = v1.Sub(v2)
	var v5 = v1.Mul(v2)
	var v6 = v1.Div(v2)
	var v7 = v1.DivRound(v2, 3) // 会进行四舍五入
	fmt.Println(v1)             //222.321321
	fmt.Println(v2)             //0.321321
	fmt.Println(v3)             //222.642642
	fmt.Println(v4)             //222
	fmt.Println(v5)             //71.436509185041
	fmt.Println(v6)             //691.8978871595694026
	fmt.Println(v7)             //691.9
}
