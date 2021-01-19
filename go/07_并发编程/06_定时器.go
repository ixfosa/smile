package main

import (
	"fmt"
	"time"
)

/*
	Timer：时间到了，执行只执行1次
 */


// 1.timer基本使用
func main061() {
	timer := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <- timer.C
	fmt.Printf("t2:%v\n", t2)
}

// 2.验证timer只能响应1次
func main062()  {
	timer := time.NewTimer(time.Second)
	for {
		//<- timer.C
		fmt.Println("时间到:", <- timer.C) //时间到: 2020-11-02 18:45:07.0474534 +0800 CST m=+1.015806201
	}
}

//  3.timer实现延时的功能
func main063()  {
	//1
	time.Sleep(time.Second)

	//2
	timer := time.NewTimer(2 * time.Second)
	<- timer.C
	fmt.Println("2秒到")

	//3
	timer1 := time.NewTimer(2 * time.Second)
	<- timer1.C
	fmt.Println("2秒到")
}

//  4.停止定时器
func main064()  {
	timer := time.NewTimer(2 * time.Second)
	go func() {
		<- timer.C
		fmt.Println("定时器执行了")
	}()
	b := timer.Stop()

	if b {
		fmt.Println("timer已经关闭...")
	}
}

// 5.重置定时器
func main065()  {
	timer := time.NewTimer(5 * time.Second)

	timer.Reset(1 * time.Second)

	fmt.Println(time.Now())
	fmt.Println(<-timer.C)
}


/*
	Ticker：时间到了，多次执行
 */
func main()  {
	ticker := time.NewTicker(1 * time.Second)

	i := 0

	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)

			if i == 5 {
				ticker.Stop()
			}
			fmt.Println("ticker")
		}
	}()

	for {
		;
	}
}

