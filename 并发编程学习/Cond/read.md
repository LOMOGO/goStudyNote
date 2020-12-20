## Cond
Go标准库提供Cond原语的目的是，为等待/通知场景下的并发问题提供支持。Cond通常应用于等待某个条件的一组goroutine，等条件变为true的时候，其中一个goroutine或者所有的goroutine都会被唤醒执行。

Cond是和某个条件相关，这个条件需要一组goroutine协作共同完成，在条件还没有满足的时候，所有等待这个条件的goroutine都会被阻塞住，只有这一组goroutine通过协作达到了这个条件，等待的goroutine才可能继续进行下去。

等待的条件，可以是某个变量达到了某个阈值或者某个时间点，也可以是一组变量分别都达到了某个阈值，还可以是某个对象的状态满足了特定的条件。总结来讲，等待的条件是一种可以用来计算结果是 true 还是 false 的条件。

从开发实践上，我们真正使用 Cond 的场景比较少，因为一旦遇到需要使用 Cond 的场景，我们更多地会使用 Channel 的方式

标准库中的Cond并发原语**初始化**的时候，需要关联一个Locker接口的实例，一般我们用**Mutex或者RWMutex**。

Cond的实现

```
type Cond

func NewCond(l locker) *Cond
func (c *Cond) Broadcast()
func (c *Cond) Signal()
func (c *Cond) Wait()
```

首先，Cond关联的Locker实例可以通过c.L访问，它内部维护着一个先入先出的等待队列

**Signal 方法**，允许调用者 Caller 唤醒一个等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，显然无需通知 waiter；如果 Cond 等待队列中有一个或者多个等待的 goroutine，则需要从等待队列中移除第一个 goroutine 并把它唤醒。在其他编程语言中，比如 Java 语言中，Signal 方法也被叫做 notify 方法。

`调用Signal方法时,不强求一定要持有c.L的锁`

**Broadcast 方法**，允许调用者 Caller 唤醒所有等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，显然无需通知 waiter；如果 Cond 等待队列中有一个或者多个等待的 goroutine，则清空所有等待的 goroutine，并全部唤醒。在其他编程语言中，比如 Java 语言中，Broadcast 方法也被叫做 notifyAll 方法。

`同样地，调用 Broadcast 方法时，也不强求你一定持有 c.L 的锁。`

**Wait 方法**，会把调用者 Caller 放入 Cond 的等待队列中并阻塞，直到被 Signal 或者 Broadcast 的方法从等待队列中移除并唤醒。

`调用Wait方法时必须要持有c.L的锁。`

有关Cond用法的例子见`demo.go`

## Cond的源码

```

type Cond struct {
    noCopy noCopy

    // 当观察或者修改等待条件的时候需要加锁
    L Locker

    // 等待队列
    notify  notifyList
    checker copyChecker
}

func NewCond(l Locker) *Cond {
    return &Cond{L: l}
}

func (c *Cond) Wait() {
    c.checker.check()
    // 增加到等待队列中
    t := runtime_notifyListAdd(&c.notify)
    c.L.Unlock()
    // 阻塞休眠直到被唤醒
    runtime_notifyListWait(&c.notify, t)
    c.L.Lock()
}

func (c *Cond) Signal() {
    c.checker.check()
    runtime_notifyListNotifyOne(&c.notify)
}

func (c *Cond) Broadcast() {
    c.checker.check()
    runtime_notifyListNotifyAll(&c.notify）
}
```

runtime_notifyListXXX 是运行时实现的方法，实现了一个等待 / 通知的队列。如果你想深入学习这部分，可以再去看看 runtime/sema.go 代码中。

copyChecker 是一个辅助结构，可以在运行时检查 Cond 是否被复制使用。

Signal 和 Broadcast 只涉及到 notifyList 数据结构，不涉及到锁。

Wait 把调用者加入到等待队列时会释放锁，在被唤醒之后还会请求锁。在阻塞休眠期间，调用者是不持有锁的，这样能让其他 goroutine 有机会检查或者更新等待变量。

## 使用Cond的2个常用错误

- 调用Wait的时候没有加锁
- 只调用了一次 Wait，没有检查等待条件是否满足，结果条件没满足，程序就继续执行了
