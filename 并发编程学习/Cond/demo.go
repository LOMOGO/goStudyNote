package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{});
	var ready int
	rand.Seed(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(5)) * time.Second)

			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员#%d，已准备就绪", i)
			//广播唤醒所有等待者
			c.Signal()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	log.Println("所有运动员准备就绪，预备 3...2...1...开始！")
}
