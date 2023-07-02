# day 06 数据类型

### 常见数据类型

- 切片Slice，用于`表示`多个数据 （数据集合）

### 切片

- 切片：动态数组
- 特点：
    - 每个切片内部维护：数组指针，切片长度，切片容量三个数据
- 扩容机制

```
1. 当切片中追加的数据 大于 容量数时 ，内部会每次自动扩充当前容量的两倍
2. 当容量超过1024时，每次扩容，只增加1/4的容量数
```

- 切片源代码

```goregexp
type slice struct {
    array unsafe.Pointer // 不安全的指针
    len   int
    cap   int
}
```

![切片-动态数组](image-20200617213213399.png)

###### 创建切片

- 创建常规切片

```go
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

}
```

- 创建指针类型的切片

```go
package main

import "fmt"

func main() {
	var v1 = new([]int)
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", v1, len(*v1), cap(*v1)) //类型：*[]int，长度：0，容量：0

	var v2 *[]int
	fmt.Println(v2)                       // <nil>
	fmt.Printf("类型：%T，长度：%d，容量：%d\n", v2) //类型：*[]int，长度：%!d(MISSING)，容量：%!d(MISSING)```

}
```

###### 自动扩容

```go
package main

import "fmt"

func main() {
	v1 := make([]int, 1, 3)                              // 长度为1，初始化索引为1的数据是int类型的默认值：0
	fmt.Println("扩容前的数据：", v1)                           // 扩容前的数据： [0]
	fmt.Println("扩容前--", "长度:", len(v1), "容量:", cap(v1)) // 扩容前-- 长度: 1 容量: 3
	v1 = append(v1, 2, 3, 4)
	fmt.Println("扩容后的数据：", v1)                           // 扩容后的数据： [0 2 3 4]
	fmt.Println("扩容后--", "长度:", len(v1), "容量:", cap(v1)) //扩容后-- 长度: 4 容量: 6

}
```

- 扩容前

```go
package main

import "fmt"

func main() {

	v1 := []int{11, 22, 33} // 长度为3，容量为3
	v2 := append(v1, 44)    // 超出v1数据的容量，自动扩容

	v1[0] = 999 // 修改v1数据，不会影响v2的数据

	fmt.Printf("%p , %v\n", &v1, v1) // 扩容前数据值：0xc000010030, [999 22 33]
	fmt.Printf("%p , %v\n", &v2, v2) // 扩容后数据值： 0xc000010048, [11 22 33 44]

}

```

###### 切片常见操作

- 获取长度和容量

```go
package main

import "fmt"

func main() {
	// 长度和容量
	v1 := []int{11, 22, 33}
	fmt.Println("长度：", len(v1), "容量：", cap(v1)) // 长度： 3 容量： 3
}
```  

- 索引

```go
package main

import "fmt"

func main() {
	// 索引取值
	v2 := []string{"a1", "a2", "a3"}
	fmt.Println("索引位置：0", "值:", v2[0]) // 索引位置：0 值: a1

	fmt.Println("索引位置：1", "值:", v2[1]) //索引位置：1 值: a2

	fmt.Println("索引位置：2", "值:", v2[2]) // 索引位置：2 值: a3

	v3 := make([]int, 2, 5)
	fmt.Println("索引位置：0", "值:", v3[0]) // 索引位置：0 值: 0
	fmt.Println("索引位置：1", "值:", v3[1]) // 索引位置：1 值: 0
	//fmt.Println("索引位置：2", "值:", v3[2]) // 会报错,因为长度不够
}

```  

- 切片 [:]

```go
package main

import "fmt"

func main() {
	//切片的切片，遵循开区间 （左边能取到，右边取不到。 不存在负索引）
	v4 := []int{11, 22, 33, 44, 55, 66}
	v5 := v4[1:3]
	v6 := v4[1:]
	v7 := v4[:3]
	fmt.Println("切片索引位置1到2：", v5)  // 切片索引位置1到2： [22 33]
	fmt.Println("切片索引位置1到最后：", v6) // 切片索引位置1到最后： [22 33 44 55 66]
	fmt.Println("切片索引位置0到2：", v7)  // 切片索引位置0到2： [11 22 33]
}


```  

- 追加

```go
package main

import "fmt"

func main() {
	// 切片追加
	v8 := []int{11, 22, 33}
	v9 := append(v8, 44)
	v10 := append(v8, 55, 66, 77, 88)
	v11 := append(v1, []int{100, 200, 300}...)    // 批量追加， ...表示可变参数，被称为 "variadic" 语法。
	fmt.Printf("追加一个数据:%v，数据地址:%p \n", v9, &v9)   // 自动扩容 追加一个数据:[11 22 33 44]，数据地址:0xc000010078
	fmt.Printf("追加四个数据:%v，数据地址:%p \n", v10, &v10) // 自动扩容 追加四个数据:[11 22 33 55 66 77 88]，数据地址:0xc000010090
	fmt.Printf("追加一堆数据:%v，数据地址:%p \n", v11, &v11) // 自动扩容 追加一堆数据:[11 22 33 100 200 300]，数据地址:0xc0000100a8

}

```  

- 删除

```go
package main

import "fmt"

func main() {
	// 切片删除 (使用到append追加 和 ...表示可变参数variadic) todo ： 效率低下，之后会使用链表删除
	v12 := []int{11, 22, 33, 44, 55, 66}
	deleteIndex := 2 // 待删除的索引下标
	result := append(v12[:deleteIndex], v12[deleteIndex+1:]...)
	fmt.Println(result) // [11 22 44 55 66]
}

```  

- 插入

```go
package main

import "fmt"

func main() {
	// 插入
	v13 := []int{11, 22, 33, 44, 55, 66}
	insertIndex := 3 // 在索引位置
	result2 := make([]int, 0, len(v13)+1)
	result2 = append(result2, v13[:insertIndex]...)
	result2 = append(result2, 99) // 此处为插入数据
	result2 = append(result2, v13[insertIndex:]...)
	fmt.Println(result2) // [11 22 33 99 44 55 66]

}

```  

- 循环

```go
package main

import "fmt"

func main() {
	// 循环
	v14 := []int{11, 22, 33, 44, 55, 66}
	for i := 0; i < len(v14); i++ {
		fmt.Println("索引：", i, "值：", v14[i])
	}
	for index, value := range v14 {
		fmt.Println("索引：", index, "值：", value)
	}
}

```

###### 切片嵌套

```go
package main

import "fmt"

func main() {

	// 切片嵌套
	// 一维列表
	v1 := []int{11, 22, 33, 44, 55, 66, 77}
	// 二维列表
	v2 := [][]int{[]int{11, 22, 33, 44, 55}, []int{66, 77, 88, 99}}
	// 切片嵌套数组(动态数组，嵌套2个值数组)
	v3 := [][2]int{[2]int{1, 2}, [2]int{3, 4}, [2]int{3, 4}}
	v4 := [][2]int{[2]int{1, 2}, [2]int{3, 4}, [2]int{3, 4}}
	fmt.Println(v1) // [11 22 33 44 55 66 77]
	fmt.Println(v2) // [[11 22 33 44 55] [66 77 88 99]]
	fmt.Println(v3) // [[1 2] [3 4] [3 4]]
	fmt.Println(v4) // [[1 2] [3 4] [3 4]]
	fmt.Println("修改切片嵌套的数据")
	v1[0] = 999
	fmt.Printf("数据类型：%T\n", v1[0]) // 数据类型：int
	v2[0][0] = 999
	fmt.Printf("数据类型：%T\n", v2[0]) // 数据类型：[]int
	v3[1][0] = 999
	fmt.Printf("数据类型：%T\n", v3[1]) // 数据类型：[2]int
	v4[1][1] = 999
	fmt.Printf("数据类型：%T\n", v4[1]) // 数据类型：[2]int
	fmt.Println(v1)                // [999 22 33 44 55 66 77]
	fmt.Println(v2)                // [[999 22 33 44 55] [66 77 88 99]]
	fmt.Println(v3)                // [[1 2] [999 4] [3 4]]
	fmt.Println(v4)                // [[1 2] [3 999] [3 4]]

}

```

### 变量赋值

- 整型

```go
package main

import "fmt"

func main() {
	v1 := false
	v2 := v1

	fmt.Printf("v1的内存地址：%p \n", &v1) // 0xc00012a002
	fmt.Printf("v2的内存地址：%p \n", &v2) // 0xc00012a003
}
```

- 浮点型

```go
package main

import "fmt"

func main() {
	v1 := 3.14
	v2 := v1

	fmt.Printf("v1的内存地址：%p \n", &v1) // 0xc000016050
	fmt.Printf("v2的内存地址：%p \n", &v2) // 0xc000016058
}
```

- 布尔类型

```go
package main

import "fmt"

func main() {
	v1 := false
	v2 := v1

	fmt.Printf("v1的内存地址：%p \n", &v1) // 0xc00012a002
	fmt.Printf("v2的内存地址：%p \n", &v2) // 0xc00012a003
}
```

- 字符串

```go
package main

import "fmt"

func main() {
	v1 := "shdeng"
	v2 := v1

	fmt.Printf("v1的内存地址：%p \n", &v1) // 0xc000010200
	fmt.Printf("v2的内存地址：%p \n", &v2) // 0xc000010210 
}
```

- 数组

```go
package main

import "fmt"

func main() {
	v1 := [2]int{6, 9}
	v2 := v1
	fmt.Println(v1, v2)
	fmt.Printf("v1的内存地址：%p \n", &v1) // 0xc0000b4010
	fmt.Printf("v2的内存地址：%p \n", &v2) // 0xc0000b4020

	v1[0] = 11111
	fmt.Println(v1, v2)
}
``` 

- 切片

```go
package main

import "fmt"

func main() {
	v1 := []int{6, 9}
	v2 := v1                                  // 出现拷贝现象，但是内部的 ptr （unsafe.Pointer）指向同一块内存地址
	fmt.Println(v1, v2)                       // [6 9] [6 9]
	fmt.Printf("v1的内存地址：%p \n", &v1)          // v1的内存地址：0xc000010030
	fmt.Printf("v2的内存地址：%p \n", &v2)          // v2的内存地址：0xc000010048
	fmt.Printf("v1下标为0数据的内存地址：%p \n", &v1[0]) // v1下标为0数据的内存地址：0xc00001c070
	fmt.Printf("v2下标为0数据的内存地址：%p \n", &v2[0]) // v2下标为0数据的内存地址：0xc00001c070
	fmt.Printf("v1下标为1数据的内存地址：%p \n", &v1[1]) // v1下标为1数据的内存地址：0xc00001c078
	fmt.Printf("v2下标为1数据的内存地址：%p \n", &v2[1]) // v2下标为1数据的内存地址：0xc00001c078

	v1[0] = 11111
	fmt.Println(v1, v2) // [11111 9] [11111 9]
}

```

`引用类型和值类型 是 人为定义的。需要根据数据类型的底层原理去理解`

`new 和 make的区别`

```go
// make 适用于 容器类型的数据结构。如：切片，map，channel链表
// new 声明指针类型数据
```