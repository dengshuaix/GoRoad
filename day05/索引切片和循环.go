package main

import "fmt"

func main() {
	var name string = "邓周周邓"
	//1。 索引获取字节
	fmt.Println("字节：", name[0])
	fmt.Println("字节：", name[1])
	fmt.Println("字节：", name[2])

	//2。 切片获取区间
	newSliceRes1 := name[0:3]
	newSliceRes2 := name[3:9]
	fmt.Println("切片：", newSliceRes1, newSliceRes2)

	//3。 循环遍历字节
	for i := 0; i < len(name); i++ {
		fmt.Println("字节索引：", i, "字节值：", name[i])
	}

	//4。 range 遍历字符串
	for index, item := range name {
		fmt.Println("字符索引：", index, "字符值：", item, string(item))

	}
	//5。 rune 集合 遍历
	runeList := []rune(name)
	for _, item := range runeList {
		fmt.Println(string(item))
	}

}
