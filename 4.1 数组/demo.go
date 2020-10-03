package main

import "fmt"

func main() {
	const n = 5
	//数组的长度必须是常量表达式，因为数组的长度是固定的，数组的长度需要在编译阶段确定。
	var ns [n]int
	fmt.Printf("%T\n", ns)

	symbol := [...]string{4: "sf", 1: "pj", 3: "rgr", 0: "jb"}
	for i, v := range symbol {
		fmt.Println(i, v)
	}

	a := [...]int{1, 3}
	b := [...]int{3, 1}
	c := [...]int{1, 3}
	fmt.Println(a == b, a == c)

	//b1 := []byte("ageargergr")
}
