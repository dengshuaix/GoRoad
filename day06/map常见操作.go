package main

import "fmt"

func main() {

	// 1. map 长度 ,
	data := map[string]string{"1": "2", "3": "4"}
	fmt.Println("map长度：", len(data)) // map长度： 2

	data2 := make(map[string]string, 10) // 10 是参考值，根据10会重新计算的到新的容量
	data2["s"] = "1"
	data2["w"] = "2"
	data2["q"] = "2"
	fmt.Println("map长度：", len(data2)) // map长度： 3

	//cap(data2)  map不能使用cap查看容量会报错

	// 2. 添加
	data3 := map[string]string{"n1": "s", "n2": "2"}
	data3["n3"] = "3"
	fmt.Println(data3) //map[n1:s n2:2 n3:3]

	//3. 修改
	data4 := map[string]string{"n1": "s", "n2": "2"}
	fmt.Println(data4) //map[n1:s n2:2]
	data4["n1"] = "1"
	fmt.Println(data4) //map[n1:1 n2:2]

	//4. 删除
	data5 := map[string]string{"n1": "s", "n2": "2"}
	fmt.Println(data5) // map[n1:s n2:2]
	delete(data5, "n1")
	delete(data5, "n3") // 删除不存在的键时，不会报错
	fmt.Println(data5)  // map[n2:2]

	//5. 查看
	data6 := map[string]string{"n1": "a", "n2": "1"}
	fmt.Println("直接通过键名取值：", data6["n1"]) //直接通过键名取值： a

	data7 := map[string]string{"n1": "a", "n2": "1"}
	for key, val := range data7 {
		fmt.Println("通过循环取键和值：", key, val) // 通过循环取键和值： n1 a ,通过循环取键和值： n2 1
	}

	data8 := map[string]string{"n1": "a", "n2": "1"}
	for key, _ := range data8 {
		fmt.Println("通过循环取key：", key) // 通过循环取key： n1 ,通过循环取key： n2 1
	}

	data9 := map[string]string{"n1": "a", "n2": "1"}
	for _, val := range data9 {
		fmt.Println("通过循环取值：", val) // 通过循环取值：a ,通过循环取值：1
	}

	// 6. 嵌套
	v1 := make(map[string]string)
	v2 := make(map[string]int)
	v3 := make(map[string]bool)
	v4 := make(map[string]float64)

	v5 := make(map[string][2]int)           // key是字符串，value是 2个数据数组
	v6 := make(map[string][]int)            // key是字符串，value是  动态切片
	v7 := make(map[string]map[int]string)   // key 是 map ，值是 map（int为key，string为value）
	fmt.Println(v1, v2, v3, v4, v5, v6, v7) // map[] map[] map[] map[] map[] map[] map[]

	v8 := make(map[string][2]map[string]string) // key是字符串，value是2个数组嵌套的map
	v8["n1"] = [2]map[string]string{{"s1": "1", "1": "s"}, {"s2": "2"}}
	v8["n2"] = [2]map[string]string{{"s3": "3"}, {"s4": "4"}}
	fmt.Println(v8) //map[n1:[map[1:s s1:1] map[s2:2]] n2:[map[s3:3] map[s4:4]]]

	//7.哪些类型支持map的key
	v17 := make(map[int]int)     //支持
	v9 := make(map[string]int)   // 支持
	v10 := make(map[float32]int) // 支持
	v11 := make(map[bool]int)    // 支持
	v12 := make(map[[2]int]int)  //支持
	//v13 := make(map[ []int ]int) // 错误,不可哈希
	//v14 := make(map[ map[int]int ]int) // 错误，不可哈希
	//v15 := make(map[ [2][]int ]int) // 报错
	//v16 := make(map[ [2]map[string]string ]int) // 报错
	fmt.Println(v17, v9, v10, v11, v12) // map[] map[] map[] map[] map[]

	//8. map 变量赋值
	v18 := map[string]string{"n1": "shdeng", "n2": "zhou"}
	v19 := v18

	v18["n1"] = "yx"

	fmt.Println(v18) // map[n1:yx n2:zhou]
	fmt.Println(v19) // map[n1:yx n2:zhou]

}
