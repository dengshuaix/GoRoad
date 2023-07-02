package main

import "fmt"

func main() {
	//8. map 变量赋值
	v18 := map[string]string{"n1": "shdeng", "n2": "zhou"}
	v19 := v18

	v18["n1"] = "yx"

	fmt.Println(v18) // map[n1:yx n2:zhou]
	fmt.Println(v19) // map[n1:yx n2:zhou]
}
