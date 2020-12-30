package main

import "fmt"

type Groundhog string

func main() {
	var g Groundhog
	g = "土拨鼠"
	g.Run()
	g.Say()
	g.Cry()
}

//在函数声明时，在其名字之前放上一个变量，就是一个方法。
//下面的代码中有一个附加的参数g，叫做方法的接收器，在go语言中不会像其他语言那样使用
//this或self作为接收器。要注意这个Groundhog必须是一个类型。
func (g Groundhog) Run() {
	fmt.Println(g + "跑了并说你写的代码真的有Goland！🤣")
}

func (g Groundhog) Say() {
	fmt.Println(g + "看了看你说：你瞅啥？🤨")
}

func (g Groundhog) Cry() {
	fmt.Println(g + "哭了，因为它的头发掉光了，它成兔子了\U0001F97A")
}
