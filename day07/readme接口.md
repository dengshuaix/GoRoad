# day 07 数据类型

### 常见数据类型
- 接口，用于 `约束` 和 `泛指数据类型`

### 接口 interface

- 特点：接口是特殊的数据类型 interface
```goregexp
type 接口名称 interface{
    方法名称() 返回值
}
```
```goregexp
// 接口中的方法只定义，不能编写具体的实现逻辑
type Base interface {
	f1()                   // 定义方法，无返回值
	f2() int               // 定义方法，返回值int类型
	f3() (int, bool)       // 定义方法，2个返回值分别是 int、bool类型
	f4(n1 int, n2 int) int // 定义方法，需要两个参数，1个返回值
}

type empty interface {}  // interface{} 
```
###### 接口的作用

- 空接口，指代任意类型
    - 接口只是代指这些数据类型
    - 在内部其实是转换为了接口类型
    - 想要再获取数据中的值时，需要再将接口转换为指定的数据类型
```go
// 示例一
package main

import (
	"fmt"
	"reflect"
)

// 定义空接口
type Base interface {
	
}

func main() {
    //定义一个切片，内部可以存放任意类型。
	dataList := make([]Base, 0)  // 推荐简写为：dataList := make([]interface{}, 0)
    
    // 切片中添加 字符串类型
	dataList = append(dataList, "shdeng")
	// 切片中添加 整型
	dataList = append(dataList, 18)
    // 切片中添加 浮点型
	dataList = append(dataList, 99.99)
}
```
```go
// 示例二
/*
 @Author:shdeng  微信号：wupeiqi666
 @Video:  https://space.bilibili.com/283478842
*/
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func something(arg interface{}) {
	fmt.Println(arg)
}

func main() {
	something("shdeng")
	something(666)
	something(4.15)
	something(Person{name: "wupeiqi", age: 18})
}
```
```go
// 示例三
package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Role struct {
	title string
	count int
}

func something(arg interface{}) {

	// 多个类型转换，将arg接口对象转换为。（断言）
	switch tp := arg.(type) {
	case Person:
		fmt.Println(tp.name)
	case Role:
		fmt.Println(tp.title)
	case string:
		fmt.Println(tp)
	default:
		fmt.Println(tp)
	}
}

func main() {
	something("shdeng")
	something(666)
	something(4.15)
	something(Person{name: "wupeiqi", age: 18})
	something(Role{title: "管理员", count: 2})
}
```
- 非空接口，规范+约束
    - 用于约束结构体中必须含有某个方法
```go
// 示例一
package main

import "fmt"

// 定义接口
type IBase interface {
	f1() int
}

// 定义结构体Person
type Person struct {
	name string
}

// 为结构体Person定义方法
func (p Person) f1() int {
	return 123
}

// 定义结构体User
type User struct {
	name string
}

// 为结构体User定义方法
func (p User) f1() int {
	return 666
}

// 基于接口的参数，可以实现传入多中类型（多态），也同时具有约束对象必须实现接口方法的功能
func DoSomething(base IBase) {
	result := base.f1() //直接调用  接口.f1()  -> 找到其对应的类型并执行其方法
	fmt.Println(result)
}

func main() {

	per := Person{name: "shdeng"}
	user := User{name: "wupeiqi"}

	DoSomething(per)
	DoSomething(user)
}
```
```go
// 示例二
package main

import "fmt"

// 定义接口
type IBase interface {
	f1() int
}

// 定义结构体
type Person struct {
	name string
}

// 为结构体定义方法
func (p *Person) f1() int {
	return 123
}

// 定义结构体
type User struct {
	name string
}

// 为结构体定义方法
func (p *User) f1() int {
	return 666
}

// 基于接口的参数，可以实现传入多中类型（多态），也同时具有约束对象必须实现接口方法的功能
func DoSomething(base IBase) {
	result := base.f1() // 由于base的约束，此处智能执行IBase中约束的接口。
	fmt.Println(result)
}

func main() {

	per := &Person{name: "shdeng"} // 创建结构体对象并获取其指针对象
	user := &User{name: "wupeiqi"}

	DoSomething(per)
	DoSomething(user)
}
```
###### 接口底层实现
- 接口是Go的一种数据类型，接口可以代指任意类型，实现参数、”容器“中可以处理多种数据类型
```goregexp
type eface struct {
   _type *_type			// 存储类型相关信息
   data  unsafe.Pointer // 存储数据
}
```  

- 空接口
    - 现其他对象赋值给空接口, 就是将其他对象相关的值存放到eface的 _type和data
    - _type是一个结构体内部存储挺多的信息，这里统称为类型相关的信息
```goregexp
// The conv and assert functions below do very similar things.
// The convXXX functions are guaranteed by the compiler to succeed.
// The assertXXX functions may fail (either panicking or returning false,
// depending on whether they are 1-result or 2-result).
// The convXXX functions succeed on a nil input, whereas the assertXXX
// functions fail on a nil input.

func convT2E(t *_type, elem unsafe.Pointer) (e eface) {
   if raceenabled {
      raceReadObjectPC(t, elem, getcallerpc(), funcPC(convT2E))
   }
   if msanenabled {
      msanread(elem, t.size)
   }
   x := mallocgc(t.size, t, true)
   // TODO: We allocate a zeroed object only to overwrite it with actual data.
   // Figure out how to avoid zeroing. Also below in convT2Eslice, convT2I, convT2Islice.
   typedmemmove(t, x, elem)
   e._type = t
   e.data = x
   return
}
```    
```go
// 示例一
package main

func main() {
	num := 666
	var object interface{}
	// 将num的类型int存储到 _type 中；值8存储到data中
	object = num
}
```
```go
// 示例二
package main

func DoSomething(arg interface{}) {
	// 将 num的类型int 存储到 _type 中；值8存储到data中;
}

func main() {
	num := 666
	DoSomething(num)
}
```
```go
// 示例三
package main

type Person struct {
	name string
	age  int
}

func DoSomething(arg interface{}) {
	// 将 info的类型Person 存储到 _type 中；值{name: "shdeng", age: 18}存储到data中;
}

func main() {
	info := Person{name: "shdeng", age: 18}
	DoSomething(info)
}
```
- 非空接口
    - 非空接口会定义一些方法来实现约束
```goregexp
type iface struct {
   tab  *itab           // 类型和方法相关
   data unsafe.Pointer  // 数据
}

type itab struct {
	inter *interfacetype // 接口信息，如：接口中定义的方法。
	_type *_type         // 类型
	hash  uint32
	_     [4]byte
	fun   [1]uintptr
}

type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod   // 接口的方法
}
```    
