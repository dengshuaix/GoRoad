package main

import "fmt"

func main() {
	// 节约内存空间
	v1 := "shdeng"
	v2 := &v1
	fmt.Println(v1, v2, *v2) // shdeng 0xc000014270 shdeng
	v1 = "zhou" // v2 指向 v1的内存地址。 修改v1 相当于修改v2
	fmt.Println(v1, v2, *v2) // zhou 0xc000014270 zhou
}
