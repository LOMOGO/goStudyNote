* golang搭建的的服务器每接收一次请求处理时都会另起一个goroutine，这样服务器就可以同一时间处理多个请求
* golang的if语句支持嵌套如以下代码所示
```
if err := r.PraseForm err != nil {
    log.Print(err)
}
```
golang`允许这样的一个简单的语句结果作为局部的变量声明出现在if语句的最前面`，这样err这个变量的作用域只在if语句这个范围内。