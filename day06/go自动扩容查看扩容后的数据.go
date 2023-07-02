package main

import "fmt"

func main() {

	v1 := make([]int, 2, 2)
	fmt.Println("v1扩容前:", v1)
	v2 := append(v1, 99)
	fmt.Println("v1扩容后:", v1)
	fmt.Println("v2扩容后:", v2) // 双倍扩容
	fmt.Println("v2扩容后长度:", len(v2), "容量:", cap(v2))
	fmt.Println("v2 扩容后最后一个长度数据是:", v2[2])
	fmt.Println("v2 扩容后最后一个容量数据是:", v2[3]) // 未赋值

}
