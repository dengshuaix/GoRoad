package main

import "fmt"

func main() {

	// 长度和容量
	v1 := []int{11, 22, 33}
	fmt.Println("长度：", len(v1), "容量：", cap(v1)) // 长度： 3 容量： 3

	// 索引取值
	v2 := []string{"a1", "a2", "a3"}
	fmt.Println("索引位置：0", "值:", v2[0]) // 索引位置：0 值: a1

	fmt.Println("索引位置：1", "值:", v2[1]) //索引位置：1 值: a2

	fmt.Println("索引位置：2", "值:", v2[2]) // 索引位置：2 值: a3

	v3 := make([]int, 2, 5)
	fmt.Println("索引位置：0", "值:", v3[0]) // 索引位置：0 值: 0
	fmt.Println("索引位置：1", "值:", v3[1]) // 索引位置：1 值: 0
	//fmt.Println("索引位置：2", "值:", v3[2]) // 会报错,因为长度不够

	//切片的切片，遵循开区间 （左边能取到，右边取不到。 不存在负索引）
	v4 := []int{11, 22, 33, 44, 55, 66}
	v5 := v4[1:3]
	v6 := v4[1:]
	v7 := v4[:3]
	fmt.Println("切片索引位置1到2：", v5)  // 切片索引位置1到2： [22 33]
	fmt.Println("切片索引位置1到最后：", v6) // 切片索引位置1到最后： [22 33 44 55 66]
	fmt.Println("切片索引位置0到2：", v7)  // 切片索引位置0到2： [11 22 33]

	// 切片追加
	v8 := []int{11, 22, 33}
	v9 := append(v8, 44)
	v10 := append(v8, 55, 66, 77, 88)
	v11 := append(v1, []int{100, 200, 300}...)    // 批量追加， ...表示可变参数，被称为 "variadic" 语法。
	fmt.Printf("追加一个数据:%v，数据地址:%p \n", v9, &v9)   // 自动扩容 追加一个数据:[11 22 33 44]，数据地址:0xc000010078
	fmt.Printf("追加四个数据:%v，数据地址:%p \n", v10, &v10) // 自动扩容 追加四个数据:[11 22 33 55 66 77 88]，数据地址:0xc000010090
	fmt.Printf("追加一堆数据:%v，数据地址:%p \n", v11, &v11) // 自动扩容 追加一堆数据:[11 22 33 100 200 300]，数据地址:0xc0000100a8

	// 切片删除 (使用到append追加 和 ...表示可变参数variadic) todo ： 效率低下，之后会使用链表删除
	v12 := []int{11, 22, 33, 44, 55, 66}
	deleteIndex := 2 // 待删除的索引下标
	result := append(v12[:deleteIndex], v12[deleteIndex+1:]...)
	fmt.Println(result) // [11 22 44 55 66]

	// 插入
	v13 := []int{11, 22, 33, 44, 55, 66}
	insertIndex := 3 // 在索引位置
	result2 := make([]int, 0, len(v13)+1)
	result2 = append(result2, v13[:insertIndex]...)
	result = append(result2, 99) // 此处为插入数据
	result2 = append(result2, v13[insertIndex:]...)
	fmt.Println(result2)

	// 循环
	v14 := []int{11, 22, 33, 44, 55, 66}
	for i := 0; i < len(v14); i++ {
		fmt.Println("索引：", i, "值：", v14[i])
	}
	for index, value := range v14 {
		fmt.Println("索引：", index, "值：", value)
	}

}
