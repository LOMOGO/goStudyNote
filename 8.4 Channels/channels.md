如果说goroutine是Go语言的并发体的话，那么channels则是它们之间的通信机制。一个channel
是一个通信机制。它可以让一个goroutine通过它给另一个goroutine发送值信息。每个channel都有
一个特殊的类型，一个int类型数据的channel一般写为chan int。

```
使用内置的make函数，创建一个channel：
ch := make(chan int)
```
两个相同类型的channel可以使用 == 运算符比较。如果两个channel引用的是相同的对象，那么比较
的结果为真。一个channel也可以和nil进行比较。

一个channel有发送和接收两个主要操作，都是通信行为。一个发送语句将一个值从一个goroutine
通过channel发送到另一个执行接收操作的goroutine。发送和接收两个操作都使用 `<-`运算符。
在发送语句中，`<-`运算符分割channel和要发送的值。在接收语句中，`<-`运算符写在channel
对象之前。一个不使用接收结果的操作也是合法的。
```
ch <- x
x = <- ch
<- ch
```
Channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic
异常。对一个已经被close过的channel进行接收操作依然可以接收到之前已经成功发送的数据；如果
channel中已经没有数据的话将产生一个零值的数据。

`close(ch)`

用make函数创建channel时可以指定第二个整型参数，对应channel的容量。如果channel容量大于
零，那么该channel就是带缓存的channel。
```
ch = make(chan int)
ch = make(chan int, 0)
ch = make(chan int, 3)
```