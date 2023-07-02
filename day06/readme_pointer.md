# day 06 数据类型

### 常见数据类型

- 指针，用于表示内存地址的类型

### 指针

- 指针：数据类型，用于表示数据的内存地址
- 特点：节约内存，实现数据驻留/参数驻留 。 可理解为 单例行为

###### 指针声明

```go
package main

import "fmt"

func main() {
	// 声明一个 字符串类型 的变量（默认初始化值为空字符串）。
	var v1 string
	// 声明一个 字符串的指针类型 的变量（默认初始化值为nil）。
	var v2 *string
	var v3 int
	var v4 *int
	fmt.Println(v1) // 空值
	fmt.Println(v2) // <nil>
	fmt.Println(v3) // 0
	fmt.Println(v4) // <nil>

	// 声明一个 字符串类型 的变量，值为 shdeng。
	var name string = "shdeng"
	// 声明一个 字符串的指针类型 的变量，值为 name 对应的内存地址。
	var pointer *string = &name
	fmt.Println(name)    // shdeng
	fmt.Println(pointer) // 0xc000014270
	var age int = 18
	var x1 *int = &age
	fmt.Println(age, x1) // 18 0xc0000ac020

}
```

###### 指针意义

```go
package main

import "fmt"

func main() {
	// 节约内存空间
	v1 := "shdeng"
	v2 := &v1
	fmt.Println(v1, v2, *v2) // shdeng 0xc000014270 shdeng
	v1 = "zhou"              // v2 指向 v1的内存地址。 修改v1 相当于修改v2
	fmt.Println(v1, v2, *v2) // zhou 0xc000014270 zhou
}

```

###### 指针场景

```go
package main

import "fmt"

func main() {

	// 场景一 ， 指针类型的数据，内存地址指向同一个
	v1 := "shdeng"
	v2 := v1
	v1 = "周"

	fmt.Println(v1, v2) // 周 shdeng

	v3 := "shdeng"
	v4 := &v3
	v3 = "周"

	fmt.Println(v3, *v4) // 周 周

}
```

```go
package main

import "fmt"

func changeData(data string) {
	fmt.Printf("%p\n", &data) // 0xc000014280
	data = "哈哈哈哈哈"
	fmt.Printf("%p\n", &data) // 0xc000014280

}

func main() {

	// 场景二 函数的参数非指针
	name := "shdeng"
	// 本质上会将name的值拷贝一份，并赋值给data
	changeData(name)
	fmt.Println(name)         // shdeng
	fmt.Printf("%p\n", &name) // 0xc000014270

}
```

```go
package main

import "fmt"

func changeData(ptr *string) {
	fmt.Printf("%p\n", &ptr) // 0xc000014280
	*ptr = "哈哈哈哈哈"
	fmt.Printf("%p\n", &ptr) // 0xc000014280

}

func main() {

	// 场景三 函数的参数指针，函数的参数作为指针，传入的数据就不进行拷贝，指针参数指向同一份内存地址
	name := "shdeng"
	changeData(&name)
	fmt.Println(name)         // 哈哈哈哈哈
	fmt.Printf("%p\n", &name) // 0xc000014270

}

```

```go
package main

import "fmt"

func main() {
	// 场景四，Scanf 输入语句，&username是指向上述定义的username参数的地址，即同一份内存地址
	var username string
	fmt.Printf("请输入用户名：")

	// 提示用户输入，当用户输入之后，将输入的值赋值给内存地址对应的区域中。
	fmt.Scanf("%s", &username)

	if username == "shdeng" {
		fmt.Println("登录成功")
	} else {
		fmt.Println("登录失败")
	}
}
```

###### 指针嵌套（指针的指针）
```go
package main

import "fmt"

func main() {
	// 指针嵌套
	name := "shdeng"

	// 声明一个指针类型变量p1，内部存储name的内存地址
	var p1 *string = &name

	// 声明一个指针的指针类型变量p2，内部存储指针p1的内存地址
	var p2 **string = &p1

	// 声明一个指针的指针的指针类型变量p3，内部存储指针p2的内存地址
	var p3 ***string = &p2
	
	//改变name值
	***p3="zhou"
	fmt.Println(name)
}

```
##### 指针高级操作
- 数组的地址 == 数组的第一个元素的地址。
```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//数组的地址 == 数组的第一个元素的地址。
	dataList := [3]int8{1, 2, 3}
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &dataList, dataList)       // 内存地址：0xc00001c062 - 数据类型：[3]int8 是数组类型
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &dataList[0], dataList[0]) // 内存地址：0xc00001c062 - 数据类型：int8 元素类型
	fmt.Println()
	// 指针的计算
	// 1. 获取数组第一个元素的地址
	var firstDataPtr *int8 = &dataList[0]
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &firstDataPtr, firstDataPtr)

	// 2. 转换成 Pointer 类型
	ptr := unsafe.Pointer(firstDataPtr)
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &ptr, ptr)

	// 3. 转换成 uinptr ， 然后进行地址计算（地址加上1个字节，等同于取出第二个索引位置的数据）
	targetAddress := uintptr(ptr) + 1
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &targetAddress, targetAddress)

	// 4. 根据新地址，重新转换成Pointer类型
	newPtr := unsafe.Pointer(targetAddress)
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &newPtr, newPtr)

	// 5. Pointer对象转换成 int8 指针类型
	newValue := (*int8)(newPtr) // 类型强制转换
	fmt.Printf("内存地址：%p - 数据类型：%T\n", &newValue, newValue)

	// 6. 根据 指针获取值
	fmt.Println(*newValue)
	/*
		// 指针计算示意图
		内存地址：0xc000012030 - 数据类型：*int8
		内存地址：0xc000012038 - 数据类型：unsafe.Pointer
		内存地址：0xc00001c068 - 数据类型：uintptr
		内存地址：0xc000012040 - 数据类型：unsafe.Pointer
		内存地址：0xc000012048 - 数据类型：*int8
		2
	*/
}

```
// 不允许在 uintptr 和 unsafe.Pointer 转换之间留出运行 GC 的空间
