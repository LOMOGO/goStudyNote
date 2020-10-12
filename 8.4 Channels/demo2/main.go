/*Go语言的类型系统提供了单方向的channel类型，分别用于只发送或只接收的channel。(channel，中文翻译就是管道，我们可以把这个channel就想象成
一个管道，管道的左边代表数据发送进来，管道的右边代表数据即将被接收，那么即可以发送又可以接收的管道可以理解成双向的channel。管道的右边被堵住的
话那么这个管道就是只能发送的channel，也就是单向channel，左边被堵住的话那这个channel就是能接收的channel，也是一个单方向的channel。但是说他是
单方向的并不意味着只能发送的单方向channel中的数据不能被获取，他里面的数据依旧可以被channel获取到。)*/
package main

import "fmt"

func main()  {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
	//因为关闭操作只用于断言不再向channel发送新的数据，所以只有在发送者所在的goroutine才会调用close函数，因此对一个只接收的channel调用close将会是一个编译错误。
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		//类型chan<- int表示一个只发送int的channel, <-chan int表示一个只接收int的channel。
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
