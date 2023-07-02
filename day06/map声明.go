package main

import (
	"fmt"
	"reflect"
)

func main() {
	// map 声明以及初始化，键值对格式，值的格式必须保持一致
	userInfo := map[string]string{"name": "shdeng", "age": "18", "gender": "男"}

	// map获取数据
	fmt.Println("name：", userInfo["name"])     //name： shdeng
	fmt.Println("age：", userInfo["age"])       //age： 18
	fmt.Println("gender：", userInfo["gender"]) //gender： 男

	// 使用 make 初始化 map.  且map不存在容量，内部也有扩容机制
	intMap := make(map[int]int, 10)
	fmt.Println("初始化后的长度为：", len(intMap)) // 0
	fmt.Printf("初始化后的地址：%p\n", &intMap)   // 初始化后的地址：0xc000012030
	intMap[0] = 1
	intMap[1] = 1
	intMap[2] = 1
	for i := 0; i < 100; i++ {
		intMap[i] = i
		fmt.Printf("intMap地址：%p,%v \n", &intMap, i) // intMap地址：0xc000012030 ,(0~99)

	}

	fmt.Println("intMap长度：", len(intMap)) // 100
	fmt.Printf("intMap地址：%p\n", &intMap)  // intMap地址：0xc000012030

	// map 被赋值 , 出现拷贝现象。重新生成一份data数据
	data := make(map[string]int)
	data["1"] = 100
	data["2"] = 200

	var p1 map[string]int
	fmt.Println(p1)                 //map[] 空 map
	fmt.Println(len(p1))            // 长度为0
	fmt.Println(reflect.TypeOf(p1)) //map[string]int
	p1 = data

	fmt.Printf("data 类型：%T---值：%v---内存地址：%p\n", data, data, &data) // data 类型：map[string]int---值：map[1:100 2:200]---内存地址：0xc000012048
	fmt.Printf("p1 类型：%T---值：%v---内存地址：%p\n", p1, p1, &p1)         // p1 类型：map[string]int---值：map[1:100 2:200]---内存地址：0xc000012050

	// 使用 new 声明指针map, new函数会生成空指针map，不能直接map进行数据操作
	v1 := make(map[string]int)
	v1["100"] = 110
	v1["200"] = 220

	v2 := new(map[string]int) // 此处v2时 指针类型的空map
	// v2["k1"] = 123  # 报错
	v2 = &v1 // 将v1，map类型的地址，传入给v2
	fmt.Println(v2)

	// map 的key可哈希类型：int/bool/string/array/float
	v3 := make(map[[2]int][3]int)
	v3[[2]int{1, 2}] = [3]int{1, 2, 3}
	v3[[2]int{4, 5}] = [3]int{3, 3, 3}
	fmt.Println("数组支持 哈希，结果：", v3) //  map[[1 2]:[1 2 3] [4 5]:[3 3 3]]

	v4 := make(map[[2]int]float32)
	v4[[2]int{1, 1}] = 1.6
	v4[[2]int{1, 2}] = 3.4
	fmt.Println("float支持 哈希，结果：", v4) //  map[[1 1]:1.6 [1 2]:3.4]

}
