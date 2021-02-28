package main

import (
	"fmt"
	"sort"
)

// 在类型中实现了 sort.Interface 的 Len，Less和 Swap 方法，这样我们就可以使用 sort 包的通用Sort 方法
/*
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
 */

// 创建一个为内置 []string 类型的别名的ByLength 类型
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

// 按字符串长度增加的顺序来排序，所以这里使用了 len(s[i]) 和 len(s[j])。
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type user struct {
	name string
	age int
}

type ByUserAge []user

func (s ByUserAge) Len() int {
	return len(s)
}

func (s ByUserAge) Less(i, j int) bool {
	return s[i].age < s[j].age
}

func (s ByUserAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)

	users := []user{
		user{"fo", 22},
		user{"qing", 18},
		user{"zhong", 21},
	}
	sort.Sort(ByUserAge(users))
	fmt.Println(users) // [{qing 18} {zhong 21} {fo 22}]
}
