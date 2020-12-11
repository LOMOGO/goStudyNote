package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//例子1
	/*var wg sync.WaitGroup
	go exp1(200, &wg)
	go exp1(400, &wg)
	go exp1(800, &wg)
	wg.Wait()*/

	//例子2
	var wg sync.WaitGroup
	//预先设置计数值
	wg.Add(3)
	go exp2(200, &wg)
	go exp2(400, &wg)
	go exp2(800, &wg)
	wg.Wait()
	fmt.Println()

	//例子3
	exp3(200, &wg)
	exp3(400, &wg)
	exp3(800, &wg)
	wg.Wait()
	fmt.Println("Done")
}

//例子1本来是想要在所有goroutine执行完毕之后输出持续时间duration，但是这段代码的错误之处就在于：将WaitGroup.Add()方法放在子goroutine中
//等主goroutine调用Wait的时候，因为三个子goroutine一开始都休眠，所以可能WaitGroup的Add方法还没有被调用，WaitGroup的计数还是0,所以他并没有等待三个
//子goroutine都执行完毕才继续执行，而是立即执行了下一步。
//导致这个错误的原因是：没有遵循先完成所有的Add之后才Wait。
func exp1(millisecs time.Duration, wg *sync.WaitGroup) {
	duration := millisecs * time.Millisecond
	time.Sleep(duration)

	wg.Add(1)
	fmt.Printf("例子1：程序持续时间为：%v\n", duration)
	wg.Done()
}

//要解决例子1的问题，一个方法是，预先设置计数值
func exp2(millisecs time.Duration, wg *sync.WaitGroup)  {
	duration := millisecs * time.Millisecond
	time.Sleep(duration)
	fmt.Printf("例子2：程序持续时间为：%v\n", duration)
	wg.Done()
}

//另外一个方法是：在启动子goroutine之前才调用Add方法
func exp3(millisecs time.Duration, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		duration := millisecs * time.Millisecond
		time.Sleep(duration)
		fmt.Printf("例子3：程序持续时间为：%v\n", duration)
		wg.Done()
	}()
}
