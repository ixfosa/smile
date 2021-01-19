package main

import (
	"fmt"
	"runtime"
	"time"
)

/******************* runtime.Gosched()********************/
/*
	让出CPU时间片，重新等待安排任务
 */
func main031() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s, "...")
		}
	}("smile")

	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("ixfosa")
		//smile ...
		//smile ...
		//ixfosa
		//ixfosa
	}

}

/******************runtime.Goexit()*********************/
/*
	退出当前协程
 */
func main032()  {
	go func() {
		defer fmt.Println("A.derfer")
		func() {
			defer fmt.Println("B.derfer")
			// 结束协程
			runtime.Goexit()

			defer fmt.Println("B.derfer")

			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	for {
		;
	}

	//B.derfer
	//A.derfer
}


/******************runtime.GOMAXPROCS*********************/
/*
Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。
默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。

Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。

Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。

Go语言中的操作系统线程和goroutine的关系：
	1.一个操作系统线程对应用户态多个goroutine。
	2.go程序可以同时使用多个操作系统线程。
	3.goroutine和OS线程是多对多的关系，即m:n。

我们可以通过将任务分配到不同的CPU逻辑核心上实现并行的效果，这里举个例子：
 */
func a()  {
	for i := 1; i <= 10; i++ {
		fmt.Println("a-", i)
	}
}

func b()  {
	for i := 1; i <= 10; i++ {
		fmt.Println("b-", i)
	}
}

func main() {
	//两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。
	//runtime.GOMAXPROCS(1)

	//将逻辑核心数设为2，此时两个任务并行执行
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}

