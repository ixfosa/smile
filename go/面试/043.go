package main

import "fmt"

// 下面 A、B 两处应该填入什么代码，才能确保顺利打印出结果？
type S043 struct {
	m string
}

func f043() *S043 {
	// return ___  // A
	return &S043{
		"foo",
	}
}

func main() {
	// p := ___   // B
	p := f043()   // *f() 或者 f()
	fmt.Println(p.m) //print "foo"
}
