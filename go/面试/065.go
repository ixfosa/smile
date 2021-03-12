package main

import "fmt"

// 下面代码里的 counter 的输出值？
	// A. 2
	// B. 3
	// C. 2 或 3

// C。for range map 是无序的，如果第一次循环到 A，则输出 3；否则输出 2。


func main() {
	m := map[string]int{
		"A": 21,
		"B": 22,
		"c": 23,
	}

	counter := 0

	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}

	fmt.Println("counter: ", counter)
}
