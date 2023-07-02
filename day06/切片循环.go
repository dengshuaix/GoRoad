package main

import "fmt"

func main() {
	var v1 = new([]int)
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", v1, len(*v1), cap(*v1)) //类型：*[]int，长度：0，容量：0

	var v2 *[]int
	fmt.Println(v2)                       // <nil>
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", v2) //类型：*[]int，长度：%!d(MISSING)，容量：%!d(MISSING)```

}
