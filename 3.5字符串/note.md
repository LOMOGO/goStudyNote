```
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
}
```
字符串值也可以用`字符串面值方式编写，只要将一系列字节序列包含在双引号内`即可：

"hello, 世界"

因为golang的文本字符串以UFT8编码的方式处理，因此我们可以将Unicode码点也写到字符串面值中。

在一个双引号字符串包含的字符串面值中，可以用反斜杠\开头的转义序列插入任意的数据。如下所示：

| 字符 | 含义 |
| ---- | ---- |
| \a |       响铃 | 
| \b |      退格 | 
| \f |      换页 | 
| \n |     换行 | 
| \r  |     回车 | 
| \t |      制表符 | 
| \v  |     垂直制表符 | 
| \'   |    单引号 (只用在 '\'' 形式的rune符号面值中) | 
| \"   |    双引号 (只用在 "..." 形式的字符串面值中) | 
| \\    |   反斜杠 | 

一个`原生的字符串面值形式是`\`...\`，使用反引号代替双引号。在原生的字符串面值中，没有转义操作；

全部的内容都是字面的意思，包含退格和换行，但是无法直接写`如何要写这个需要字符串拼接+"`"即可。

原生字符串面值被广泛应用于HTML模板、json面值、命令行提示信息以及那些需要拓展到多行的场景。

一个字符串是包含只读字节的数组，一旦创建，是不可变的。而一个 字节slice的元素则可以自由的修改。

**字符串和字节slice之间可以相互转换：**
`s := "abc"
b := []byte(s)
s2 := string(b)`

为避免不必要的内存分配，bytes包和string包都提供了很多实用的函数，
比如string包中的六个函数：
``` 
 func Contains(s, substr string) bool
 func Count(s, sep string) int
 func Fields(s string) []string
 func HasPrefix(s, prefix string) bool
 func Index(s, sep string) int
 func Join(a []string, sep string) string
```
bytes包中对应的六个函数：
```
func Contains(b, subslice []byte) bool
func Count(s, sep []byte) int
func Fields(s []byte) [][]byte
func HasPrefix(s, prefix []byte) bool
func Index(s, sep []byte) int
func Join(s [][]byte, sep []byte) []byte
```
它们之间唯一的区别就是字符串类型的参数被替换成了字节slice类型的参数

字符串和数字的转换：

将整数转换为字符串，一种方法使用fmt.Sprintf返回一个格式化的字符串；另一种方法是用strconv.Itoa()

fmt.Sprintf函数的%b、%d、%o和%x等参数可以很方便的进行进制之间的转换。

如果要将一个字符串解析为整数，可以使用strconv包的Atoi或ParseInt函数，还有用于解析无符号整数的ParseUnit函数
```
x, err := strconv.Atoi("123")             // x is an int
y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
```
ParseInt函数的第三个参数是用于指定整型数的大小，例如16表示int16，0则表示int。在任何情况下y总是int64类型

