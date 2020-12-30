package main

import "fmt"

type num struct {
	num int
}

func main() {
	var n num
	n.num = 1
	fmt.Println(n.num)
	n.Add(1)
	fmt.Println(n.num)
	n.Mul(3)
	fmt.Println(n.num)
	n.Mul(3)
	fmt.Println(n.num)
	n.Add(2)
	fmt.Println(n.num)
}

//type num *int
//Invalid receiver type 'num' ('num' is a pointer type)
//如果一个接收器类型本身就是指针的话，是不被允许的
/*func (n num) Add(par int)  {

}*/

func (n num) Add(par int) {
	n.num += par
}

func (n *num) Mul(par int) {
	n.num *= par
}
