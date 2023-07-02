package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	// os.Stdin 是标准输入
	// bufio 是 读取文本模块
	reader := bufio.NewReader(os.Stdin)
	// line，从stdin中读取一行的数据（字节集合 -> 转化成为字符串）
	// reader默认一次能4096个字节（4096/3)
	//    1. 一次性读完，isPrefix=false
	// 	  2. 先读一部分，isPrefix=true，再去读取isPrefix=false

	line, _, _ := reader.ReadLine()      //  接收输入数据
	fmt.Println(reflect.TypeOf(line))    // []uint8
	fmt.Println(reflect.TypeOf(line[0])) // uint8
	data := string(line)
	fmt.Println(data)
}
