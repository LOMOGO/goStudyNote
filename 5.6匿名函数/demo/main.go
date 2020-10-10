package main

import "fmt"

//函数字面量是一种表达式(func 关键字后没有函数名),它的值被称为匿名函数。在包函数中定义的匿名函数可以引用该函数的变量

func squares() func() int {
	var x int

	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()

	fmt.Println(squares()()) //1
	fmt.Println(squares()()) //1
	fmt.Println(squares()()) //1
	fmt.Println(squares()()) //1
	fmt.Println(squares()()) //1

	fmt.Println("---------------")
	//这个例子证明，函数值不仅仅是一串代码，还记录了状态，在squares定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引用。这是函数值属于引用类型和不可比较的原因。Go使用闭包技术实现函数值，因此也可以把函数值叫做闭包
	//变量的声明周期不由它的作用域决定: squares返回后，变量x仍然隐式的存在于f中。
	fmt.Println(f()) //1
	fmt.Println(f()) //4
	fmt.Println(f()) //9
	fmt.Println(f()) //16
	fmt.Println(f()) //25
}
