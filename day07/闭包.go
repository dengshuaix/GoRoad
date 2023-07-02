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
			fmt.Println(i) // i在每个函数内部传递
		}
		functionList = append(functionList, function)
	}
	// 执行到此步骤时， i 已经等于5
	functionList[0]() // 5
	fmt.Println(&functionList[0]) // 0xc00007c040
	functionList[1]() // 5
	fmt.Println(&functionList[1]) // 0xc00007c048

	functionList[2]() // 5
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
