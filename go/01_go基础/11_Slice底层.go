package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*----------------切片和数组--------------------*/
/*
在 Go 中，与 C 数组变量隐式作为指针使用不同，Go 数组是值类型，赋值和函数传参操作都会复制整个数组数据。
 */
func main110() {
	//三个内存地址都不同，这也就验证了 Go 中数组赋值和函数传参都是值复制的。
	arrA := [2]int{1, 2}
	var arrB [2]int
	arrB = arrA
	fmt.Printf("arrA:%p, %v\n", &arrA, arrA) //arrA:0xc0420080c0, [1 2]
	fmt.Printf("arrB:%p, %v\n", &arrB, arrB) //arrB:0xc0420080d0, [1 2]
	testArr(arrA)                                   //arr :0xc042008110, [1 2]
}

func testArr(arr [2]int)  {
	fmt.Printf("arr:%p, %v", &arr, arr)
}

/*-----------------函数传参用数组的指针-------------------*/
func main111() {
	arr1 := [2]int{1, 2}
	testArrPoint1(&arr1) //testArrPoint1 arr:0xc04205e080, [1 2]
	fmt.Printf("arr1:%p, %v \n", &arr1, arr1)  //arr1:0xc04205e080, [1 3]
	arr2 := arr1[:]
	testArrPoint2(arr2) //testArrPoint2 arr:0xc0420583e0, [1 3]
	fmt.Printf("arr2:%p, %v \n", &arr2, arr2) //arr2:0xc0420583c0, [1 4]
}
func testArrPoint1(x *[2]int)  {
	fmt.Printf("testArrPoint1 arr:%p, %v \n", x, *x)
	(*x)[1] += 1
}
func testArrPoint2(x []int)  {
	fmt.Printf("testArrPoint2 arr:%p, %v \n", &x, x)
	x[1] += 1
}

func main112() {
	/*data := [5]int{1, 2, 3, 4, 5}
	s := data[:2:3]
	s = append(s, 9, 8)
	fmt.Printf("data:%v\n", data) //data:[1 2 3 4 5]
	fmt.Printf("s:%v \n", s) //s:[1 2 9 8] */
	data := [5]int{1, 2, 3, 4, 5}
	s := data[:2:3]
	s = append(s, 9)
	s = append(s, 8)
	fmt.Printf("data:%v\n", data) //data:[1 2 9 4 5]
	fmt.Printf("s:%v \n", s) //s:[1 2 9 8]
}



/*----------------切片的数据结构--------------------*/
/*
切片本身并不是动态数组或者数组指针。
它内部实现的数据结构通过指针引用底层数组，设定相关属性将数据读写操作限定在指定的区域内。
切片本身是一个只读对象，其工作机制类似数组指针的一种封装。

切片（slice）是对数组一个连续片段的引用，所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型）。
这个片段可以是整个数组，或者是由起始和终止索引标识的一些项的子集。需要注意的是，终止索引标识的项不包括在切片内。
切片提供了一个与指向数组的动态窗口。

给定项的切片索引可能比相关数组的相同元素的索引小。
和数组不同的是，切片的长度可以在运行时修改，最小为 0 最大为相关数组的长度：切片是一个长度可变的数组。

Slice 的数据结构定义如下:
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}
切片的结构体由3部分构成，unsafe.Pointer 是指向一个数组的指针，len 代表当前切片的长度，cap 是当前切片的容量。cap 总是大于等于 len 的。
 */
func main113() {
	//u如果想从 slice 中得到一块内存地址，可以这样做：
	/*s := make([]int, 100)
	ptr := unsafe.Pointer(&s[0])
	fmt.Printf("ptr: %p", ptr)*/

	//如果反过来呢？从 Go 的内存地址中构造一个 slice。
	var ptr unsafe.Pointer
	length := 1
	s := struct {
		addr unsafe.Pointer
		len int
		cap int
	}{ptr, length, length}

	s1 := *(*[]byte)(unsafe.Pointer(&s))

	fmt.Printf("s1:%p, %v\n", &s1, s) //s1:0xc0420023e0, {<nil> 1 1}

	//在 Go 的反射中就存在一个与之对应的数据结构 SliceHeader，我们可以用它来构造一个 slice
	var o []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&o))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(ptr)
	fmt.Printf("sliceHeader:%p, %v", &sliceHeader, sliceHeader) //sliceHeader:0xc042004030, &{0 1 1}
}
func main114() {
	//u如果想从 slice 中得到一块内存地址，可以这样做：
	s := make([]int, 5)
	ptr := unsafe.Pointer(&s[0])
	intPtr := (*int)(ptr)
	strPtr := (*string)(ptr)
	fmt.Printf("ptr: %p\n", ptr) //ptr: 0xc04200a2d0
	fmt.Printf("intPtr: %p \n", intPtr) //intPtr: 0xc04200a2d0
	fmt.Printf("strPtr: %p \n", strPtr) //strPtr: 0xc04200a2d0

	*intPtr = 1
	fmt.Printf("intPtr: %p, %v\n", intPtr, *intPtr) //intPtr: 0xc04200a2d0, 1
	fmt.Printf("s[0]: %p, %v\n", &s[0], s) //s[0]: 0xc04200a2d0, [1 0 0 0 0]

	*strPtr = "ixfosa"

	fmt.Printf("strPtr: %p, %v\n", strPtr, string(*strPtr)) //strPtr: 0xc04200a2d0, ixfosa
	fmt.Printf("s[0]: %p, %v\n", &s[0], s) //s[0]: 0xc04200a2d0, [5015717 6 0 0 0]
	fmt.Println(s) //[5015717 6 0 0 0]

}
/*-----------------创建切片-------------------*/
/*
make 函数允许在运行期动态指定数组长度，绕开了数组类型必须使用编译期常量的限制。
创建切片有两种形式，make 创建切片，空切片。

make 和切片字面量
	func makeslice(et *_type, len, cap int) slice {
		// 根据切片的数据类型，获取切片的最大容量
		maxElements := maxSliceCap(et.size)
		// 比较切片的长度，长度值域应该在[0,maxElements]之间
		if len < 0 || uintptr(len) > maxElements {
			panic(errorString("makeslice: len out of range"))
		}
		// 比较切片的容量，容量值域应该在[len,maxElements]之间
		if cap < len || uintptr(cap) > maxElements {
			panic(errorString("makeslice: cap out of range"))
		}
		// 根据切片的容量申请内存
		p := mallocgc(et.size*uintptr(cap), et, true)
		// 返回申请好内存的切片的首地址
		return slice{p, len, cap}
	}
 */



/*----------------nil 和空切片--------------------*/
/*
	var slice []int
nil 切片被用在很多标准库和内置函数中，描述一个不存在的切片的时候，就需要用到 nil 切片。
比如函数在发生异常的时候，返回的切片就是 nil 切片。nil 切片的指针指向 nil。

空切片一般会用来表示一个空的集合。比如数据库查询，一条结果也没有查到，那么就可以返回一个空切片。
 */
func main116() {
	//返回一个空切片。
	s := make([]int, 0)
	s1 := []int{}
	fmt.Println("s = ", s ,"s1 = ", s1) //s =  [] s1 =  []

	//空切片和 nil 切片的区别在于，空切片指向的地址不是nil，指向的是一个内存地址，但是它没有分配任何内存空间，即底层元素包含0个元素。
	//最后需要说明的一点是。不管是使用 nil 切片还是空切片，对其调用内置函数 append，len 和 cap 的效果都是一样的。

	var s2 []int
	fmt.Println("s2 = ", s2) //s2 =  []
}



/*-----------------切片扩容-------------------*/
/*
当一个切片的容量满了，就需要扩容了。怎么扩，策略是什么？
	主要需要关注的有两点，一个是扩容时候的策略，还有一个就是扩容是生成全新的内存地址还是在原来的地址后追加。
	func growslice(et *_type, old slice, cap int) slice {
		if raceenabled {
			callerpc := getcallerpc(unsafe.Pointer(&et))
			racereadrangepc(old.array, uintptr(old.len*int(et.size)), callerpc, funcPC(growslice))
		}
		if msanenabled {
			msanread(old.array, uintptr(old.len*int(et.size)))
		}

		if et.size == 0 {
			// 如果新要扩容的容量比原来的容量还要小，这代表要缩容了，那么可以直接报panic了。
			if cap < old.cap {
				panic(errorString("growslice: cap out of range"))
			}

			// 如果当前切片的大小为0，还调用了扩容方法，那么就新生成一个新的容量的切片返回。
			return slice{unsafe.Pointer(&zerobase), old.len, cap}
		}

	  // 这里就是扩容的策略
		newcap := old.cap
		doublecap := newcap + newcap
		if cap > doublecap {
			newcap = cap
		} else {
			if old.len < 1024 {
				newcap = doublecap
			} else {
				for newcap < cap {
					newcap += newcap / 4
				}
			}
		}

		// 计算新的切片的容量，长度。
		var lenmem, newlenmem, capmem uintptr
		const ptrSize = unsafe.Sizeof((*byte)(nil))
		switch et.size {
		case 1:
			lenmem = uintptr(old.len)
			newlenmem = uintptr(cap)
			capmem = roundupsize(uintptr(newcap))
			newcap = int(capmem)
		case ptrSize:
			lenmem = uintptr(old.len) * ptrSize
			newlenmem = uintptr(cap) * ptrSize
			capmem = roundupsize(uintptr(newcap) * ptrSize)
			newcap = int(capmem / ptrSize)
		default:
			lenmem = uintptr(old.len) * et.size
			newlenmem = uintptr(cap) * et.size
			capmem = roundupsize(uintptr(newcap) * et.size)
			newcap = int(capmem / et.size)
		}

		// 判断非法的值，保证容量是在增加，并且容量不超过最大容量
		if cap < old.cap || uintptr(newcap) > maxSliceCap(et.size) {
			panic(errorString("growslice: cap out of range"))
		}

		var p unsafe.Pointer
		if et.kind&kindNoPointers != 0 {
			// 在老的切片后面继续扩充容量
			p = mallocgc(capmem, nil, false)
			// 将 lenmem 这个多个 bytes 从 old.array地址 拷贝到 p 的地址处
			memmove(p, old.array, lenmem)
			// 先将 P 地址加上新的容量得到新切片容量的地址，然后将新切片容量地址后面的 capmem-newlenmem 个 bytes 这块内存初始化。为之后继续 append() 操作腾出空间。
			memclrNoHeapPointers(add(p, newlenmem), capmem-newlenmem)
		} else {
			// 重新申请新的数组给新切片
			// 重新申请 capmen 这个大的内存地址，并且初始化为0值
			p = mallocgc(capmem, et, true)
			if !writeBarrier.enabled {
				// 如果还不能打开写锁，那么只能把 lenmem 大小的 bytes 字节从 old.array 拷贝到 p 的地址处
				memmove(p, old.array, lenmem)
			} else {
				// 循环拷贝老的切片的值
				for i := uintptr(0); i < lenmem; i += et.size {
					typedmemmove(et, add(p, i), add(old.array, i))
				}
			}
		}
		// 返回最终新切片，容量更新为最新扩容之后的容量
		return slice{p, old.len, newcap}
	}
 */
func main117() {
	slice := []int{10, 20, 30, 40}
	newSlice := append(slice, 50)
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
}
/*Go 中切片扩容的策略是这样的：
如果切片的容量小于 1024 个元素，于是扩容的时候就翻倍增加容量。上面那个例子也验证了这一情况，总容量从原来的4个翻倍到现在的8个。
一旦元素个数超过 1024 个元素，那么增长因子就变成 1.25 ，即每次增加原来容量的四分之一。
注意：扩容扩大的容量都是针对原来的容量而言的，而不是针对原来数组的长度而言的。*/

/*------------------切片拷贝------------------*/
/*

 */
func main118() {
	/*array := []int{10, 20, 30, 40}
	slice := make([]int, 6)
	n := copy(slice, array)
	fmt.Println(n,slice)*/

	slice := make([]byte, 3)
	n := copy(slice, "abcdef")
	fmt.Println(n,slice) //  3 [97,98,99]
}

//如果用 range 的方式去遍历一个切片，拿到的 Value 其实是切片里面的值拷贝。所以每次打印 Value 的地址都不变。
//由于 Value 是值拷贝的，并非引用传递，所以直接改 Value 是达不到更改原切片值的目的的，需要通过 &slice[index] 获取真实的地址。
func main() {
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("value = %d , value-addr = %x , slice-addr = %x\n", value, &value, &slice[index])
	}
}
/*
value = 10 , value-addr = c4200aedf8 , slice-addr = c4200b0320
value = 20 , value-addr = c4200aedf8 , slice-addr = c4200b0328
value = 30 , value-addr = c4200aedf8 , slice-addr = c4200b0330
value = 40 , value-addr = c4200aedf8 , slice-addr = c4200b0338
*/
