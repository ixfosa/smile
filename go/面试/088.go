package main

import "fmt"

// 下面代码能编译通过吗？

type Info struct {
	result int
}

func work() (int, error) {
	return 88, nil
}
func main() {
	var data Info
	// data.result, err := work() // non-name data.result on left side of :=
	var err error
	data.result, err = work()

	if err != nil {
		return
	}
	fmt.Println(data)
}
