package main

import (
	"fmt"
	"time"
)

// 这是我们将要在多个并发实例中支持的任务了。这些执行者将从 jobs 通道接收任务，并且通过 results 发送对应的结果。
// 我们将让每个任务间隔 1s 来模仿一个耗时的任务。

func worker34(id int, jobs <- chan int, results chan <- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 1; i <= 3; i++ {
		go worker34(i, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		<-results
	}
}

// worker 3 processing job 1
// worker 1 processing job 2
// worker 2 processing job 3
// worker 1 processing job 4
// worker 3 processing job 5
// worker 2 processing job 6
// worker 3 processing job 7
// worker 2 processing job 9
// worker 1 processing job 8