* `os.Args`变量是一个字符串的切片
* golang中序列的元素数目为len(s)，Go语言的区间`索引同样采用左闭右开的形式`
* os.Args的第一个元素，os.Args[0]，是命令本身的名字；其它的元素则是程序启动时传给它的参数。
* 如果变量没有的显示初始化，则会被`隐式地赋予其类型的零值`，数值类型是0，`字符串类型是空字符串`""。
* golang中 `i++或i--是语句而不是表达式`，因此任何将他们赋值给其他变量的写法都是错误的，所以`j = i++非法`，并且++和--都只能放在变量名的后面，因此`--i也非法`。
* golang中只有for循环这一种循环语句。for循环有多种形式，其中一种如下所示：
```
for initialization; condition; post {
    // zero or more statements
}
```
for循环的这三个部分都可以省略(initialization; condition; post)
for循环的另外一种形式，在某种数据类型的区间（range）上遍历，如以下代码所展示：
```
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
```
每次循环迭代，range会产生一对值：索引以及在该索引处的元素值，range的语法要求，要处理元
素就必须处理索引。如果想省略索引，那么可以使用空标识符：_ 。

声明一个变量有好几种方式，以下这些都等价：
s := ""
var s string
var s = ""
var s string = ""
第一种是短变量声明，虽然简洁但是只能用在函数内部，而不能用于包变量。实践中一般采用前两种形式。