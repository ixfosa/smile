package main

import "fmt"

// f1()、f2()、f3() 函数分别返回什么？

func f1() (r int) {
	defer func() {
		r++
	}()
	return r
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r++
	}(r)
	return r
}



func main() {
	fmt.Println(f1())  // 1
	fmt.Println(f2())  // 5
	fmt.Println(f3())  // 0
}
