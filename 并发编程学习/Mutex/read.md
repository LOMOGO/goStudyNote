##  Mutex需要注意的地方
Mutex的Unlock方法可以被任意的goroutine调用释放锁，即使是没持有这个互斥锁的goroutine也可以进行这个操作。这是因为，Mutex本身并没有包含持有这把锁的goroutine的信息，所以，Unlock也不会对此进行检查。

这个功能有时候就很危险了，因为其他goroutine可以强制释放锁，这是一个非常危险的操作，因为在临界区的goroutine可能不知道锁已经被释放了，可能还有执行临界区的业务操作，这可能会带来意想不到的结果，因为这个goroutine还以为自己持有锁，因此有可能导致data race问题。

所以我们在使用Mutex的时候，必须要保证goroutine尽可能不去释放自己未持有的锁，一定要遵循`谁申请，谁释放`的原则，在真实的实践中，我们使用互斥锁的时候，很少在一个方法中单独申请锁，而在另外一个方法中单独释放锁，一般都会在同一个方法中获取锁、释放锁。

以前会基于性能的考虑，及时释放掉锁，所以在一些if-else分支中加上释放锁的代码，从go1.14版本起，go对defer进行了优化，defer对于性能的影响很小，所以使用defer mu.Unlock()来替换掉用if-else这样的分支执行mu.unlock() 也是没有问题的。

## Mutex 四种常见易错点
- Lock和Unlock不是成对出现
- Copy已经使用了的Mutex
- 重入
- 死锁

保证Lock/Unlock成对出现，尽可能的使用 defer mutex.Unlock的方式把他们成对、紧凑的写在一起