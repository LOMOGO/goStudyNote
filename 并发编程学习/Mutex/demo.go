package main

import (
	"fmt"
	"sync"
)

//go语言中的互斥锁就是Mutex，互斥锁是并发控制的一个基本手段，是为了避免竞争而建立的一种并发控制机制。

//临界区：在并发编程中，如果程序中的一部分会被并发访问或修改，那么，为了避免并发访问导致的意想不到的结果，这部分
//程序需要被保护起来，这部分保护起来的程序，就叫做临界区。可以说临界区就是一个被共享的资源，比如对数据库的访问，
//对某一个共享数据结构的操作、对一个I/O设备的使用。。。。。。

//如果很多线程同步访问临界区，就会造成访问或操作错位，所以我们可以使用互斥锁，限定临界区只能同时由一个线程持有。

//Mutex、channel可以叫做同步原语或者并发原语，关于同步原语和并发原语没有一个严格的定义，可以把它看作解决并发问题的一个
//基础的数据结构。

//同步原语的使用场景：共享资源   任务编排   消息传递

//创建十个goroutine，每个goroutine执行十万次加1操作，最后期望公共变量count的值为一百万
func Exm1() {
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

func Exm2() {
	//不需要对Mutex的零值进行额外的初始化，直接使用变量var mu sync.Mutex即可。
	var mu sync.Mutex
	var count = 0

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				//这个例子中的共享资源是count变量，临界区是count++，只要在临界区前面获取锁
				//离开临界区的时候释放锁，就可以解决数据竞争的问题了。
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

//很多情况下，Mutex会嵌入到其他struct中使用，假如有多个字段，一般会把Mutex放在要控制的字段上面。
type Counter struct {
	mu sync.Mutex
	count int
}

func Exm3() {
	var counter Counter
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.mu.Lock()
				counter.count++
				counter.mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
}

func main() {
	/*Exm1每次执行的结果并不一定相同，造成这种结果的原因是count++不是一个原子操作，它至少包含
	读取变量count的当前值，对这个值加1，再把结果写回到count变量中，因为不是原子操作，就可能有并发的问题。
	比如10个goroutine同时读取到count的值为9527,接着各自按照自己的逻辑加1,值变成了9528,然后把这个结果再写会到
	count变量中。但是，实际上，此时我们增加的总数应该是10才对，这里却只增加了1,这就是并发访问共享
	数据的常见错误。*/
	Exm1()
	Exm2()
	Exm3()
}