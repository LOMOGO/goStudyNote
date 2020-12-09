package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu sync.Mutex
}

func (c *Counter) Write() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

//注意读取的时候也要对临界区进行互斥锁保护
func (c *Counter) Read() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

//注意这里传的是指针，如果不传指针的话互斥锁就会出现复制错误
func Worker(c *Counter, w *sync.WaitGroup) {
	defer w.Done()
	c.Write()
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Worker(&counter, &wg)
	}
	wg.Wait()
	fmt.Println(counter.Read())
}
