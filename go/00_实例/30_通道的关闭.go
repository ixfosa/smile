package main

import "fmt"

// 关闭 一个通道意味着不能再向这个通道发送值了。
// 这个特性可以用来给这个通道的接收方传达工作已经完成的信息。
func main() {
	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		for true {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}

	close(jobs)
	fmt.Println("sent all jobs")
	<- done

}
