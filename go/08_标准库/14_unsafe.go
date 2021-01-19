package main

import (
	"fmt"
	"unsafe"
)

/****************unsafe 包***********************/
/*
	unsafe包的文档是说明的： 导入unsafe的软件包可能不可移植，并且不受Go 1兼容性指南的保护。

	func Alignof(x ArbitraryType) uintptr
	func Offsetof(x ArbitraryType) uintptr
	func Sizeof(x ArbitraryType) uintptr
	type ArbitraryType int      arbitrary  adj.	任意的; 武断的; 随心所欲的; 专横的; 专制的;
	type Pointer *ArbitraryType



	在unsafe包中，只提供了3个函数，两个类型。就这么少的量，却有着超级强悍的功能。
一般我们在C语言中通过指针，在知道变量在内存中占用的字节数情况下，就可以通过指针加偏移量的操作，直接在地址中，修改，访问变量的值。
在Go 语言中不支持指针运算，那怎么办呢？其实通过unsafe包，我们可以完成类似的操作。

	ArbitraryType 是以int为基础定义的一个新类型，但是Go 语言unsafe包中，对ArbitraryType赋予了特殊的意义，
通常，我们把interface{}看作是任意类型，那么ArbitraryType这个类型，在Go 语言系统中，比interface{}还要随意。

Pointer 是ArbitraryType指针类型为基础的新类型，在Go语言系统中，可以把Pointer类型，理解成任何指针的亲爹。

Go 语言的指针类型长度与int类型长度，在内存中占用的字节数是一样的。ArbitraryType类型的变量也可以是指针。

	func Alignof(x ArbitraryType) uintptr
	func Offsetof(x ArbitraryType) uintptr
	func Sizeof(x ArbitraryType) uintptr
这三个函数的参数均是ArbitraryType类型。
	1. Alignof返回变量对齐字节数量
	2. Offsetof返回变量指定属性的偏移量，所以如果变量是一个struct类型，不能直接将这个struct类型的变量当作参数，
		只能将这个struct类型变量的属性当作参数。
	3. Sizeof 返回变量在内存中占用的字节数，切记，如果是slice，则不会返回这个slice在内存中的实际占用长度。

	unsafe中，通过ArbitraryType 、Pointer 这两个类型，可以将其他类型都转换过来，然后通过这三个函数，
分别能取长度，偏移量，对齐字节数，就可以在内存地址映射中，来回游走。

 */


/******************指针运算*********************/
/*
	uintptr这个基础类型，在Go 语言中，字节长度是与int一致。
通常Pointer不能参与指针运算，比如你要在某个指针地址上加上一个偏移量，Pointer是不能做这个运算的，那么谁可以呢？
这里要靠 uintptr类型了，只有将Pointer类型先转换成uintptr类型，做完地址加减法运算后，再转换成Pointer类型，通过*操作达到取值、修改值的目的。

	unsafe.Pointer其实就是类似 C的void *，在Go 语言中是用于各种指针相互转换的桥梁，也即是通用指针。
它可以让任意类型的指针实现相互转换，也可以将任意类型的指针转换为 uintptr 进行指针运算。

uintptr是Go 语言的内置类型，是能存储指针的整型， uintptr 的底层类型是int，它和unsafe.Pointer可相互转换。

uintptr和unsafe.Pointer的区别就是：
	* unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
		* 任何类型的指针值都可以转换为unsafe.Pointer。
		* unsafe.Pointer可以转换为任何类型的指针值。
	* 而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
	* unsafe.Pointer 可以和 普通指针 进行相互转换；
	* unsafe.Pointer 可以和 uintptr 进行相互转换。

uintptr和intptr是无符号和有符号的指针类型，并且确保在64位平台上是8个字节，在32位平台上是4个字节，uintptr主要用于Go 语言中的指针运算。
 */

// 通过unsafe包来实现对V的成员i和j赋值，然后通过GetI()和GetJ()来打印观察输出结果。
type V struct {
	i int32
	j int64
}

func (v V) GetI() {
	fmt.Printf("i=%d\n", v.i)
}

func (v V) GetJ() {
	fmt.Printf("j=%d\n", v.j)
}

func main141() {
	// 定义指针类型变量
	var v *V = &V{6, 9}

	// 取得v的指针并转为*int32的值，对应结构体的i。
	var i *int32 = (*int32)(unsafe.Pointer(v))

	fmt.Println("指针地址：", i) // 指针地址： 0xc04205e080
	fmt.Println("指针uintptr值:", uintptr(unsafe.Pointer(i))) // 指针uintptr值: 825741402240
	*i = int32(98)

	// 根据v的基准地址加上偏移量进行指针运算，运算后的值为j的地址，使用unsafe.Pointer转为指针
	var j *int64 = (*int64)(unsafe.Pointer((uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int64(0))))))

	*j = 998

	v.GetI() // i=98
	v.GetJ() // j=998
}

/***************************************/
/*
	要修改struct字段的值，需要提前知道结构体V的成员布局，然后根据字段计算偏移量，以及考虑对齐值，最后通过指针运算得到成员指针，
利用指针达到修改成员值得目的。由于结构体的成员在内存中的分配是一段连续的内存，因此结构体中第一个成员的地址就是这个结构体的地址，
我们也可以认为是相对于这个结构体偏移了0。相同的，这个结构体中的任一成员都可以相对于这个结构体的偏移来计算出它在内存中的绝对地址。

	var v *V = &V{6, 9}
通过&来分配一段内存(并按类型初始化)，返回一个指针。所以v就是类型为V的一个指针。和new函数的作用类似。

	var i *int32 = (*int32)(unsafe.Pointer(v))
将指针v转成通用指针，再转成int32指针类型。这里就看到了unsafe.Pointer的作用了，不能直接将v转成int32类型的指针，
那样将会panic，但是unsafe.Pointer是可以转为任何指针。刚才说了v的地址其实就是它的第一个成员的地址，所以这个i就很显然指向了v的成员i，
通过给i赋值就相当于给v.i赋值了，但是别忘了i只是个指针，要赋值得解引用。

	*i = int32(98)
现在已经成功的改变了v的私有成员i的值。,但是对于v.j来说，怎么来得到它在内存中的地址呢？
其实我们可以获取它相对于v的偏移量(unsafe.Sizeof可以为我们做这个事)，但上面的代码并没有这样去实现。

	var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int64(0)))))
v是有两个成员的，包括i和j，并且在定义中，i位于j的前面，而i是int32类型，也就是说i占4个字节。
所以j是相对于v偏移了4个字节。可以用uintptr(4)或uintptr(unsafe.Sizeof(int64(0)))来做这个事。
unsafe.Sizeof方法用来得到一个值应该占用多少个字节空间。注意这里跟C的用法不一样，C是直接传入类型，而Go 语言是传入值。
之所以转成uintptr类型是因为需要做指针运算。v的地址加上j相对于v的偏移地址，也就得到了v.j在内存中的绝对地址，
然后通过unsafe.Pointer转为指针，别忘了j的类型是int64，所以现在的j就是一个指向v.j的指针，接下来给它赋值：

	*j = int64(763)

另外，我们可以看到两种地址表示上的差异：
	指针地址： 0xc04205e080
	指针uintptr值: 825741402240

 */

// 上面结构体V中，定义了2个成员属性，如果我们定义一个byte类型的成员属性。我们来看下它的输出：
type M struct {
	x byte
	y int32
	z int64
}

func (m M) GetX() {
	fmt.Printf("i=%d\n", m.x)
}
func (m M) GetY() {
	fmt.Printf("j=%d\n", m.y)
}
func (m M) GetZ() {
	fmt.Printf("j=%d\n", m.z)
}

func main()  {
	var m *M = new(M)

	// m的长度
	fmt.Printf("size=%d\n", unsafe.Sizeof(*m))

	//  Alignof 返回变量对齐字节数量
	fmt.Println("Alignof m:", unsafe.Alignof(m))
	fmt.Println("Alignof m.x:", unsafe.Alignof(m.x))
	fmt.Println("Alignof m.y:", unsafe.Alignof(m.y))
	fmt.Println("Alignof m.z:", unsafe.Alignof(m.z))

	//	Offsetof返回变量指定属性的偏移量，所以如果变量是一个struct类型，不能直接将这个struct类型的变量当作参数，
	//	只能将这个struct类型变量的属性当作参数。
	//fmt.Println("Offsetof m:", unsafe.Offsetof(m)) // invalid expression unsafe.Offsetof(m)
	fmt.Println("Offsetof m.x:", unsafe.Offsetof(m.x))
	fmt.Println("Offsetof m.y:", unsafe.Offsetof(m.y))
	fmt.Println("Offsetof m.z:", unsafe.Offsetof(m.z))

	fmt.Println("szieof byte: ", unsafe.Sizeof(byte(0)))      // szieof byte:  1
	fmt.Println("szieof int8: ", unsafe.Sizeof(int8(0)))      // szieof int8:  1
	fmt.Println("szieof int16: ", unsafe.Sizeof(int16(0)))    // szieof int26:  2
	fmt.Println("szieof int32: ", unsafe.Sizeof(int32(0)))    // szieof int32:  4
	fmt.Println("szieof int64: ", unsafe.Sizeof(int64(0)))    // szieof int64:  8
	fmt.Println("szieof string: ", unsafe.Sizeof(string(0)))  // szieof string:  16
	fmt.Println("szieof []int: ", unsafe.Sizeof([]int{0, 0})) // szieof []int:  24


	// 取得v的指针考虑对齐值计算偏移量，然后转为*int32的值，对应结构体的i。
	var y *int32 = (*int32)(unsafe.Pointer((uintptr(unsafe.Pointer(m)) + uintptr(unsafe.Offsetof(m.y)))))

	fmt.Println("指针地址 y：", y) // 指针地址 y： 0xc04205e084
	*y = 66
	m.GetY() // j=66

	// 根据v的基准地址加上偏移量进行指针运算，运算后的值为j的地址，使用unsafe.Pointer转为指针
 	var z *int64 = (*int64)(unsafe.Pointer((uintptr(unsafe.Pointer(m)) + uintptr(unsafe.Offsetof(m.z)))))
	fmt.Println("指针地址z ：", z) // 指针地址z ： 0xc04205e088
 	*z = 99
 	m.GetZ() // j=99

 	var x *byte = (*byte)(unsafe.Pointer(m))
	fmt.Println("指针地址 x：", x) // 指针地址 x： 0xc04205e080
 	*x = 22
 	m.GetX() // 22

	fmt.Printf("指针地址 m：%p\n", m) // 指针地址 m：0xc04205e080

	// 新结构体的长度为size=16，好像跟我们想像的不一致。
	// 我们计算一下：b是byte类型，占1个字节；i是int32类型，占4个字节；j是int64类型，占8个字节，1+4+8=13。这是怎么回事呢？
	//这是因为发生了对齐。在struct中，它的对齐值是它的成员中的最大对齐值。

	//每个成员类型都有它的对齐值，可以用unsafe.Alignof方法来计算，比如unsafe.Alignof(v.b)就可以得到b的对齐值为1 。
	//但这个对齐值是其值类型的长度或引用的地址长度（32位或者64位），和其在结构体中的size不是简单相加的问题。经过在64位机器上测试，

	fmt.Println("x", unsafe.Pointer(x))
	fmt.Println("y", unsafe.Pointer(y))
	fmt.Println("z", unsafe.Pointer(z))
	// x 0xc04205e080
	// y 0xc04205e084
	// z 0xc04205e088
}
