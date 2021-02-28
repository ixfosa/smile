package main

import "fmt"

// 返回目标字符串 t 出现的的第一个索引位置，或者在没有匹配值时返回 -1。
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 如果目标字符串 t 在这个切片中则返回 true
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 如果这些切片中的字符串有一个满足条件 f 则返回true
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 如果切片中的所有字符串都满足条件 f 则返回 true
// 如果这些切片中的字符串有一个满足条件 f 则返回true
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs{
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs{
		vsm[i] = f(v)
	}
	return vsm
}


func main() {
	persons:= []string{"fo", "long", "zhong", "qing"}

	fmt.Println("Index: ", Index(persons, "qing")) // Index:  3

	fmt.Println("Include: ", Include(persons, "long"))

	fmt.Println("Any: ", Any(persons, func(v string) bool { // Any:  true
		return len(v) == 2
	}))

	fmt.Println("All: ", All(persons, func(v string) bool { // All:  false
		return len(v) == 2
	}))

	fmt.Println("All: ", All(persons, func(v string) bool { // All:  true
		return len(v) < 6
	}))

	fmt.Println("Filter: ", Filter(persons, func(v string) bool { // Filter:  [long qing]
		return len(v) == 4
	}))

	fmt.Println("Map: ", Map(persons, func(v string) string {
		return "hello " + v
	})) // Map:  [hello fo hello long hello zhong hello qing]
}
