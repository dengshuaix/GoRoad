# day04 go 语言中的编码

### go 编译器 使用 utf-8 编码

- unicode（万国码）
    - ucs2，用16位来表示所有的情况。2**16=65535
      ```go
        // 两个字节
        00000000 00000000 
        ....
	    11111111 11111111
      ```
    - ucs4，用32位来表示所有的情况。2**32=4294967296
      ```go
        // 四个字节
        0000000 0000000 0000000 0000000
        ........
        11111111 11111111 11111111 11111111
      ```
    - ucs2和ucs4应该根据自己的业务场景来进行选择。

- utf-8编码 ， 3个字节
    - 对unicode进行压缩，到底是如何压缩?
        - 找模板
        - 码位以二进制展示，再根据模板进行转换
        - wusir ： https://pythonav.com/wiki/detail/1/80/
        - 码典 ：https://www.unicode.org/versions/Unicode15.0.0/

### 示例

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 定义字符串，字符串是以什么形式存在于Go编译器（utf-8编码）
	name := "邓周"

	// 其他语言：
	// 0 邓
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2)) // 233 11101001
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2)) // 130 10000010
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2)) // 147 10010011

	// 1 周
	fmt.Println(name[3])
	fmt.Println(name[4])
	fmt.Println(name[5])

}

```