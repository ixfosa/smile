package main

import "fmt"

func main() {
	// slice 的类型仅由它所包含的元素决定（不像数组中还需要元素的个数）。
	// 要创建一个长度非零的空slice，需要使用内建的方法 make。
	a := make([]string, 3)
	fmt.Println("a = ", a) // a =  [ ]

	// 和数组一样设置和得到值
	a[0] = "a"
	a[1] = "b"
	a[2] = "c"
	fmt.Println("a[1] = ", a[1]) // a[1] =  b
	fmt.Println("a = ", a) // a =  [a b c]

	// len 返回 slice 的长度
	fmt.Println("len = ", len(a)) // len =  3

	// 其中一个是内建的 append，它返回一个包含了一个或者多个新值的 slice。
	// 注意我们接受返回由 append返回的新的 slice 值。
	// a[3] = "f" // panic: runtime error: index out of range
	a = append(a, "d")
	fmt.Println("a = ", a) // a =  [a b c d]
	fmt.Println("len = ", len(a)) // len =  4
	fmt.Println("cap = ", cap(a)) // cap =  6
	// a[5] = "e" // panic: runtime error: index out of range
	a = append(a, "e", "f")
	fmt.Println("a = ", a, "len = ", len(a), "cap = ", cap(a)) // a =  [a b c d e f] len =  6 cap =  6

	b := make([]string, len(a))
	copy(b, a)
	fmt.Println("b = ", b, "len = ", len(b), "cap = ", cap(b)) // b =  [a b c d e f] len =  6 cap =  6

	// Slice 支持通过 slice[low:high] 语法进行“切片”操作。
	l := a[2:5]
	fmt.Println("a[2:5]:", l, "len = ", len(l), "cap = ", cap(l)) // a[2:5]: [c d e] len =  3 cap =  4

	l[0] = "cc"
	fmt.Println("l = ", l, "a = ", a) // l =  [cc d e] a =  [a b cc d e f]

	l = append(l, "g")
	fmt.Println("l:", l, "len = ", len(l), "cap = ", cap(l)) // l: [cc d e g] len =  4 cap =  4
	fmt.Println("a = ", a) // a =  [a b cc d e g]

	l = append(l, "h")
	fmt.Println("l:", l, "len = ", len(l), "cap = ", cap(l)) // l: [cc d e g] len =  4 cap =  4
	fmt.Println("a = ", a) // a =  [a b cc d e g]

	l[0] = "ccccc"
	fmt.Println("l = ", l, "a = ", a) // l =  [ccccc d e g h] a =  [a b cc d e g]


	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // 2d:  [[0] [1 2] [2 3 4]]
}
