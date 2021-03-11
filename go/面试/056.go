package main

// 下面的代码有什么问题？
func main() {
	// go 中不同类型是不能比较的，而数组长度是数组类型的一部分，所以 […]int{1} 和 [2]int{1} 是两种不同的类型，不能比较；
	// invalid operation: [1]int literal == [2]int literal (mismatched types [1]int and [2]int)
	// fmt.Println([...]int{1} == [2]int{1})

	// invalid operation: []int literal == []int literal (slice can only be compared to nil)
	// fmt.Println([]int{1} == []int{1}) // 切片是不能比较的；
}
