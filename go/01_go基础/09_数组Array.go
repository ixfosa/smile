package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*-----------------数组Array-------------------*/
/*
Golang Array和以往认知的数组有很大不同。
    1. 数组：是同一种数据类型的固定长度的序列。
    2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
    3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
    4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
    for i := 0; i < len(a); i++ {
    }
    for index, v := range a {
    }
    5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
    6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
    7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
    8.指针数组 [n]*T，数组指针 *[n]T。
 */


/*----------------一维数组--------------------*/
/*
 全局：
    var arr0 [5]int = [5]int{1, 2, 3}
    var arr1 = [5]int{1, 2, 3, 4, 5}
    var arr2 = [...]int{1, 2, 3, 4, 5, 6}
    var str = [5]string{3: "hello world", 4: "tom"}
    局部：
    a := [3]int{1, 2}           // 未初始化元素值为 0。
    b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
    c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素。
    d := [...]struct {
        name string
        age  uint8
    }{
        {"user1", 10}, // 可省略元素类型。
        {"user2", 20}, // 别忘了最后一行的逗号。
    }
 */
func main91() {
	const a int = 5
	//var b int = 5
	var arr0 [a]int = [5]int{1, 2, 3}

	//var arr1 [b]int = [5]int{1, 2, 3} //数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。

	fmt.Println("arr0 = ", arr0) //arr0 =  [1 2 3 0 0]  未初始化元素值为 0。

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr2 = ", arr2)

	arr3 := [...]int{1, 2, 3}  // 通过初始化值确定数组长度。
	fmt.Println("arr3 = ", arr3)

	arr4 := [5]string{1:"smile", 3:"ixfosa"}  // 使用索引号初始化元素。
	fmt.Println("arr4 = ", arr4)

	arr5 := [...]struct{
		name string
		age int
	}{
		{"ixfosa", 22}, // 可省略元素类型。
		{"大龙虾", 21},  //// 别忘了最后一行的逗号。
	}
	fmt.Println("arr5 = ", arr5)

	/*arr0 =  [1 2 3 0 0]
	arr2 =  [1 2 3 4 5]
	arr3 =  [1 2 3]
	arr4 =  [ smile  ixfosa ]
	arr5 =  [{ixfosa 22} {大龙虾 21}]*/

}

/*-----------------多维数组-------------------*/
/*
 全局
    var arr0 [5][3]int
    var arr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

局部：
    a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
    b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
 */

var arr0 [5][3]int
var arr1 = [5][3]int{{1, 2, 3}, {4, 5, 5}}
var arr3 = [...][3]int{{1, 2, 3}, {4, 5, 6}} //// 第 2 纬度不能用 "..."。 use of [...] array outside of array literal
func main92() {
	a := [3][2]int{{1, 1}, {2, 2}, {3, 3}}
	b := [...][3]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(arr1, arr3)
	fmt.Println(a, b)

	//[[1 2 3] [4 5 5] [0 0 0] [0 0 0] [0 0 0]] [[1 2 3] [4 5 6]]
	//[[1 1] [2 2] [3 3]] [[1 1 0] [2 2 0] [3 3 0]]
}

/*----------------数组遍历--------------------*/
func main93() {
	//一维数组遍历
	arr1 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("arr1[%d] = %v  ", i, arr1[i])
		//arr1[0] = 1  arr1[1] = 2  arr1[2] = 3  arr1[3] = 4  arr1[4] = 5
	}

	fmt.Println()

	for idx, v := range arr1 {
		fmt.Printf("arr1[%d] = %v  ", idx, v)
		//arr1[0] = 1  arr1[1] = 2  arr1[2] = 3  arr1[3] = 4  arr1[4] = 5
	}

	fmt.Println()

	//多维数组遍历
	arr2 := [2][2]string{{"hello", "world"}, {"ixfosa", "smile"}}
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++{
			fmt.Printf("arr2[%d][%d] = %v  ", i, j, arr2[i][j])
			//arr2[0][0] = hello  arr2[0][1] = world  arr2[1][0] = ixfosa  arr2[1][1] = smile
		}
	}

	fmt.Println()

	for idx1, v1 := range arr2 {
		for idx2, v2 := range v1 {
			fmt.Printf("arr2[%d][%d] = %v  ", idx1, idx2, v2)
			//arr2[0][0] = hello  arr2[0][1] = world  arr2[1][0] = ixfosa  arr2[1][1] = smile
		}
	}
}


/*----------------数组拷贝和传参--------------------*/
/*
	值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针
 */


/*------------------------------------*/
/*

 */
func test(arr [3]int)  {
	fmt.Printf("p = %p \n", &arr)
	arr[2] = 1
	fmt.Printf("test arr = %v\n", arr)
}

func test2(arr *[3]int)  {
	fmt.Printf("p = %p \n", arr) //p = 0xc042010320
	arr2 := *arr
	arr[2] = 1
	arr2[2] = 3
	(*arr)[0] = 9
	fmt.Printf("test2 arr = %v\n", arr) //test2 arr = &[9 1 1]
	fmt.Printf("test2 arr2 = %v\n", arr2) //test2 arr2 = [1 1 3]
}

func main94() {
	arr := [3]int{1, 1, 2}
	fmt.Printf("p = %p \n", &arr)

	test(arr)
	fmt.Printf("main arr = %v\n", arr)

	//p = 0xc042010320
	//p = 0xc042010360
	//test arr = [1 1 1]
	//main arr = [1 1 2]

	test2(&arr)
	fmt.Printf("main arr = %v\n", arr) //main arr = [9 1 1]
}

/*-----------------内置函数 len 和 cap 都返回数组长度 (元素数量)。-------------------*/
func main95() {
	arr := [3]int{1, 1, 2}
	fmt.Println("len = ", len(arr), "cap = ", cap(arr)) //len =  3 cap =  3
}


/*-----------------求数组所有元素之和-------------------*/
func sumArr(arr [10]int)  int {

	sum := 0

	for _, v := range arr {
		sum += v
	}
	return sum
}

func main96() {
	// 生成时间数种子
	// seed()种子默认是1
	//rand.Seed(1)
	rand.Seed(time.Now().UnixNano())

	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000) // 产生一个0到1000随机数
	}
	fmt.Println("arr = ", arr)
	fmt.Println("sumArr = ", sumArr(arr))
}

//找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，找出两个元素之和等于8的下标分别是（0，4）和（1，2）
func getIdxBySum(a [5]int, target int) {
	// 遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 继续遍历
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
func main() {
	arr := [5]int{1,3,5,8,7}
	getIdxBySum(arr, 8)
}

