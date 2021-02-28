package main

import (
	"fmt"
	"sort"
)

// Go 的 sort 包实现了内置和用户自定义数据类型的排序功能。

func main() {
	// 排序方法是针对内置数据类型的；这里是一个字符串的例子。
	// 注意排序是原地更新的，所以他会改变给定的序列并且不返回一个新值。

	strs := []string{"a", "c", "b"}
	sort.Strings(strs)
	fmt.Println("strs = ", strs) // strs =  [a b c]

	ints := []int{1, 3, 2}
	sort.Ints(ints)
	fmt.Println("ints = ", ints) // ints =  [1 2 3]

	// 使用 sort 来检查一个序列是不是已经是排好序的。
	isOk := sort.IntsAreSorted(ints)
	fmt.Println("isOk = ", isOk) // isOk =  true
}
