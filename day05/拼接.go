package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	// strings.Join 拼接字符串
	dataList := []string{"i", "love"}
	result1 := strings.Join(dataList, "-")
	fmt.Println(result1) //i-love

	// 建议：效率更高一些（go 1.10之前）
	var buffer bytes.Buffer
	buffer.WriteString("i")
	buffer.WriteString("l")
	buffer.WriteString("o")
	buffer.WriteString("v")
	buffer.WriteString("e")
	result2 := buffer.String()
	fmt.Println(result2)

	// 建议：效率更更更更高一些（go 1.10之后）
	var builder strings.Builder
	builder.WriteString("l")
	builder.WriteString("o")
	builder.WriteString("v")
	builder.WriteString("e")
	result3 := builder.String()
	fmt.Println(result3)
}
