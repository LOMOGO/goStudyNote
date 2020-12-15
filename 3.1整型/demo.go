package main

import "fmt"

func main() {
	//Go语言中有两种内置的类型别名：byte是uint8的内置别名。我们可以把byte和uint8看作是同一个类型
	//rune是int32的内置别名。我们可以把rune和int32看作同一个类型

	//从Go1.13开始，下划线_可以出现在整数、浮点数和虚部数字面量中，但是需要注意的是一个下划线_是不能出现在此字面量的首尾的
	num := 1000_0001
	_ = 0x1_23
	_ = 0b100_1
	_ = 1.05_24
	//_ = 1_.05 这种是不合法的
	//_ = 6__9 这种也不合法
	fmt.Println(num)
	//整数类型值有四种字面量形式：十进制形式（decimal）、八进制形式(octal)、十六进制形式(hex)、二进制形式(binary)
	//下面的三个字面量均表示十进制的15
	hex := 0xF //十六进制必须使用0X或者0x开头
	octal := 017//八进制必须使用0、0o、或者0O开头
	binary := 0b1111//二进制必须使用0b或者0B开头

	//%b表示以二进制格式打印  %o表示以八进制格式打印  %x表示以十六进制打印
	fmt.Printf("binary:%b\toctal:%o\thex:%x\n", binary, octal, hex)
	fmt.Printf("B2D:%d\tO2D:%d\tH2D:%d\n", binary, octal, hex)
}
