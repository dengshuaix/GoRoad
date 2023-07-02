### 常见数据类型

- 函数，用于 将重复的动作以函数的形式集成，便于多次调用

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
	fmt.Println(r1, r2)
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

```go
package main

import "fmt"

func main() {
	// 闭包的作用
	/*
		1。 保存变量状态：闭包可以用来保存函数执行时所需的变量状态，使得这些状态在函数调用之间保持不变。这样可以实现一些特殊的功能，比如记忆化或者延迟计算。
		2。 实现私有变量：闭包可以创建一个私有作用域，使得外部无法访问其中的变量，从而保证变量的安全性和不被意外修改。
		3。 实现高阶函数：闭包可以作为高阶函数的参数或者返回值，从而实现一些更加灵活的编程模式，比如函数式编程。
		4。 实现回调函数：闭包可以作为回调函数，用于在异步编程中处理异步任务的结果，从而实现更加高效的程序设计。
	*/
	var functionList []func()

	for i := 0; i < 5; i++ {
		function := func() {
			fmt.Println(&i) // 0xc0000ac008
			fmt.Println(i)  // i在每个函数内部传递
		}
		functionList = append(functionList, function)
	}
	// 执行到此步骤时， i 已经等于5
	functionList[0]()             // 5
	fmt.Println(&functionList[0]) // 0xc00007c040
	functionList[1]()             // 5
	fmt.Println(&functionList[1]) // 0xc00007c048

	functionList[2]()             // 5
	fmt.Println(&functionList[2]) //0xc00007c050

	// 存储着5个函数
	var functionList2 []func()

	for i := 0; i < 5; i++ {

		// 以参数形式传入给函数，（重新copy一份，保存本次i的值）
		function := func(arg int) func() {
			return func() {
				fmt.Println(arg)
			}
		}(i) // 匿名函数
		functionList2 = append(functionList2, function)
	}

	// 运行函数
	functionList2[0]() // 0
	functionList2[1]() // 1
	functionList2[2]() // 2
}

```

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

- 结构体做参数和返回值时，在执行时候都会被重新拷贝一份，如果不想被拷贝，则可以通过指针的形式进行处理。
    - 非指针类型结构体

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
}

var P Person = Person{"shdeng", 18}

func doSomething() Person {
	return P
}

func main() {
	data := doSomething() // 返回一个结构体类型数据
	P.name = "zhou"       // 修改 P 不影响 data，因为 data相当于重新 拷贝一份
	fmt.Println(data)     // {shdeng 18}
	fmt.Println(P)        // {zhou 18}
}
```

    - 指针类型结构体

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
}

var P Person = Person{"shdeng", 18}

func doSomething() *Person {
	fmt.Printf("%p\n", &P) // 0x113d230
	return &P              // 返回结构体地址
}

func main() {
	data := doSomething()
	P.name = "zhou"
	fmt.Println(data)         // &{zhou 18}
	fmt.Printf("%p\n", &data) // 0xc000012028

	fmt.Println(P)         // {zhou 18}
	fmt.Printf("%p\n", &P) // 0x113d230

}

```

###### 类型方法

- 为type声明的类型编写一些方法，从而实现对象.方法的操作

```go
package main

import "fmt"

// 声明类型
type MyInt int

// 为MyInt类型自定义一个指针方法
// 可以是指针/可以是类型：*MyInt   MyInt
// 不使用对象，可以用 _ 代替
func (_ *MyInt) DoSomething(a1 int, a2 int) int {
	return a1 + a2
}

func do(a1 int, a2 int) int {
	return a1 + a2
}

func main() {
	var v1 MyInt = 1
	result := v1.DoSomething(1, 2)
	fmt.Println(result)
}
```

- 结构体也是基于type声明的类型

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
	blog string
}

// 为Person结构体类型自定义一个指针方法
// 注意：此处如果不是指针类型的话，再执行方法时结构体对象就会被重复拷贝一份。
func (p *Person) DoSomething(a1 int, a2 int) int {
	return a1 + a2 + p.age
}

func main() {
	p1 := Person{name: "武沛齐", age: 18, blog: "https://www.pythonav.com"}
	// 调用Person的DoSomething方法
	result := p1.DoSomething(1, 2)
	fmt.Println(result)
}
```

###### 方法继承
- 结构体之前存在匿名嵌套关系，则 子结构体可以继承父结构体中的方法
```go
package main

import "fmt"

type Base struct {
   name string
}
type Son struct {
   Base    // 匿名的方式，如果改成 base Base 则无法继承Base的方法。
   age int
}

// Base结构体的方法
func (b *Base) m1() int {
   return 666
}

// Son结构体的方法
func (s *Son) m2() int {
   return 999
}

func main() {
    son := Son{age: 18, Base: Base{name: "武沛齐"}}
	result1 := son.m1()
	result2 := son.m2()
	fmt.Println(result1, result2)
	
	//如果Son结构体中还有与其他结构体嵌套，那么他可以继承所有嵌套结构体中的方法
}
```
###### 结构体工厂
- 工厂的名字以 new 或 New 开头
```goregexp
type File struct {
    fd      int     
    name    string  
}

f := File{10, "xxxxxx"}
```
-  强制使用工厂方法
```goregexp
type File struct {
    fd      int     
    name    string  
}

func NewFile(fd int, name string) *File {
	// 20...
    return &File{fd, name}
}

func main() {
	f1 := NewFile(10, "./test.txt")
    f2 := File{10, "xxxxxx"}
}
```   

-  公有 和 私有
```goregexp
// 私有
type matrix struct {
    ...
}

// 共有
func NewMatrix(params) *matrix {
    m := new(matrix) // 初始化 m
    return m
}
```
```goregexp
package main
import "matrix"
...
// wrong := new(matrix.matrix)     // 编译失败（matrix 是私有的）
right := matrix.NewMatrix(...)  // 实例化 matrix 的唯一方式
```