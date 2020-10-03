package main

import (
	"fmt"
	"go_program_tour/4.4结构体/model"
)

func main() {
	var num int = 5
	var p *int = &num
	fmt.Println(*p)

	type Employee struct {
		ID        int
		Name      string
		Address   string
		Salary    int
		ManagerID int
		position  string
	}
	var dilbert Employee
	var pEmployee *Employee = &dilbert
	//下面三行代码等价
	fmt.Println(dilbert.Address)
	fmt.Println((*pEmployee).Address)
	fmt.Println(pEmployee.Address)

	//通常一行对应一个结构体成员，成员的名字在前，类型在后。不过如果 相邻的成员 的类型如果相同的话，可以被合并到一行，如下面的Name和Address成员那样：
	//结构体成员的输入顺序也有重要的意义，成员的输入顺序不同，结构体的类型也就不同
	//如果结构体成员名字是以大写字母开头的，那么该成员就是导出的，一个结构体可能同时包含导出和未导出(小写)的成员。
	type Employee1 struct {
		ID            int
		Name, Address string
		Salary        int
		ManagerID     int
		position      string
		//一个命名为S的结构体类型不能再包含S类型的成员：因为一个聚合的值不能包含它自身。(该限制同样适用于数组)但是S类型的结构体可以包含 *S 指针类型的成员
		dilbert1 *Employee1
	}

	/*结构体值也可以用结构体字面值表示，结构体字面可以指定每个成员的值。
	  但是下面这种写法要求以结构体成员定义的 顺序 为 每个 结构体成员指定一个字面值。
	  这种语法一般只在定义结构体的包内部使用。*/
	e := Employee{
		1,
		"alta",
		"bet",
		500,
		1,
		"stc",
	}
	fmt.Println("here", e)

	/*而且上面这种结构体字面值的方法不能在外部包中用顺序赋值的方法偷偷初始化未导出的成员，会报错*/
	/*employee3 := model.Employee3{
		1,
		"alta",
		"bet",
		500,
		1,
		"stc",//implicit assignment of unexported field 'position' in model.Employee3 literal
	}*/

	/*还可以用以成员名字和相应的值来初始化的方式。这种方式可以填写部分或全部的成员*/
	_ = model.Employee3{
		ID:      2,
		Name:    "sdit",
		Address: "daf",
	}
	
	/*结构体可以作为函数的参数和返回值。如果考虑到效率的话，较大的结构体通常会用指针的方式传入和放回
	* 如果要在函数内部修改结构体成员的话，用指针传入是必须的；在Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。
	* 因为结构体通常通过指针处理，可以用下面的写法来创建并初始化一个结构体变量，并放回结构体的地址：*/
	_ = &model.Employee3{
		ID:   3,
		Name: "wll",
	}
	//下面的写法等价于上面的写法。
	e3 := new(model.Employee3)
	*e3 = model.Employee3{
		ID:     6,
		Salary: 432,
	}

	//如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，那样的话两个结构体将可以使用==或者!=运算符进行比较。相等比较运算符==将比较两个结构体的每个成员
	//可比较的结构体类型和其他可比较的类型一样，可以用于map的key类型

	/*结构体嵌入和匿名成员：
	  下面我们将看到如何使用go语言提供的结构体嵌入机制让一个命名的结构体包含另一个结构体类型的匿名成员，从而通过简单的点运算符x.f来访问匿名成员链中嵌套的x.d.e.f成员*/

	//例子：一个车轮可以由圆和辐条组成，圆由圆心坐标和圆的半径组成，因此要定义一个车轮结构体可以这样定义出 坐标结构体，圆结构体， 以方便复用
	type Point struct {
		X, Y int
	}

	type Circle struct {
		Center Point
		Radius int
	}

	type Whell struct {
		Circle Circle
		Spokes int
	}

	//改动后的结构体类型变得更加清晰，但是这种修改的同时也导致了访问每个成员变得繁琐，例如：
	var w Whell
	w.Circle.Center.X = 0

	//golang有的匿名成员特性可以让我们只声明一个成员对应的数据类型而不用指明成员的名字，匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。使用方法如下
	type Circle1 struct {
		Point
		Radius int
	}

	type Wheel1 struct {
		Circle1
		Spokes int
	}

	//得益于匿名嵌入的特性，我们可以直接访问叶子属性，而不用给出完整的路径，要注意的是完整路径依旧可以使用
	var w1 Wheel1
	w1.X = 0
	w1.Y = 0
	w1.Radius = 20

	//遗憾的是，结构体字面值并没有简单表示匿名成员的语法，结构体字面值必须遵循类型声明时的结构。以下两种语法等价：
	w2 := Wheel1{Circle1{Point{8, 8}, 6}, 20}
	w2 = Wheel1{
		Circle1: Circle1{
			Point:  Point{
				X: 2,
				Y: 4,
			},
			Radius: 43,
		},
		Spokes:  454,
	}
	fmt.Println(w2)
}
