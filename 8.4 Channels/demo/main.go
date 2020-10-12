//channels也可以用于将多个goroutine连接在一起，一个channel的输出作为下一个channel的输入。这种串联的channels就是所谓的管道(pipeline)
package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x <= 100; x++ {
			naturals <- x
		}
		/*当一个channel被关闭后，再向这个channel发送数据将导致panic异常。当一个被关闭的channel中已发送的数据都被成功接收后，
		后续的接受操作将不再阻塞，它们会立即返回一个零值。
		* 其实不需要关闭每一个channel。只有当需要告诉接收者goroutine，所有的数据已经 全部发送 时才需要关闭channel。不管一个channel是非
		被关闭，当然没有被引用时将会被Go的垃圾自动回收器回收。(不要将关闭一个打开文件的操作和关闭一个channel操作混淆。对于每个打开的文件
		，都需要在不使用的时候调用对应的Close方法来关闭文件。)
		* 试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常。关闭一个channel将会触发一个广播机制。*/
		close(naturals)
	}()
	//squares
	go func() {
		//go语言的range循环可以直接在channels上面迭代，它依次从channel接收数据，当channel被关闭并且没有值可以接收时跳出循环。
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	//上面的这个函数等效于：
	/*
	//没有办法直接测试一个channel是否被关闭，但是接收操作有一个变体形式： 当它接收两个结果的时候，多接收的第二个结果是一个布尔值ok，
	//true表示成功从channels接收到值，false表示channels已经被关闭并且里面没有值可接收。
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()
	*/

	for x := range squares{
		fmt.Println(x)
	}

}
