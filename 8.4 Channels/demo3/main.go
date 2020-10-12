package main

/*
带缓存的channel内部持有一个元素队列。队列的最大容量是在调用make函数创建channel时通过第二个参数指定的。下面的语句创建了一个可以持有三个字符串元素的带缓存channel。
ch = make(chan string, 3)
向缓存channel的发送操作就是向内部缓存队列的尾部插入元素，接受操作就是从队列的头部删除元素。如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收
操作而释放了新的队列空间。相反，如果channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素。channel的缓存队列解耦了接收和发送的goroutine。

在某些特殊情况下，程序可能需要知道channel内部缓存的容量，可以用内置的cap函数获取：cap(ch)
同样，对于内置的len()函数，如果传入的是channel，那么将返回channel内部缓存队列中有效元素的个数。
*/
