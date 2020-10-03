package main

import "fmt"

func main() {
	s := "hello,world"
	fmt.Println(len(s))
	s = "我hello,world"
	fmt.Println(len(s))
	//索引操作s[i]返回第i个字节,汉字的占的字节数是3，中文标点符号也是3
	fmt.Println(s[0], s[1], s[2], s[3])
	//子字符串操作s[i:j]基于原始的s字符串的第i个字节开始到第j个字节（不包括j本身），生成一个新的字符串
	fmt.Println(s[0:3])
	//字符串是不可修改的，因此尝试修改字符串内部的数据的操作也是被禁止的。
	s = `你可以开开心心的去做很多的事情
"不需要什么理由"
` + "`" +`哈哈哈`
	fmt.Println(s)
}
