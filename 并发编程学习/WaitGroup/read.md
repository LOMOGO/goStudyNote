![WaitGroup知识梳理](https://static001.geekbang.org/resource/image/84/ff/845yyf00c6db85c0yy59867e6de77dff.jpg)

## WaitGroup
#### WaitGroup要解决的是并发等待的问题：
现在有一个goroutine A 在检查点(checkpoint)等待一组goroutine全部完成，如果在执行任务的这些goroutine还没全部完成，那么goroutine A 就会阻塞在检查点，直到所有goroutine都完成后才能继续执行。
#### WaitGroup的基本用法：
Go标准库中的 WaitGroup 提供了三个方法：
```
    func (wg *WaitGroup) Add(delta int)
    func (wg *WaitGroup) Done()
    func (wg *WaitGroup) Wait()
```

- Add，用来设置 WaitGroup 的计数值； 
  
- Done，用来将 WaitGroup 的计数值减 1，其实就是调用了 Add(-1)；
  
- Wait，调用这个方法的 goroutine 会一直阻塞，直到 WaitGroup 的计数值变为 0。

#### WaitGroup的例子：
见demo.go

#### WaitGroup实现:
WaitGroup的数据结构：
```

type WaitGroup struct {
    // 避免复制使用的一个技巧，可以告诉vet工具违反了复制使用的规则
    noCopy noCopy
    // 64bit(8bytes)的值分成两段，高32bit是计数值，低32bit是waiter的计数
    // 另外32bit是用作信号量的
    // 因为64bit值的原子操作需要64bit对齐，但是32bit编译器不支持，所以数组中的元素在不同的架构中不一样，具体处理看下面的方法
    // 总之，会找到对齐的那64bit作为state，其余的32bit做信号量
    state1 [3]uint32
}


// 得到state的地址和信号量的地址
func (wg *WaitGroup) state() (statep *uint64, semap *uint32) {
    if uintptr(unsafe.Pointer(&wg.state1))%8 == 0 {
        // 如果地址是64bit对齐的，数组前两个元素做state，后一个元素做信号量
        return (*uint64)(unsafe.Pointer(&wg.state1)), &wg.state1[2]
    } else {
        // 如果地址是32bit对齐的，数组后两个元素用来做state，它可以用来做64bit的原子操作，第一个元素32bit用来做信号量
        return (*uint64)(unsafe.Pointer(&wg.state1[1])), &wg.state1[0]
    }
}
```
WaitGroup的数据结构中包含一个noCopy的辅助字段，一个state1记录WaitGroup状态的数组。

- noCopy的辅助字段，主要是辅助vet工具检查是否通过copy赋值这个WaitGroup实例
- state1,一个具有复合意义的字段，包含WaitGroup的计数、阻塞在检查点的waiter数和信号量。

因为对64位整数的原子操作要求整数的地址是64位对齐的，所以针对64位和32位环境的state字段的组成是不一样的。

在64位环境下，state1的第一个元素是waiter数，第二个元素是WaitGroup的计数值，第三个元素是信号量
![64位环境下 state1 各位表示的含义](https://static001.geekbang.org/resource/image/71/ea/71b5fyy6284140986d04c0b6f87aedea.jpg)
在32位环境下，如果state1不是64位对齐的地址，那么state1的第一个元素是信号量，后两个元素分别是waiter数和计数值。
![64位环境下 state1 各位表示的含义](https://static001.geekbang.org/resource/image/22/ac/22c40ac54cfeb53669a6ae39020c23ac.jpg)

接下来是Add、Done、Wait这三个方法的实现。

#### Add方法的逻辑：
Add方法主要操作的是state的技术部分。你可以为计数值增加一个delta值，内部通过原子操作把这个值加到计数值上。
这个delta也可以是个负数，相当于为计数值减去一个值，Done方法内部其实就是通过Add(-1)实现的。

Add方法的主干实现代码如下:
```
func (wg *WaitGroup) Add(delta int) {
    statep, semap := wg.state()
    // 高32bit是计数值v，所以把delta左移32，增加到计数上
    state := atomic.AddUint64(statep, uint64(delta)<<32)
    v := int32(state >> 32) // 当前计数值
    w := uint32(state) // waiter count

    if v > 0 || w == 0 {
        return
    }

    // 如果计数值v为0并且waiter的数量w不为0，那么state的值就是waiter的数量
    // 将waiter的数量设置为0，因为计数值v也是0,所以它们俩的组合*statep直接设置为0即可。此时需要并唤醒所有的waiter
    *statep = 0
    for ; w != 0; w-- {
        runtime_Semrelease(semap, false, 0)
    }
}
```

Done方法的实现代码如下：
```
// Done方法实际就是计数器减1
func (wg *WaitGroup) Done() {
    wg.Add(-1)
}
```

Wait方法的实现逻辑是：不断检查state的值。如果其中的计数值变成了0,那么说明所有的任务已经完成，调用者不必在等待，直接返回。如果计数值
大于0,说明此时还有任务没有完成，那么调用者就变成了等待者，需要加入waiter队列，并且阻塞自己。

Wait的主干实现代码如下：
```

func (wg *WaitGroup) Wait() {
    statep, semap := wg.state()
    
    for {
        state := atomic.LoadUint64(statep)
        v := int32(state >> 32) // 当前计数值
        w := uint32(state) // waiter的数量
        if v == 0 {
            // 如果计数值为0, 调用这个方法的goroutine不必再等待，继续执行它后面的逻辑即可
            return
        }
        // 否则把waiter数量加1。期间可能有并发调用Wait的情况，所以最外层使用了一个for循环
        if atomic.CompareAndSwapUint64(statep, state, state+1) {
            // 阻塞休眠等待
            runtime_Semacquire(semap)
            // 被唤醒，不再阻塞，返回
            return
        }
    }
}
```

#### 使用WaitGroup时的常见错误

###### 常见错误1：计数器设置为负值

WaitGroup的计数器的值必须大于等于0.我们在更改这个计数值的时候，WaitGroup会先做检查，如果计数值被设置为负数，就会导致panic。

常见的使WaitGroup的计数器变为零值的场景有：`调用Add的时候传递一个负数`和`调用Done方法的次数过多，超过了WaitGroup的计数值`。

**使用WaitGroup的正确姿势是：预先确定好WaitGroup的计数值，然后调用相同次数的Done完成相应的任务。**
比如在WaitGroup变量声明之后，就立即设置它的计数值，或者在goroutine启动之前增加1,然后调用Done。 

###### 常见错误2：不期望的Add时机
在使用WaitGroup的时候一定要遵循：等所有的Add方法调用之后再调用Wait，否则就可能导致panic或者不期望的结果。
见例子demo2.go

###### 常见错误3：前一个Wait还没结束就重用WaitGroup
因为WaitGroup是可以重用的，只要WaitGroup的计数值恢复到零值的状态，那么就可以被看作是新创建的Waitgroup，被重复使用.
错误代码如下：
```
func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        time.Sleep(time.Millisecond)
        wg.Done() // 计数器减1
        wg.Add(1) // 计数值加1
    }()
    wg.Wait() // 主goroutine等待，有可能和第7行并发执行
}
```

在这个例子中，第 6 行虽然让 WaitGroup 的计数恢复到 0，但是因为第 9 行有个 waiter 在等待，如果等待 Wait 的 goroutine，刚被唤醒就和 Add 调用（第 7 行）有并发执行的冲突，所以就会出现 panic。

- 总结一下：WaitGroup 虽然可以重用，但是是有一个前提的，那就是必须等到上一轮的 Wait 完成之后，才能重用 WaitGroup 执行下一轮的 Add/Wait，如果你在 Wait 还没执行完的时候就调用下一轮 Add 方法，就有可能出现 panic。

####noCopy： 辅助vet检查
我们刚刚在学习 WaitGroup 的数据结构时，提到了里面有一个 noCopy 字段。你还记得它的作用吗？其实，它就是指示 vet 工具在做检查的时候，这个数据结构不能做值复制使用。更严谨地说，是不能在第一次使用之后复制使用 ( must not be copied after first use)。

noCopy 是一个通用的计数技术，其他并发原语中也会用到。如果你想要自己定义的数据结构不被复制使用，或者说，不能通过 vet 工具检查出复制使用的报警，就可以通过嵌入 noCopy 这个数据类型来实现。

#### 关于如何避免错误使用 WaitGroup 的情况，我们只需要尽量保证下面 5 ：
- 不重用 WaitGroup。新建一个 WaitGroup 不会带来多大的资源开销，重用反而更容易出错。
- 保证所有的 Add 方法调用都在 Wait 之前。 
- 不传递负数给 Add 方法，只通过 Done 来给计数值减 1。
- 不做多余的 Done 方法调用，保证 Add 的计数值和 Done 方法调用的数量是一样的。
- 不遗漏 Done 方法的调用，否则会导致 Wait hang 住无法返回。


