package main

import (
	"fmt"
)

/********************循环语句range*******************/
/*
Golang range类似迭代器操作，返回 (索引, 值) 或 (键, 值)。
for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：

	for key, value := range oldMap {
		newMap[key] = value
	}
            	1st value	2nd value
string	    	index		s[index]	unicode, rune
array/slice		index		s[index]
map				key			m[key]
channel			element
可忽略不想要的返回值，或 "_" 这个特殊变量。
 */
func main51() {
	s := "ixfosa"

	//// 忽略 2nd value，支持 string/array/slice/map。
	for i := range s {
		fmt.Print(string(s[i]))
	}
	fmt.Println()
	//// 忽略 index。
	for _, ch := range s {
		print("ch : ", ch, " ") //ch : 105 ch : 120 ch : 102 ch : 111 ch : 115 ch : 97
	}
	fmt.Println("----------------------")
	// 忽略全部返回值，仅迭代。
	for range s {

	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println("map: ", k, v)
	}
	//map:  a 1
	//map:  b 2
	//map:  c 3
}

/******************注意，range 会复制对象。*********************/
func main52() {
	//注意，range 会复制对象。
	a := [3]int{0, 1, 2}
	for i, v := range a {   //index、value 都是从复制品中取出。
		if i == 0 {  //在修改前，我们先修改原数组。
			a[1], a[2] = 999, 999
			fmt.Print(a, " ") //[0 999 999]
		}
		a[i] = v + 100  //使用复制品中取出的 value 修改原数组。
	}
	fmt.Print(a, " ") //[100 101 102]
}


/*****************建议改用引用类型，其底层数据不会被复制。**********************/
func main() {
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {  //复制 struct slice { pointer, len, cap }。
		if i == 0 {
			s = s[:3]  //对 slice 的修改，不会影响 range。
			s[2] = 100  //对底层数据的修改。
		}
		fmt.Println(i, v)
	}
	//0 1
	//1 2
	//2 100
	//3 4
	//4 5
	fmt.Println(s) //[1 2 100]

	//另外两种引用类型 map、channel 是指针包装，而不像 slice 是 struct。
}


/********************for 和 for range有什么区别?*******************/
/*
主要是使用场景不同
for可以  遍历array和slice
		遍历key为整型递增的map
		遍历string

for range可以完成所有for可以做的事情，却能做到for不能做的，包括
	遍历key为string类型的map并同时获取key和value
	遍历channel
 */