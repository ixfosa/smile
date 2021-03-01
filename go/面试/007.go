package main

import "fmt"

// 1.下面这段代码能否通过编译？不能的话，原因是什么？如果通过，输出什么？
func main() {
	sn1 := struct {
		name string
		age int
	}{"Zi Qing", 18}

	sn2 := struct {
		name string
		age int
	}{"Zi Qing", 18}

	if sn1 == sn2 {
		 fmt.Println("sn1 == sn2") // sn1 == sn2
	}


	
	sn3 := struct {
		name string
		m map[int]int
	}{"Zi Qing", make( map[int]int)}

	sn4 := struct {
		name string
		m map[int]int
	}{"Zi Qing", make( map[int]int)}

	// if sn3 == sn4 {    // invalid operation: sn3 == sn4 (struct containing map[int]int cannot be compared)
	// 	fmt.Println(" sn3 == sn4") // sn1 == sn2
	// }
	fmt.Println(sn3, sn4)
}
/*
	参考答案及解析：编译不通过 invalid operation:  sn3 == sn4

	这道题目考的是结构体的比较，有几个需要注意的地方：
		结构体只能比较是否相等，但是不能比较大小。
		相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关，sn5 与 sn1 就是不同的结构体；
			 sn5:= struct {
					age  int
					name string
				}{age:11,name:"qq"}
	如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，比较时逐个项进行比较，
	如果每一项都相等，则两个结构体才相等，否则不相等；

	那什么是可比较的呢，常见的有 bool、数值型、字符、指针、数组等，像切片、map、函数等是不能比较的。
 */