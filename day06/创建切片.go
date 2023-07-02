package main

import "fmt"

func main() {
	// 创建切片的三种方式
	//1.直接声明，nums 容量和长度都是0
	var nums []int
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", nums, len(nums), cap(nums)) //类型：[]int，长度：0，容量：0
	nums = append(nums, 1)
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", nums, len(nums), cap(nums)) //类型：[]int，长度：1，容量：1

	nums = append(nums, 2)
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", nums, len(nums), cap(nums)) //类型：[]int，长度：2，容量：2
	nums = append(nums, 3)
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", nums, len(nums), cap(nums)) //类型：[]int，长度：3，容量：4

	// 2. 直接初始化切片数据 , 长度和容量一致
	var data = []int{11, 22, 33}
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", data, len(data), cap(data)) //类型：[]int，长度：3，容量：3

	//3. make 创建：切片，字典，channel
	var users = make([]int, 1, 3)
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", users, len(users), cap(users)) //类型：[]int，长度：1，容量：3
	users = append(users, 1)
	users = append(users, 2)
	users = append(users, 3)
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", users, len(users), cap(users)) //类型：[]int，长度：4，容量：6

	var v1 = new([]int)
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", v1, len(*v1), cap(*v1)) //类型：*[]int，长度：0，容量：0

	var v2 *[]int
	fmt.Println(v2)                       // <nil>
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", v2) //类型：*[]int，长度：%!d(MISSING)，容量：%!d(MISSING)

}
