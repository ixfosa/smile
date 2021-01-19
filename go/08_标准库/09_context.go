package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*******************上下文 Context********************/
/*
	在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理。
请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和RPC服务。
用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的token、请求的截止时间。
当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源。

上下文 context.Context 是用来设置截止日期、同步信号，传递请求相关值的结构体。上下文与 Goroutine 有比较密切的关系。

context.Context 是 Go 语言在 1.7 版本中引入标准库的接口，该接口定义了四个需要实现的方法，其中包括：
	1.Deadline 	— 返回 context.Context 被取消的时间，也就是完成工作的截止日期；
	2.Done 		— 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消之后关闭，多次调用 Done 方法会返回同一个 Channel；
	3.Err 		— 返回 context.Context 结束的原因，它只会在 Done 返回的 Channel 被关闭时才会返回非空的值；
		1.如果 context.Context 被取消，会返回 Canceled 错误；
		1.如果 context.Context 超时，会返回 DeadlineExceeded 错误；
	4.Value 	— 从 context.Context 中获取键对应的值，对于同一个上下文来说，
	              多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；

	type Context interface {
		Deadline() (deadline time.Time, ok bool)
		Done() <-chan struct{}
		Err() error
		Value(key interface{}) interface{}
	}

context 包中提供的 context.Background、context.TODO、context.WithDeadline 和 context.WithValue 函数会返回实现该接口的私有结构体
 */


/******************为什么需要Context*********************/
/*
我们可能会创建多个 Goroutine 来处理一次请求，
而 context.Context 的作用就是在不同 Goroutine 之间同步请求特定数据、取消信号以及处理请求的截止日期。

每一个 context.Context 都会从最顶层的 Goroutine 一层一层传递到最下层。
context.Context 可以在上层 Goroutine 执行出现错误时，将信号及时同步给下层。

不使用 Context 同步信号
	当最上层的 Goroutine 因为某些原因执行失败时，下层的 Goroutine 由于没有接收到这个信号所以会继续工作；
	但是当我们正确地使用 context.Context 时，就可以在下层及时停掉无用的工作以减少额外资源的消耗：

使用 Context 同步信号
 */

func handle(ctx context.Context, duration time.Duration)  {
	select {
	case <- ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <- time.After(duration):
		fmt.Println("process request with", duration)
	}
}
// 创建了一个过期时间为 1s 的上下文，并向上下文传入 handle 函数，该方法会使用 500ms 的时间处理传入的『请求』
func main091() {
	// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)

	defer cancel()

	//go handle(ctx, 500 * time.Millisecond)

	//将处理『请求』时间增加至 1500ms，整个程序都会因为上下文的过期而被中止，：
	// handle context deadline exceeded
	// main context deadline exceeded
	go handle(ctx, 1500 * time.Millisecond)

	select {
	case <- ctx.Done():
		fmt.Println("main", ctx.Err())
	}
	// 因为过期时间大于处理时间，所以我们有足够的时间处理该『请求』，运行上述代码会打印出如下所示的内容：
	// process request with 500ms
	// main context deadline exceeded
	// handle 函数没有进入超时的 select 分支，但是 main 函数的 select 却会等待 context.Context 的超时
	// 并打印出 main context deadline exceeded。
}

// 初始的例子
var wg092 sync.WaitGroup
func worker92()  {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	// 如何接收外部命令实现退出
	wg092.Done()
}
func main092() {
	wg092.Add(1)
	go worker92()

	// 如何实现结束子goroutine
	wg092.Wait()
	fmt.Println("over!!!")
}

/***************************************/

// 全局变量方式
var wg093 sync.WaitGroup
var exit093 bool

// 全局变量方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易统一
// 2. 如果worker中再启动goroutine，就不太好控制了。

func worker93()  {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)

		if exit093 {
			break
		}
	}
	// 如何接收外部命令实现退出
	wg093.Done()
}

func main093()  {
	wg093.Add(1)
	go worker93()
	time.Sleep(3 * time.Second) // sleep3秒以免程序过快退出
	exit093 = true   // 修改全局变量实现子goroutine的退出
	wg093.Wait()
	fmt.Println("over!!!")
}

/***************************************/
// 通道方式

var wg094 sync.WaitGroup

// 管道方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel

func worker94(exitChan chan struct{})  {
	LOOP:
		for {
			fmt.Println("worker")
			time.Sleep(time.Second)
			select {
			case <- exitChan:
				break LOOP
			default:

			}
		}
		// 如何接收外部命令实现退出
		wg094.Done()
}

func main094()  {
	exitChan := make(chan struct{})
	wg094.Add(1)
	go worker94(exitChan)
	time.Sleep(3 * time.Second) // sleep3秒以免程序过快退出
	close(exitChan)  // 给子goroutine发送退出信号
	wg094.Wait()
	fmt.Println("over!!!")
}



/***************************************/
// 官方版的方案
var wg095 sync.WaitGroup

func worker95(ctx context.Context)  {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <- ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	// 如何接收外部命令实现退出
	wg095.Done()
}

func main095()  {
	// func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
	ctx, cancel := context.WithCancel(context.Background())
	wg095.Add(1)
	go worker95(ctx)
	time.Sleep(3 * time.Second)
	cancel() // 通知子goroutine结束
	wg095.Wait()
	fmt.Println("over")
}

// 当子goroutine又开启另外一个goroutine时，只需要将ctx传入即可：
var wg096 sync.WaitGroup

func worker096(ctx context.Context)  {
	worker966(ctx)
	LOOP:
		for {
			fmt.Println("worker96")
			time.Sleep(time.Second)
			select {
			case <- ctx.Done(): // 等待上级通知
				break LOOP
			default:
			}
		}
		// 如何接收外部命令实现退出
		wg096.Done()
}

func worker966(ctx context.Context)  {
LOOP:
	for {
		fmt.Println("worker966")
		time.Sleep(time.Second)
		select {
		case <- ctx.Done(): // 等待上级通知
			break LOOP
		default:

		}
	}
}

func main096() {
	ctx, cancel := context.WithCancel(context.Background())
	wg096.Add(1)
	go worker096(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg096.Wait()
	fmt.Println("over")
}


/********************默认上下文*******************/
/*
	context 包中最常用的方法还是 context.Background、context.TODO，
这两个方法都会返回预先初始化好的私有变量 background 和 todo，它们会在同一个 Go 程序中被复用：

	Go内置两个函数：Background()和 TODO()，这两个函数分别返回一个实现了 Context 接口的 background 和 todo。
我们代码中最开始都是以这两个内置的上下文对象作为最顶层的partent context，衍生出更多的子上下文对象。
	func Background() Context {
		return background
	}

	func TODO() Context {
		return todo
	}

这两个私有变量都是通过 new(emptyCtx) 语句初始化的，它们是指向私有结构体 context.emptyCtx 的指针，这是最简单、最常用的上下文类型：
	var (
		background = new(emptyCtx)
		todo       = new(emptyCtx)
	)

	type emptyCtx int

	func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
		return
	}

	func (*emptyCtx) Done() <-chan struct{} {
		return nil
	}

	func (*emptyCtx) Err() error {
		return nil
	}

	func (*emptyCtx) Value(key interface{}) interface{} {
		return nil
	}

context.emptyCtx 通过返回 nil 实现了 context.Context 接口，它没有任何特殊的功能。

Context 层级关

	background和todo本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context，它们只是在使用和语义上稍有不同：

	context.Background 是上下文的默认值，所有其他的上下文都应该从它衍生（Derived）出来；
	context.TODO 应该只在不确定应该使用哪种上下文时使用；

在多数情况下，如果当前函数没有上下文作为入参，我们都会使用 context.Background 作为起始的上下文向下传递。
 */



/*******************With系列函数--WithCancel********************/
/*
WithCancel的函数签名如下：
    func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
WithCancel返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，
无论先发生什么情况。

取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。
 */


// gen函数在单独的goroutine中生成整数并将它们发送到返回的通道。
// gen的调用者在使用生成的整数之后需要取消上下文，以免gen启动的内部goroutine发生泄漏。
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <- ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func main097()  {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

/******************With系列函数--WithDeadline*********************/
/*
WithDeadline的函数签名如下：
	func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
返回父上下文的副本，并将deadline调整为不迟于d。如果父上下文的deadline(截至时间)已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。
当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。

取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。
 */

func main098() {
	// 定义了一个50毫秒之后过期的deadline
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	// 使用一个select让主程序陷入等待：等待1秒后打印overslept退出或者等待ctx过期后退出。
	select {
	case <- time.After(1 * time.Second):
		fmt.Println("overslept")
	case <- ctx.Done():  //因为 ctx-50秒后就过期，所以ctx.Done()会先接收到值，上面的代码会打印ctx.Err()取消原因。
		fmt.Println(ctx.Err())
	}
}


/*****************With系列函数--WithTimeout**********************/
/*
WithTimeout的函数签名如下：
	func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
WithTimeout返回WithDeadline(parent, time.Now().Add(timeout))。

取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制。
 */

var wg099 sync.WaitGroup

func worker099(ctx context.Context)  {
	LOOP:
		for {
			fmt.Println("db connecting ...")
			time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
			select {
			case <- ctx.Done():
				break LOOP
			default:
			}
		}
	fmt.Println("worker done!")
	wg099.Done()
}
func main099()  {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), 50 * time.Millisecond)

	wg099.Add(1)

	go worker099(ctx)

	//time.Sleep(5 * time.Second)

	cancel() // 通知子goroutine结束

	wg099.Wait()
	fmt.Println("over")
}

/*****************With系列函数--WithValue**********************/
/*
WithValue函数能够将请求作用域的数据与 Context 对象建立关系。声明如下：
    func WithValue(parent Context, key, val interface{}) Context
WithValue返回父节点的副本，其中与key关联的值为val。

仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。

	所提供的键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。
WithValue的用户应该为键定义自己的类型。为了避免在分配给interface{}时进行分配，上下文键通常具有具体类型struct{}。
或者，导出的上下文关键变量的静态类型应该是指针或接口。
 */

type TraceCode string

var wg0910 sync.WaitGroup

func work0910(ctx context.Context)  {
	key := TraceCode("TRACE_CODE")
	// Value(key interface{}) interface{}
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code

	if !ok {
		fmt.Println("invalid trace code")
	}

	LOOP:
		for {
			fmt.Printf("worker, trace code:%s\n", traceCode)
			time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
			select {
			case <- ctx.Done(): // 50毫秒后自动调用
				break LOOP
			default:
			}
		}
	fmt.Println("worker done!")
	wg0910.Done()
}

func main()  {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 50)

	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "ixfosa")

	wg0910.Add(1)

	go work0910(ctx)

	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg0910.Wait()
	fmt.Println("over")
}

/******************使用Context的注意事项*********************/
/*
	推荐以参数的方式显示传递Context
	以Context作为参数的函数方法，应该把Context作为第一个参数。
	给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
	Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
	Context是线程安全的，可以放心的在多个goroutine中传递
 */


