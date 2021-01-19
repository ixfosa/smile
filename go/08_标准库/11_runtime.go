package main

import (
	"fmt"
	"runtime"
)

/*******************runtime的使用********************/
/*
runtime 调度器是个非常有用的东西，关于 runtime 包几个方法:
	Gosched()：让当前线程让出 cpu 以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行
	NumCPU：返回当前系统的 CPU 核数量
	GOMAXPROCS()：设置最大的可同时使用的 CPU 核数
	Goexit()：退出当前 goroutine(但是defer语句会照常执行)
		return结束当前函数,并返回指定值
		runtime.Goexit()结束当前goroutine,其他的goroutine不受影响,主程序也一样继续运行
		os.Exit 会结束当前程序
	NumGoroutine：返回正在执行和排队的任务总数
	GOOS：目标操作系统
 */



/*******************NumCPU********************/

func main111() {
	fmt.Println("cpus: ", runtime.NumCPU()) // cpus:  16
	fmt.Println("goroot: ", runtime.GOROOT()) // goroot:  D:\env\go
	fmt.Println("archive", runtime.GOOS) //  archive windows
}


/*******************GOMAXPROCS********************/
/*
	Golang 默认所有任务都运行在一个 cpu 核里，如果要在 goroutine 中使用多核，可以使用 runtime.GOMAXPROCS 函数修改，
当参数小于 1 时使用默认值。
 */
func init() {
	runtime.GOMAXPROCS(1)
}

/******************Gosched*********************/
/*
	让当前 goroutine 让出 CPU，当一个 goroutine 发生阻塞，Go 会自动地把与该 goroutine
处于同一系统线程的其他 goroutine 转移到另一个系统线程上去，以使这些 goroutine 不阻塞
 */
func main() {
	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("a:", i)

		if i == 1 {
			runtime.Gosched() //切换任务
		}
	}
	<-exit
}
