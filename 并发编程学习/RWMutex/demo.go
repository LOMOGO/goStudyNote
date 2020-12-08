//本来写下面这个demo是想比较一下Mutex和RWMutex的性能对比来着，后来发现我这个例子没办法进行两者的性能对比测试，就当是练手了吧，自行忽略
package main

import (
	"fmt"
	"sync"
	"time"
)

/*
我们可以使用Mutex来保证读写共享资源的安全性。不管是读还是写，我们都通过 Mutex 来保证只有一个 goroutine 访问共享资源，
这在某些情况下有点“浪费”。比如说，在写少读多的情况下，即使一段时间内没有写操作，大量并发的读访问也不得不在 Mutex 的保护下变成了串行访问，
这个时候，使用 Mutex，对性能的影响就比较大。*/

/*如果可以区分读写操作的话，那么就可以提高性能，解释如下：
如果某个读操作的goroutine持有了锁，在这种情况下，其它读操作的goroutine就不必一直傻傻的等待了
而是可以并发地访问共享变量，这样我们就可以将原本Mutex下的串行的读变成并行的读，从而提高读操作的性能。
当写操作的goroutine持有锁的时候，他就是一个排外锁，其它的写操作和读操作的goroutine，需要阻塞等待持有这个
锁的goroutine释放锁。
说白了就是RWMutex,写操作的goroutine占有的状态下其他goroutine不能读写，读操作的goroutine占有的状态下允许读操作的goroutine并发的读，不允许写操作的goroutine写
*/

//这一类并发读写问题叫作readers-writers 问题，意思就是，同时可能有多个读或者多个写，但是只要有一个线程在执行写操作，其它的线程都不能执行读写操作。

type MuCounter struct {
	count int
	mu sync.Mutex
}

type RWMuCounter struct {
	count int
	mu sync.RWMutex
}

/*
RWMutex的方法也很少，总共只有5个
- Lock/Unlock: 写操作时调用的方法。  如果锁已经被reader或者writer持有，那么，Lock方法会一直
阻塞，直到能获取到锁；Unlock则是配对的释放锁的方法。
- RLock/RUnlock: 读操作时调用的方法。  如果锁已经被writer持有的话，Rlock方法会一致阻塞，直到
能获取到锁，否则就直接返回；而RUnlock是reader释放锁的方法。
- Rlocker: 这个方法的作用是为读操作返回一个Locker接口的对象。它的Lock方法会调用RWMutex的Rlock
方法，它的Unlock方法会调用RWMutex的RUnlock方法。
*/

/*
- 什么时候考虑使用RWMutex：
如果遇到明确区分reader和writer goroutine的场景，且有大量的并发读、少量的并发写，并且有强烈的
性能需求，就可以考虑使用读写所RWMutex替换Mutex。

- Go语言的RWMutex设计时优先是写操作的goroutine优先

- RWMutex的3个踩坑点：
1：不可复制（传指针）
2：重入导致死锁
3：释放未加锁的RWMutex
*/

func main() {
	var mc MuCounter
	var rwmc RWMuCounter
	for i := 0; i < 10; i++ {
		go func() {
			for {
				fmt.Println("Mutex: ", mc.Read())
				fmt.Println("RWMutex: ---------------", rwmc.Read())
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for {
		mc.Write()
		rwmc.Write()
		time.Sleep(time.Second)
	}
}

//Mutex写
func (mc *MuCounter) Write() {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.count++
}

//Mutex读
func (mc *MuCounter) Read() int {
	return mc.count
}

//RWMutex写
func (rwmc *RWMuCounter) Write() {
	rwmc.mu.Lock()
	defer rwmc.mu.Unlock()
	rwmc.count++
}

//RWMutex写
func (rwmc *RWMuCounter) Read() int {
	rwmc.mu.RLock()
	defer rwmc.mu.RUnlock()
	return rwmc.count
}
