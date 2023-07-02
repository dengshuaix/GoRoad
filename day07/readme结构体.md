# day 07 数据类型

### 常见数据类型

- 结构体，自定义数据集合

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

  p1 := Person{name: "shdeng", age: 18}
  p2 := p1 // 内部将p1重新拷贝一份

  fmt.Println(p1, reflect.TypeOf(p1), &p1) // {shdeng 18} main.Person &{shdeng 18}

  fmt.Println(p2, reflect.TypeOf(p2), &p2) //{shdeng 18} main.Person &{shdeng 18}

  p1.name = "zhou"

  fmt.Println(p1, reflect.TypeOf(p1), &p1) // {zhou 18} main.Person &{zhou 18}

  fmt.Println(p2, reflect.TypeOf(p2), &p2) // {shdeng 18} main.Person &{shdeng 18}

}

```

- 结构体指针赋值
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

- 结构体嵌套赋值拷贝
```go
package main

import "fmt"

func main() {
	type Address struct {
		city, state string
	}

	type Person struct {
		name    string
		age     int
		address Address
	}

	p1 := Person{name: "二狗子", age: 19, address: Address{"北京", "BJ"}}
	p2 := p1

	fmt.Println(p1.address, &p1.address) // {{北京 BJ} &{北京 BJ}

	fmt.Println(p2.address, &p2.address) // {北京 BJ} &{北京 BJ}

	p1.address.city = "上海"

	fmt.Println(p1.address, &p1.address) // {上海 BJ} &{上海 BJ}

	fmt.Println(p2.address, &p2.address) // {北京 BJ} &{北京 BJ}

}

```
- 谁不拷贝？
  其实本质上都拷贝了，只不过由于数据存储方式的不同，导致拷贝的有些是数据，有些是内存地址（指针）。

  - 感觉拷贝：字符串、数组、整型等。
  - 感觉不拷贝：map、切片。
  
```go
type Address struct {
    city, state string
}
type Person struct {
    name    string
    age     int
    hobby   *[2]string
    num     []int
    parent  map[string]string
    address Address
}

p1 := Person{
    name:   "二狗子",
    age:    19,
    hobby:  &[2]string{"裸奔", "大保健"},
    num:    []int{69, 19, 99, 38},
    parent: map[string]string{"father": "Alex", "mother": "Monika"},
}
p2 := p1
p1.hobby[0] = "洗澡"

fmt.Println(p1.hobby) // &[洗澡 大保健]
fmt.Println(p2.hobby) // &[洗澡 大保健]

```
###### 结构体标签 tag
```go
/*
 @Author:武沛齐  微信号：wupeiqi666
 @Description: 老男孩IT教育 & 路飞学城
 @Video:  https://space.bilibili.com/283478842
*/
package main

import (
  "fmt"
  "reflect"
)

func main() {
  type Person struct {
    name string "姓名" // 姓名就是tag
    age  int32  "年龄"
    blog string "博客"
  }

  p1 := Person{name: "武沛齐", age: 18, blog: "https://www.pythonav.com"}

  p1Type := reflect.TypeOf(p1)

  // 方式1 获取 tag
  filed1 := p1Type.Field(0) 
  fmt.Println(filed1.Tag) // 姓名   filed1.Name -> name

  // 方式2 获取 tag
  filed2, _ := p1Type.FieldByName("blog")
  fmt.Println(filed2.Tag) // 

  // 循环获取 获取 tag
  fieldNum := p1Type.NumField() // 总共有多少个字段 3
  // 循环：0 1 2 
  for index := 0; index < fieldNum; index++ {
    field := p1Type.Field(index)
    fmt.Println(field.Name, field.Tag) //    name  姓名;    age 年龄  ;   blog 博客
  }
}
```
