package main

import (
	"fmt"
)

func main() {
	var v1 float32
	v1 = 3.14
	v2 := 99.9
	v3 := float64(v1) + v2 // 同类型数据，才能做处理

	fmt.Println(v1, v2, v3) // 3.14 99.9 103.04000010490418

	// 上述  103.04000010490418 结果，就是非精确方式 ，在举例
	v4 := 0.1
	v5 := 0.2
	result1 := v4 + v5
	fmt.Println(result1) // 0.30000000000000004

	v6 := 0.3
	v7 := 0.2
	result2 := v6 + v7
	fmt.Println(result2) // 0.5
}
