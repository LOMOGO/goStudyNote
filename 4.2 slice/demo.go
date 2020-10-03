package main

import "fmt"

func main() {
	//内置的len()、cap()函数分别用来返回slice的长度和容量
	s := []int{1, 2, 3}
	fmt.Println(len(s), cap(s))//3 3
	fmt.Printf("%T\n", s)//[]int
	//s[4] = 34 引发panic异常
	fmt.Println(s)//[1 2 3]
	//slice唯一合法的比较操作是和nil比较，如果要测试一个slice对象是否是空的，使用len(s) == 0来判断
	fmt.Println(s == nil)//flase
	//
	a := make([]int, 7)
	fmt.Println(len(a), cap(a))//7 7
	b := make([]int, 7, 11)
	fmt.Println(len(b), cap(b))//7 11
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
}
