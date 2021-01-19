package main

import "fmt"

/*----------------切片Slice--------------------*/
/*
需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。
	1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
    2. 切片的长度可以改变，因此，切片是一个可变的数组。
    3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
    4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
    5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
    6. 如果 slice == nil，那么 len、cap 结果都等于 0。
 */


/*------------------创建切片的各种方式------------------*/
func main101() {
	//1. 声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("nil") //nil
	}else {
		fmt.Println("!= nil")
	}

	// 2.:=
	s2 := []int{}

	//3.make
	var s3 []int = make([]int, 0)

	var s4 []int
	s4 =  make([]int, 0)

	var s5 = make([]int, 0)
	fmt.Println(s1, s2, s3, s4, s5) //[] [] [] [] []

	// 4.初始化赋值
	var s6 []int = make([]int, 2, 2)
	s6[0] = 0
	s6[1] = 1
	//s6[2] = 2  //panic: runtime error: index out of range
	s6 = append(s6, 2)
	fmt.Println(s6) //[0 1 2]

	s7 := []int{1, 2, 3}
	fmt.Println(s7) //[1 2 3]

	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s8 []int
	// 前包后不包
	s8 = arr[1:4]
	fmt.Println(s8) //[2 3 4]
}




/*-----------------切片初始化-------------------*/
/*
全局：
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice0 []int = arr[start:end]
	var slice1 []int = arr[:end]
	var slice2 []int = arr[start:]
	var slice3 []int = arr[:]
	var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
局部：
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr[start:end]
	slice6 := arr[:end]
	slice7 := arr[start:]
	slice8 := arr[:]
	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素
 */

/*切片名称 [ low : high : max ]
low: 起始下标位置
high：结束下标位置	len = high - low
容量：cap = max - low
截取数组，初始化 切片时，没有指定切片容量时， 切片容量跟随原数组（切片）。
s[:high:max] : 从 0 开始，到 high结束。（不包含）
s[low:] :	从low 开始，到 末尾
s[: high]:	从 0 开始，到 high结束。容量跟随原先容量。【常用】*/


//全局
/*var arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 = arr[1:6] //[1 2 3 4 5]    包头不包尾
var slice1 = arr[0:5] //[0 1 2 3 4]
var slice2 = arr[:5] //[0 1 2 3 4]
var slice3 = arr[5:10]  //[5 6 7 8 9]
var slice4 = arr[5:] //[5 6 7 8 9]
var slice5 = arr[:] //[0 1 2 3 4 5 6 7 8 9]
var slice6 = arr[:len(arr)-1]  //[0 1 2 3 4 5 6 7 8]*/

/*func main102() {
	fmt.Printf("len=%d, cap=%d\n", len(slice0), cap(slice0)) //len=5, cap=9
	fmt.Printf("len=%d, cap=%d\n", len(slice1), cap(slice1)) //len=5, cap=10
	fmt.Printf("len=%d, cap=%d\n", len(slice2), cap(slice2)) //len=5, cap=10
	fmt.Printf("len=%d, cap=%d\n", len(slice3), cap(slice3)) //len=5, cap=5
	fmt.Printf("len=%d, cap=%d\n", len(slice4), cap(slice4)) //len=5, cap=5
	fmt.Printf("len=%d, cap=%d\n", len(slice5), cap(slice5)) //len=10, cap=10
	fmt.Printf("len=%d, cap=%d\n", len(slice6), cap(slice6)) //len=9, cap=10

	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr[2:8]
	slice6 := arr[0:6]         //可以简写为 slice := arr[:end]
	slice7 := arr[5:10]        //可以简写为 slice := arr[start:]
	slice8 := arr[0:len(arr)]  //slice := arr[:]
	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素
	fmt.Printf("局部变量： arr2 %v\n", arr2)
	fmt.Printf("局部变量： slice5 %v\n", slice5) //局部变量： slice5 [2 3 4 5 6 7]
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9)
}*/

/*-----------------通过make来创建切片-------------------*/
/*
	var slice []type = make([]type, len)
    slice  := make([]type, len)
    slice  := make([]type, len, cap)
 */
var slice0 []int = make([]int, 10)
var slice1 = make([]int, 10)
var slice2 = make([]int, 10, 10)
func main103()  {
	fmt.Printf("make全局slice0 ：%v\n", slice0)
	fmt.Printf("make全局slice1 ：%v\n", slice1)
	fmt.Printf("make全局slice2 ：%v\n", slice2)

	fmt.Println("--------------------------------------")

	slice3 := make([]int, 10)
	slice4 := make([]int, 10, 10)
	fmt.Printf("make局部slice3 ：%v\n", slice3)
	fmt.Printf("make局部slice4 ：%v\n", slice4)


	//make全局slice0 ：[0 0 0 0 0 0 0 0 0 0]
	//make全局slice1 ：[0 0 0 0 0 0 0 0 0 0]
	//make全局slice2 ：[0 0 0 0 0 0 0 0 0 0]
	//--------------------------------------
	//make局部slice3 ：[0 0 0 0 0 0 0 0 0 0]
	//make局部slice4 ：[0 0 0 0 0 0 0 0 0 0]
	//
	//Process finished with exit code 0
}

/*------------------读写操作实际目标是底层数组，只需注意索引号的差别------------------*/
func main104() {
	data1 := [...]int{0, 1, 2, 3, 4, 5}
	data2 := [...]int{0, 1, 2, 3, 4, 5}
	s1 := data1[2:4]
	s2 := data2[2:4]
	fmt.Println("s1 = ", s1)  //s1 =  [2 3]
	fmt.Println("s2 = ", s2)  //s2 =  [2 3]
	fmt.Printf("s1 - len = %d, cap = %d \n", len(s1), cap(s1))  //s1 - len = 2, cap = 4
	fmt.Printf("s2 - len = %d, cap = %d \n", len(s2), cap(s2))  //s2 - len = 2, cap = 4

	s1 = append(s1, 6)
	s1 = append(s1, 7)
	s1 = append(s1, 8)

	s2 = append(s2, 8)

	s1[0] += 98
	s1[1] += 97

	s2[0] += 98
	s2[1] += 97

	fmt.Println("s1 = ", s1) //s1 =  [100 100 6 7 8]
	fmt.Printf("s1 - len = %d, cap = %d \n", len(s1), cap(s1)) //s1 - len = 5, cap = 8

	fmt.Println("s2 = ", s2) //s2 =  [100 100 8]
	fmt.Printf("s2 - len = %d, cap = %d \n", len(s2), cap(s2)) //s2 - len = 3, cap = 4

	fmt.Println("data = ", data1) //data =  [0 1 2 3 6 7]
	fmt.Println("data = ", data2) //data =  [0 1 100 100 8 5]
}

func main105() {
	data1 := [...]int{0, 1, 2, 3, 4, 5}
	s1 := data1[2:4]

	fmt.Println("s1 = ", s1)  //s1 =  [2 3]

	fmt.Printf("s1 - len = %d, cap = %d \n", len(s1), cap(s1))  //s1 - len = 2, cap = 4

	//data1 p = 0xc04200a2d0, data1[0] p = 0xc04200a2d0, data1[2] p = 0xc04200a2e0
	fmt.Printf("data1 p = %p, data1[0] p = %p, data1[2] p = %p \n", &data1, &data1[0], &data1[2])
	//s1 p = 0xc0420023e0, s1[0] p = 0xc04200a2e0
	fmt.Printf("s1 p = %p, s1[0] p = %p \n", &s1, &s1[0])

	s1[0] += 98
	s1[1] += 97

	s1 = append(s1, 6)
	s1 = append(s1, 7)

	fmt.Printf("s1 - len = %d, cap = %d \n", len(s1), cap(s1)) //s1 - len = 4, cap = 4
	//s1 p = 0xc0420023e0, s1[0] p = 0xc04200a2e0
	fmt.Printf("s1 p = %p, s1[0] p = %p \n", &s1, &s1[0])

	s1 = append(s1, 8)

	//s1 p = 0xc0420023e0, s1[0] p = 0xc04200e2c0
	fmt.Printf("s1 p = %p, s1[0] p = %p \n", &s1, &s1[0]) //重新分配了底层数组，并复制数据 当s1.len > s1.cap时


	fmt.Println("s1 = ", s1) //s1 =  [100 100 6 7 8]
	fmt.Printf("s1 - len = %d, cap = %d \n", len(s1), cap(s1)) //s1 - len = 5, cap = 8


	fmt.Println("data = ", data1) //data =  [0 1 100 100 6 7]
	//data1 p = 0xc04200a2d0, data1[0] p = 0xc04200a2d0, data1[2] p = 0xc04200a2e0
	fmt.Printf("data1 p = %p, data1[0] p = %p, data1[2] p = %p \n", &data1, &data1[0], &data1[2])
}


/*----------------可直接创建 slice 对象，自动分配底层数组。--------------------*/
func main106() {
	s1 := []int{0, 1, 2, 3, 4, 5,  8: 100}  /// 通过初始化表达式构造，可使用索引号。
	fmt.Println("s1 = ", s1) //s1 =  [0 1 2 3 4 5 0 0 100]
	fmt.Printf("s1 - len = %d, cap = %d \n", len(s1), cap(s1)) //s1 - len = 9, cap = 9
	//s1[9] = 6 //panic: runtime error: index out of range  //读写操作实际目标是底层数组
	s1 = append(s1, 111)
	fmt.Printf("s1 - len = %d, cap = %d \n", len(s1), cap(s1)) //s1 - len = 10, cap = 18

	s2 := make([]int, 6, 10) // // 使用 make 创建，指定 len 和 cap 值。
	s2[0] = 0
	fmt.Printf("s2 - len = %d, cap = %d \n", len(s2), cap(s2)) //s2 - len = 6, cap = 10

	s3 := make([]int, 10) //// 省略 cap，相当于 cap = len。
	fmt.Printf("s3 - len = %d, cap = %d \n", len(s3), cap(s3)) //s3 - len = 10, cap = 10

}


/*------------------------------------*/
/*
使用 make 动态创建slice，避免了数组必须用常量做长度的麻烦。还可用指针直接访问底层数组，退化成普通数组操作。
 */
func main107() {
	s := []int{0, 1, 2}
	p := &s[0]
	*p = 10
	fmt.Println(s) //[10 1 2]
}


/*----------------[][]T，是指元素类型为 []T 。--------------------*/
func main108()  {
	ss := [][]int{
		[]int{1, 2},
		[]int{3, 4},
		[]int{5, 6},
	}
	fmt.Println(ss) //[[1 2] [3 4] [5 6]]
}


/*----------------可直接修改 struct array/slice 成员。--------------------*/
func main109() {
	d := [5]struct{
		x int
	}{}
	s := d[:]

	s[0].x = 1
	s[1].x = 1

	fmt.Println(s) //[{1} {1} {0} {0} {0}]
	fmt.Printf("%p, %p\n", &d, &d[0]) //0xc04200a2d0, 0xc04200a2d0
}

/*-----------------用append内置函数操作切片（切片追加）-------------------*/
func main1010() {
	a := []int{1, 2, 3}
	fmt.Println("a = ", a)//a =  [1 2 3]
	b := []int{4, 5, 6}
	fmt.Println("b = ", b) //b =  [4 5 6]
	c := append(a, b...)
	fmt.Println("c = ", c) //c =  [1 2 3 4 5 6]

	d := append(c, 7)
	fmt.Println("d = ", d)  //d =  [1 2 3 4 5 6 7]

	e := append(d, 8, 9)
	fmt.Println("e = ", e) //e =  [1 2 3 4 5 6 7 8 9]


	//append ：向 slice 尾部添加数据，返回新的 slice 对象。
	s1 := []int{1, 2, 3}
	fmt.Printf("s1-p=%p \n", &s1) //s1-p=0xc0420024a0
	s2 := append(s1, 4)
	fmt.Printf("s2-p=%p \n", &s2) //s2-p=0xc0420024c0
	fmt.Println("s1=", s1, "s2=", s2) //s1= [1 2 3] s2= [1 2 3 4]
}


/*
超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
 */
func main1011() {
	arr := [...]int{0, 1, 2, 3, 9:10}
	fmt.Println("arr=", arr)

	s := arr[:2:3]
	fmt.Println("s=", s) //s= [0 1]
	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s)) //len=2, cap=3

	s = append(s, 99) //此时未超出 s.cap 限制。会在原数组上修改
	fmt.Printf("&s[0]=%p, &arr[0]=%p \n", &s[0], &arr[0]) //&s[0]=0xc0420141e0, &arr[0]=0xc0420141e0

	s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制。  重新分配底层数组，与原数组无关。

	fmt.Println("s=", s, "arr=", arr) //[0 1 99 100 200] [0 1 99 3 0 0 0 0 0 10]

	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s)) //len=5, cap=6

	fmt.Printf("&s[0]=%p, &arr[0]=%p \n", &s[0], &arr[0]) //&s[0]=0xc04207a030, &arr[0]=0xc04208c0f0

	/*从输出结果可以看出，append 后的 s 重新分配了底层数组，并复制数据。
	如果只追加一个值，则不会超过 s.cap 限制，也就不会重新分配。
	通常以 2 倍容量重新分配底层数组。在大批量添加数据时，建议一次性分配足够大的空间，以减少内存分配和数据复制开销。
	或初始化足够长的 len 属性，改用索引号进行操作。及时释放不再使用的 slice 对象，避免持有过期数组，造成 GC 无法回收。*/

}

/*-----------------slice中cap重新分配规律：-------------------*/
func main1012() {
	 s := make([]int, 0, 1)
	 c := cap(s)
	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
		/*cap: 1 -> 2
		cap: 2 -> 4
		cap: 4 -> 8
		cap: 8 -> 16
		cap: 16 -> 32
		cap: 32 -> 64*/
}


/*-----------------切片拷贝-------------------*/
func main1013() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s1=%v\n", s1) //s1=[1 2 3 4 5]

	s2 := make([]int, 10)
	fmt.Printf("s2=%v\n", s2) //s2=[0 0 0 0 0 0 0 0 0 0]

	copy(s2, s1)
	fmt.Printf("s2=%v\n", s2) //s2=[1 2 3 4 5 0 0 0 0 0]

	s3 := []int{5, 4, 3, 2, 1}
	copy(s2, s3)
	fmt.Printf("s2=%v\n", s2) //s2=[5 4 3 2 1 0 0 0 0 0]
}

func main1014() {
	//copy ：函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠。
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array data : ", data)  //array data :  [0 1 2 3 4 5 6 7 8 9]
	s1 := data[8:]
	s2 := data[:5]
	fmt.Printf("slice s1 : %v\n", s1) ////slice s1 : [8 9]
	fmt.Printf("slice s2 : %v\n", s2) //slice s2 : [0 1 2 3 4]

	/*copy(s1, s2)
	fmt.Printf("slice s1 : %v\n", s1) //slice s1 : [0 1]
	fmt.Printf("slice s2 : %v\n", s2) //slice s2 : [0 1 2 3 4]*/

	copy(s2, s1)
	fmt.Printf("slice s1 : %v\n", s1) ////slice s1 : [8 9]
	fmt.Printf("slice s2 : %v\n", s2) //slice s2 : [8 9 2 3 4]

	fmt.Println("last array data : ", data) //last array data :  [8 9 2 3 4 5 6 7 8 9]
	
}

/*--------------slice遍历：----------------------*/
func main1016() {
	arr := [...]int{0, 1, 2, 3, 4, 5}
	s := arr[:]
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d]=%d \n", i, s[i])
	}
	fmt.Println("-------------------------------------")
	for idx, v := range s {
		fmt.Printf("idx=%d-v=%d \n", idx, v)
	}
}

/*------------------切片resize（调整大小）------------------*/
func main1017() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("a=%v, len=%d, cap=%d\n", a, len(a), cap(a)) //a=[1 2 3 4 5], len=5, cap=5
	b := a[1:2]
	fmt.Printf("b=%v, len=%d, cap=%d\n", b, len(b), cap(b)) //b=[2], len=1, cap=4
	c := b[0:4]
	fmt.Printf("c=%v, len=%d, cap=%d\n", c, len(c), cap(c)) //c=[2 3 4 5], len=4, cap=4
}


/*------------------字符串和切片（string and slice）------------------*/
/*
string底层就是一个byte的数组，因此，也可以进行切片操作。
 */
func main1018() {
	str := "hello world"
	s := str[:5]
	fmt.Println("s = ", s) //s =  hello

	s1 := s[:]
	fmt.Println("s1 = ", s1) //s1 =  hello
}



/*
string本身是不可变的，因此要改变string中字符。需要如下操作： 英文字符串：
 */
func main1019() {
	str := "hello world"
	byteArr := []byte(str) ////中文字符需要用[]rune(str)
	byteArr[0] = 'H'
	//byteArr[6] = "W"
	byteArr[6] = 'W'

	byteArr = append(byteArr, '!')
	//str= hello world byteArr= [72 101 108 108 111 32 87 111 114 108 100 33] string(byteArr)= Hello World!
	fmt.Println("str=", str, "byteArr=", byteArr, "string(byteArr)=", string(byteArr))

	byteArr[6] = 'G'
	byteArr = byteArr[:8]
	fmt.Println(string(byteArr)) //Hello Go
}



/*----------------含有中文字符串：--------------------*/
func main() {
	str := "你好，世界"
	s := []rune(str)
	s[3] = '够'
	s[4] = '浪'
	fmt.Println(string(s)) //你好，够浪
}



