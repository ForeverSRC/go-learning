# 第八章 Goroutines和Channels

Go语言中的并发程序可以用两种手段来实现。goroutine和channel支持“顺序通信进程”(communicating sequential processes)或被简称为CSP

## 8.1 Goroutines

主函数所在的goroutine称为`main goroutine`

`main()`返回时，所有goroutine都会被直接打断，程序退出。

除了从主函数退出或者直接终止程序之外，没有其他编程方法能够让一个goroutine打断另一个goroutine的执行。

## 8.4 Channels

每个channel都有一个特殊的类型，也就是channels可发送数据的类型。

使用make函数可创建一个channel

```go
ch:=make(chan int)
```

channel的零值是nil。两个相同类型的channel可以使用==运算符比较。如果两个channel引用的是相同的对象，那 么比较的结果为真。一个channel也可以和nil进行比较。

一个channel有发送和接受两个主要操作，都是通信行为。

```go
ch <- x //send x to channel
x=<-ch  // receive from channerl to x
<-ch   //receive from channel and discard
```

Channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常。对一个已经被close过的channel进行接收操作依然可以接受到之前已经成功发送的数据;如果channel中已经没有数据的话将产生一个零值的数据。

```go
close(ch)
```

可在make时指定channel的缓存容量。

### 8.4.1 不带缓存的channel

一个基于无缓存Channel的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channel上执行接收操作，当发送的值通过Channels成功传输之后，两个goroutine可以 继续执行后面的语句。反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到有 另一个goroutine在相同的Channels上执行发送操作。

基于无缓存Channels的发送和接收操作将导致两个goroutine做一次**同步操作**。因此，无缓存Channels有时候也被称为同步Channels。当通过一个无缓存Channels发送数据时，接收者收到数据发生在唤醒发送者goroutine**之前(happens before)**

> 在讨论并发编程时，x事件在y事件之前发生(*happens before*)，并不是说x事件在时间上比y时间更早，要表达的意思是要保证在此之前的事件都已经完成了，例如在此之前的更新某些变量的操作已经完成，可以放心依赖这些已完成的事件了。
>
> x事件既不是在y事件之前发生也不是在y事件之后发生，则称x事件和y事件是**并发**的。这并不是意味着x事件和y事件就一定是同时发生的，只是不能确定这两个事件发生的先后顺序。

### 8.4.2 串联的Channels（Pipeline）

Channels也可以用于将多个goroutine连接在一起，一个Channel的输出作为下一个Channel的 输入。这种串联的Channels就是所谓的管道(pipeline)。

没有办法直接测试一个channel是否被关闭，但是接收操作有一个变体形式：多接收一个结果，多接收的第二个结果是一个布尔值ok：

* ture表示成功从channels接收到值
* false表示channels已经被关闭并且里面没有值可接收

### 8.4.3 单方向的Channel

```go
chan <- int //只发送int的channel，不能接收
<- chan int //只接收int的channel，不能发送
```

因为关闭操作只用于断言不再向channel发送新的数据，所以只有在发送者所在的goroutine才会调用close函数，因此对一个只接收的channel调用close将是一个编译错误。

### 8.4.4 带缓存的Channel

向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作而释放了新的队列空间。相反，如果channel是空的，接收操作将阻塞直到有另一个 goroutine执行发送操作而向队列插入元素。

```go
cap(ch) //查看通道缓存容量
len(ch) //查看有效元素个数
```

## 8.7 基于select的多路复用

```go
select { 
case <-ch1:
	// ...
case x := <-ch2: 
  // ...use x...
case ch3 <- y: 
  // ...
default: 
  // ...
}
```

每个case代表一个通信操作（在某个channel上进行发送或者接收）

select会等待case中有能够执行的case时去执行。当条件满足时，select才会去通信并执行 case之后的语句;这时候其它通信是不会执行的。一个没有任何case的select语句写作 select{}，会永远地等待下去

如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会

default来设置当其它的操作都不能够马上被处理时程序需要执行哪些逻辑

在select语句中操作nil的channel永远都不 会被select到。

## 8.9 并发的退出

