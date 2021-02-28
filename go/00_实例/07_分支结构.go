package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 使用逗号来分隔多个表达式
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}


	// 不带表达式的 switch 是实现 if/else 逻辑的另一种方式
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon", t.Hour())
	default:
		fmt.Println("it's after noon", t.Hour())



	}


}
