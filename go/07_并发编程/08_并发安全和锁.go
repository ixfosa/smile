package main

import (
	"fmt"
	"sync"
	"time"
)

/*******************并发安全和锁********************/
/*
有时候在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）。
类比现实生活中的例子有十字路口被各个方向的的汽车竞争；还有火车上的卫生间被车厢里的人竞争。
 */
var x int64
var wg081 sync.WaitGroup

func add()  {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg081.Done()
}
//上面的代码中我们开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符。
func main081() {
	wg081.Add(2)
	go add()
	go add()
	wg081.Wait()
	fmt.Println(x)

}

/******************互斥锁*********************/
/*
互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。

	使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
	当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。

Go语言中使用sync包的Mutex类型来实现互斥锁。 使用互斥锁来修复上面代码的问题：
 */
var y int64
var wg082 sync.WaitGroup
var lock sync.Mutex

func add082()  {
	for i := 0; i < 5000; i++ {
		lock.Lock()   // 加锁
		y = y + 1
		lock.Unlock() // 解锁
	}
	wg082.Done()
}

func main082()  {
	wg082.Add(2)

	go add082()
	go add082()

	wg082.Wait()

	fmt.Println(y)
}


/******************读写互斥锁*********************/
/*
互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，
这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型。

读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。

需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。
 */

var (
	x083 int64
	wg083 sync.WaitGroup
	lock083 sync.Mutex
	rwlock083 sync.RWMutex
)

func write083() {
	//lock083.Lock()  // 加互斥锁
	rwlock083.Lock() // 加写锁
	x083 = x083 + 1
	time.Sleep(time.Millisecond) // 假设读操作耗时10毫秒
	rwlock083.Unlock() // 解写锁
	//lock083.Unlock()  // 解互斥锁
	wg083.Done()
}

func read083()  {
	//lock083.Lock()
	rwlock083.Lock()
	time.Sleep(time.Millisecond)
	rwlock083.Unlock()
	//lock083.Unlock()
	wg083.Done()
}

func main()  {
	start := time.Now()

	for i := 0; i < 10; i++ {
		wg083.Add(1)
		go write083()
	}

	for i := 0; i < 1000; i++ {
		wg083.Add(1)
		go read083()
	}

	wg083.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}