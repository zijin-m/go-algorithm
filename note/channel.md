# 数据结构

```go
type hchan struct {
    qcount   uint           // 当前队列中剩余元素个数
    dataqsiz uint           // 环形队列长度，即可以存放的元素个数
    buf      unsafe.Pointer // 环形队列指针
    elemsize uint16         // 每个元素的大小
    closed   uint32            // 标识关闭状态
    elemtype *_type         // 元素类型
    sendx    uint           // 队列下标，指示元素写入时存放到队列中的位置
    recvx    uint           // 队列下标，指示元素从队列的该位置读出
    recvq    waitq          // 等待读消息的goroutine队列
    sendq    waitq          // 等待写消息的goroutine队列
    lock mutex              // 互斥锁，chan不允许并发读写
}
```

## 等待队列
从channel读数据，如果channel缓冲区为空或者没有缓冲区，当前goroutine会被阻塞。 向channel写数据，如果channel缓冲区已满或者没有缓冲区，当前goroutine会被阻塞。

被阻塞的goroutine将会挂在channel的等待队列中：

* 因读阻塞的goroutine会被向channel写入数据的goroutine唤醒；
* 因写阻塞的goroutine会被从channel读数据的goroutine唤醒

# Channel 读写
##  创建channel
创建channel的过程实际上是初始化hchan结构。其中类型信息和缓冲区长度由make语句传入，buf的大小则与元素大小和缓冲区长度共同决定。

创建channel的伪代码如下所示：

```go
func makechan(t *chantype, size int) *hchan {
    var c *hchan
    c = new(hchan)
    c.buf = malloc(元素类型大小*size)
    c.elemsize = 元素类型大小
    c.elemtype = 元素类型
    c.dataqsiz = size

    return c
}
```

# 写

```go
if recvq 不为空
    取出一个 G
    写入数据到 G
    唤醒 G
    return
if buf 缓冲已满
    加入 sendq，等待被唤醒
else 写入到 buf 尾部
```

# 读
从一个channel读数据简单过程如下：

* 如果等待发送队列sendq不为空，且没有缓冲区，直接从sendq中取出G，把G中数据读出，最后把G唤醒，结束读取过程；
* *果等待发送队列sendq不为空，此时说明缓冲区已满，从缓冲区中首部读出数据，把G中数据写入缓冲区尾部，把G唤醒，结束读取过程；
* 如果缓冲区中有数据，则从缓冲区取出数据，结束读取过程；
* 将当前goroutine加入recvq，进入睡眠，等待被写goroutine唤醒；

# 用法

## select
使用select可以监控多channel，比如监控多个channel，当其中某一个channel有数据时，就从其读出数据

## range
通过range可以持续从channel中读出数据，好像在遍历一个数组一样，当channel中没有数据时会阻塞当前goroutine，与读channel时阻塞处理机制一样。