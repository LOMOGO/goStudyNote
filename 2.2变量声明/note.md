Go语言主要有四种类型的声明语句：var、const、type和func。分别对应变量、常量、类型和函数

变量声明： 

var 变量名字 类型 = 表达式

其中 “类型”或“= 表达式”两个部分可以省略其中一个。如果省略的是类型信息，那么将根据初始化表达式来推导变量的类型信息。如果是初始化表达式被省略，那么将用零值初始化该变量，其中：

* 数值变量对应的零值是0
* 布尔类型对应的零值是false
* 字符串类型对应的零值是空字符串
* 接口或引用类型（包括slice、map、chan和函数）变量的零值是 nil
也可以在一个声明语句中同时声明一组变量，或用一组初始化表达式声明并初始化一组变量。如果省略每个变量的类型，将可以声明多个不同类型的变量
```
var i, j, k int // int, int, int
var b, f, s = true, 2.3, "four"// bool, float64, string
```
简短变量声明：在函数内部，可以使用 “名字  := 表达式”形式声明变量。简短变量声明语句也可以用来声明和初始化一组变量： i, j := 0, 1; 但是需要注意一个比较微妙的地方是：简短变量声明左边的变量可能并不是全部都是刚刚声明的，如果有一些已经声明了，那么简短变量声明语句对这些已经声明过的变量就只有赋值行为了。`简短变量声明语句中必须至少要声明一个新的变量。`
﻿