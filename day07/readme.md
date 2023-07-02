# day 07 数据类型

### 常见数据类型

- 结构体，自定义数据集合
- 函数，用于 将重复的动作以函数的形式集成，便于多次调用
- 接口，用于 `约束` 和 `泛指数据类型`

### 结构体

- 什么是结构体：
    - 结构体是复合类型的数据，用于表示一组数据
    - 结构体 由 一系列 属性组成，每个属性都有自己的类型和值

###### 结构体定义

```go
package main

import "fmt"

/*
	type 结构体名称 struct {
		字段 类型
		...
}
*/

type Address struct {
	city, state string
}

// 定义结构体
type Person struct {
	name        string
	age         int
	email       string
	sex, idCard string // 同一类型的字段，可以这个写
	hobby       []string
	Address     // 匿名结构体 ， 等同于 Address Address
}

func main() {
	// 初始化结构体
	var p1 = Person{"deng", 18, "dengshuai37@163.com", "nan", "000",
		[]string{"1", "2"}, Address{"bj", "cp"}}

	// 获取结构体中的数据
	fmt.Println(p1.name)          // deng
	fmt.Println(p1.age)           // 18
	fmt.Println(p1.email)         // dengshuai37@163.com
	fmt.Println(p1.sex)           // nan
	fmt.Println(p1.idCard)        // 000
	fmt.Println(p1.city)          // bj 等同于 p1.Address.city
	fmt.Println(p1.Address.city)  // bj
	fmt.Println(p1.state)         // cp
	fmt.Println(p1.Address.state) // cp 等同于 p1.Address.state

	// 修改 结构体数据
	p1.age = 20
	fmt.Println(p1.age) // 20
}

```

###### 结构体初始化 并 使用

- 根据结构体创建一个对象

```go
package main

import "fmt"

func main() {
	// 定义一个结构体（类型），每个结构体包含 name、age、hobby 三个元素
	type Address struct {
		city, state string
	}
	type Person struct {
		name    string
		age     int
		Address // 匿名字段，那么默认Person就包含了Address的所有字段
	}

	//方式1：先后顺序
	p1 := Person{"shdeng", 19, Address{"北京", "中国"}}
	fmt.Println(p1.name, p1.age, p1.city, p1.state) // shdeng 19 北京 中国

	//方式2：关键字
	p2 := Person{name: "shdeng", age: 19, Address: Address{city: "北京", state: "中国"}}
	fmt.Println(p2.name, p2.age, p2.city, p2.state, p2.Address.city, p2.Address.state) // shdeng 19 北京 中国 北京 中国

	//方式3：先声明再赋值
	var p3 Person
	p3.name = "shdeng"
	p3.age = 50
	p3.Address = Address{
		city:  "北京",
		state: "BJ",
	}
	fmt.Println(p3.name, p3.age, p3.Address.city, p3.Address.state) // shdeng 50 北京 BJ
	// 或
	var p4 Person
	p4.name = "shdeng"
	p4.age = 50
	p4.city = "北京"
	p4.state = "BJ"
	fmt.Println(p3.name, p3.age, p3.Address.city, p3.Address.state) // shdeng 50 北京 BJ

}

```

###### 结构体指针

- 创建 指针类型 结构体

```go
package main

import "fmt"

func main() {

	type Person struct {
		name string
		age  int
	}
	// 初始化指针类型结构体
	// 方式一
	people := Person{"shdeng", 18}
	fmt.Println(people.name, people.age)
	p2 := &people
	fmt.Printf("p2类型：%T , p2地址：%p\n", p2, &p2) // p2类型：*main.Person , p2地址：0xc000012030

	// 方式二
	p3 := &Person{"zhou", 25}
	fmt.Printf("p3类型：%T , p3地址：%p\n", p3, &p3) // p3类型：*main.Person , p3地址：0xc000012038

	// 方式三
	var p4 *Person = new(Person)
	p4.name = "dz"
	p4.age = 5

	fmt.Println(p4.name, p4.age)
	fmt.Printf("p4类型：%T , p4地址：%p\n", p4, &p4) // p3类型：*main.Person ,p4地址：0xc000012040

}

```  

- 内存管理

```go
package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}

	// 初始化结构体
	p1 := Person{"shdeng", 18}
	fmt.Println(p1.name, p1.age)

	// 初始化结构体指针
	p2 := &Person{"shdeng", 1} // p2 指向 结构体的内存地址
	fmt.Println(p2.name, p2.age)

}

```

###### 结构体赋值

- 拷贝赋值

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Person struct {
		name string
		age  int
	}

	p1 := &Person{"shdeng", 18} // 指针类型结构体
	p2 := p1                    // 指针类型结构体

	fmt.Println(p1, reflect.TypeOf(p1), &p1) // &{shdeng 18} *main.Person 0xc000012028

	fmt.Println(p2, reflect.TypeOf(p2), &p2) // &{shdeng 18} *main.Person 0xc000012030

	p1.name = "zhou"
	fmt.Println(p1, reflect.TypeOf(p1), &p1) // &{zhou 18} *main.Person 0xc000012028

	fmt.Println(p2, reflect.TypeOf(p2), &p2) // &{zhou 18} *main.Person 0xc000012030

}

```

- 结构体指针赋值
- 结构体嵌套赋值拷贝

###### 结构体标签 tag

### 函数

- 目的：代码块，提高代码重复利用

```goregexp
func 函数名(参数) 返回值 {
  函数体
}
```

###### 函数定义

```go
package main

import "fmt"

// 定义函数
func SendEmail() {
	fmt.Println("发送邮件了...")
}

func main() {
	// 1000行 or 10000行
	// ...
	// 执行函数
	SendEmail()
	// ...
	// 10行代码可以实现发邮件
}
```

###### 函数参数

传值时会拷贝一份数据（等同于赋值拷贝）

- 多个参数

```go
package main

import "fmt"

func add(num1 int, num2 int) (int, bool) {
	result := num2 + num1 // num1 和 num2 是新拷贝的数据
	return result, true
}

func main() {
	data, flag := add(1, 8)
	fmt.Println(data, flag) // 9 true
}

```  

```go
package main

import "fmt"

func SendEmail(arg [2]int) {
	arg[0] = 666
}

func main() {
	dataList := [2]int{11, 22}
	SendEmail(dataList)
	fmt.Println(dataList) // [11,22]
}
```

- 指针参数

```go
package main

import "fmt"

func changeData(dataList *[3]string) {
	dataList[1] = "shdeng" // dataList 指针类型的参数
}

func main() {
	userList := [3]string{"zhou", "yun", "xia"}
	changeData(&userList) // 将地址传入，使用的是同一份数据
	fmt.Println(userList) // [zhou shdeng xia]
}

```  

- 函数做参数

```go
package main

import "fmt"

func add100(arg int) (int, bool) {
	return arg + 100, true
}

// 函数 作为参数
func proxy(data int, exec func(int) (int, bool)) int {
	data, flag := exec(data)
	if flag {
		return data
	} else {
		return 9999
	}
}

func main() {
	result := proxy(123, add100) // 函数 作为参数
	fmt.Println(result)          // 223

}
```  
```go
package main

import "fmt"

func add100(arg int) (int, bool) {
   return arg + 100, true
}

type f1 func(arg int) (int, bool) // type 定义类型，将 f1 指向就是add100函数，当传入add100时，触发

func proxy(data int, exec f1) int {
   data, flag := exec(data)
   if flag {
      return data
   } else {
      return 9999
   }
}

func main() {
   result := proxy(123, add100)
   fmt.Println(result)

}
```
- 变长（动态）长参数
```go
package main

import "fmt"

// ... 省略表示可长参数 ， num 将会以 切片形式存储参数
func do(num ...int) int {
   sum := 0
   for _, value := range num {
      sum += value
   }
   return sum
}

func main() {
   r1 := do(1, 2, 3, 3)
   r2 := do(0, 1, 1)
   fmt.Println(r1,r2)
}
```
###### 函数返回值

- 多个返回值
```go
package main

import "fmt"

func add100(arg int) (int, bool, int) {
   return arg + 100, true, 888
}

func main() {
   v1, v2, v3 := add100(555)
   fmt.Println(v1, v2, v3)
}
```  

- 返回函数
```go
package main

import (
	"fmt"
)

func exec(num1 int, num2 int) string {
	fmt.Println("执行函数了")
	return "成功"
}

type f1 func(int, int) string // 定义函数类型

func getFunction() f1 {
	return exec
}

func main() {
	function := getFunction() // 获取 exec函数
	result := function(111, 222)
	fmt.Println(result)
}
```  

- 匿名函数和返回函数
```go
package main

import "fmt"

func F1(n1 int, n2 int) func(int) string {

	return func(n1 int) string {
		fmt.Println("匿名函数")
		return "匿名"
	}
}

func main() {
	// 匿名函数
	v1 := func(n1 int, n2 int) int {
		return 123
	}
	data := v1(11, 22)
	fmt.Println(data)

	value := func(n1 int, n2 int) int {
		return 123
	}(11, 22)
	fmt.Println(value)

}
```
###### 闭包

###### defer

- 特点：函数执行完成之后触发，一般用于结束操作后，释放资源
```package main

import "fmt"

func do() int {
	fmt.Println("风吹")
	defer fmt.Println("函数执行完毕了")  // 如果在这行之前执行return，那么defer就不再执行
	fmt.Println("屁屁凉")
	return 666
}

func main() {
	ret := do()
	fmt.Println(ret)
}
// 风吹
// 屁屁凉
// 函数执行完毕了
```

```go
// 多个defer时，最终函数执行完毕后会按照倒序的方式去执行
package main

import "fmt"

func other(a1 int, a2 int) {
   fmt.Println("defer函数被执行了")
}

func do() int {
   fmt.Println("风吹")
   defer fmt.Println("函数执行完毕了")
   defer other(1, 22)
   fmt.Println("屁屁凉")
   return 666
}


func main() {
   ret := do()
   fmt.Println(ret)
   // 风吹
   // 屁屁凉
   // defer函数被执行了
   // 函数执行完毕了
   // 666
}


```
###### 自执行函数
```go
package main

import "fmt"

func main() {
	
	// 可以理解为匿名函数
	result := func(arg int) int {
		return arg + 100
	}(123)

	fmt.Println(result) // 223
}

```
###### 结构体作为 函数的 参数 和 返回值

###### 类型方法

###### 方法继承

###### 结构体工厂

### 接口 interface

- 特点：接口是特殊的数据类型 interface

###### 接口的作用

- 空接口，指代任意类型

- 非空接口，规范+约束

###### 接口底层实现

- 空接口
- 非空接口
- 