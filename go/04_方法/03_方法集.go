package main

import "fmt"

/*****************方法集**********************/
/*
Golang方法集 ：每个类型都有与之关联的方法集，这会影响到接口实现规则。
    • 类型 T 方法集包含全部 receiver T 方法。
    • 类型 *T 方法集包含全部 receiver T + *T 方法。
    • 如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
    • 如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。
    • 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。
用实例 value 和 pointer 调用方法 (含匿名字段) 不受方法集约束，编译器总是查找全部方法，并自动转换 receiver 实参。
 */


type T struct {
	int
}

//类型 T 方法集包含全部 receiver T 方法。
func (t T) test() {
	fmt.Println("类型 T 方法集包含全部 receiver T 方法。")
}

//类型 *T 方法集包含全部 receiver T + *T 方法。
func (t T) testT() {
	fmt.Println("类型 *T 方法集包含全部 receiver T 方法。")
}

func (t *T) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
}

func main31() {
	t1 := T{1}
	fmt.Printf("t1 is : %v\n", t1)
	t1.test()
	//t1 is : {1}
	//类型 T 方法集包含全部 receiver T 方法。

	t2 := T{1}
	t3 := &t2
	fmt.Printf("t2 is : %v\n", t2)
	t3.testT()
	t3.testP()
	//t2 is : {1}
	//类型 *T 方法集包含全部 receiver T 方法。
	//类型 *T 方法集包含全部 receiver *T 方法。
}

/***************************************/
/*
给定一个结构体类型 S 和一个命名为 T 的类型，方法提升像下面规定的这样被包含在结构体方法集中：
如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
这条规则说的是当我们嵌入一个类型，嵌入类型的接受者为值类型的方法将被提升，可以被外部类型的值和指针调用
 */
type S struct {
	T
}

func (t T) testST() {
	fmt.Println("如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。")
}

func main32() {
	s1 := S{T{1}}
	s2 := &s1
	fmt.Printf("s1 is : %v\n", s1)
	s1.testST()
	fmt.Printf("s2 is : %v\n", s2)
	s2.testST()
	//s1 is : {{1}}
	//如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
	//s2 is : &{{1}}
	//如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
}

/***************************************/
/*
如类型 SS 包含匿名字段 *T，则 SS 和 *SS 方法集包含 T + *T 方法。

这条规则说的是当我们嵌入一个类型的指针，嵌入类型的接受者为值类型或指针类型的方法将被提升，可以被外部类型的值或者指针调用。
 */
type SS struct {
	*T
}

func (t T) testSST() {
	fmt.Println("如类型 SS 包含匿名字段 *T，则 SS 和 *SS 方法集包含 T 方法")
}

func (t *T) testSSP() {
	fmt.Println("如类型 SS 包含匿名字段 *T，则 SS 和 *SS 方法集包含 *T 方法")
}
func main() {
	ss1 := SS{&T{1}}
	ss2 := &ss1

	fmt.Printf("ss1 is : %v\n", ss1)
	ss1.testSST()
	ss1.testSSP()

	fmt.Printf("ss2 is : %v\n", ss2)
	ss2.testSST()
	ss2.testSSP()

	//ss1 is : {0xc0420080a8}
	//如类型 SS 包含匿名字段 *T，则 SS 和 *SS 方法集包含 T 方法
	//如类型 SS 包含匿名字段 *T，则 SS 和 *SS 方法集包含 *T 方法
	//ss2 is : &{0xc0420080a8}
	//如类型 SS 包含匿名字段 *T，则 SS 和 *SS 方法集包含 T 方法
	//如类型 SS 包含匿名字段 *T，则 SS 和 *SS 方法集包含 *T 方法
}
/***************************************/

/*

 */