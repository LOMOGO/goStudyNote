//如果你的代码的循环结构中使用了函数值，且函数值需要使用循环变量的话，那么你就需要注意了
package main

import "fmt"

/*
var rmdirs []func()
//处理这类问题的时候你需要引入一个局部变量并将循环变量的值赋值给局部变量，然后才能在函数值中使用局部变量，如果不进行赋值操作的话，那么函数值操作的值都是循环变量最后一次迭代的值，因为它们都会等待循环结束后再执行函数值
for _, d := range tempDirs() {
	dir := d
	os.MkdirAll(dir, 0755)
	rmdirs = append(rmdirs, func() {
		os.RemoveAll(dir)
	})
}
for _, rmdir := range rmdirs{
	rmdir()
}

*/

func main() {
	slice_func := []func(){}
	for i := 0; i < 5; i++ {
		ni := i
		slice_func = append(slice_func, func() {
			fmt.Println(ni)
		})
	}
	for _, v := range slice_func {
		v()
	}
}
