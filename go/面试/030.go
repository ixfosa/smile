package main

// 对 add() 函数调用正确的是（）
	// A. add(1, 2)
	// B. add(1, 3, 7)
	// C. add([]int{1, 2})
	// D. add([]int{1, 3, 7}…)

// 参考答案及解析：ABD。知识点：可变函数。


func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}


func main() {
	
}
