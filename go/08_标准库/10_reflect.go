package main

import (
	"fmt"
	"reflect"
)

/******************Golang反射（reflect）*********************/
/*
	反射是指在 程序运行期对程序本身进行访问和修改的能力，程序在编译时变量被转换为内存地址，变量名不会被编译器写入到可执行部分，
在运行程序时程序无法获取自身的信息。

	支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，
这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。

Go语言提供了 `reflect` 包来访问程序的反射信息。

reflect 包
	Go语言中的反射是由 reflect 包提供支持的，它定义了两个重要的类型 `Type `和 `Value`
任意接口值在反射中都可以理解为由` reflect.Type `和 `reflect.Value` 两部分组成，
并且 reflect 包提供了` reflect.TypeOf` 和 `reflect.ValueOf `两个函数来获取任意对象的 Value 和 Type。
	func TypeOf(i interface{}) Type
	func ValueOf(i interface{}) Value

变量的内在机制
	变量包含类型信息和值信息 var arr [10]int arr[0] = 10
	类型信息：是静态的元信息，是预先定义好的
	值信息：是程序运行过程中动态改变的

reflect包封装了反射相关的方法
	获取类型信息：reflect.TypeOf，是静态的
	获取值信息：reflect.ValueOf，是动态的
 */


/******************Type类型与Value类型*********************/
/*
`类型`（Type）和 `种类`（Kind）的区别。
	编程中，使用最多的是类型，但在反射中，当需要区分一个大品种的类型时，就会用到种类（Kind）。
	例如需要统一判断类型中的指针时，使用种类（Kind）信息就较为方便。

	反射包中类型 `Type` 是 Golang 反射包中定义的一个接口，我们可以使用 `TypeOf` 函数获取任意值的变量的的类型，
`MethodByName` 可以获取当前类型对应方法的引用、`Implements` 可以判断当前类型是否实现了某个接口：
	type Type interface {
			Align() int
			FieldAlign() int
			Method(int) Method
			MethodByName(string) (Method, bool)
			NumMethod() int
			Name() string
			PkgPath() string
			Size() uintptr
			String() string
			Kind() Kind
			Implements(u Type) bool
			...
	}

	反射包中 `Value` 的类型却与 `Type` 不同，`Type` 是一个接口类型，
但是 `Value` 在 reflect`包中的定义是一个结构体，这个结构体没有任何对外暴露的成员变量，
但是却提供了很多方法让我们获取或者写入 `Value` 结构体中存储的数据：
	type Value struct {
			// contains filtered or unexported fields
	}

	func (v Value) Addr() Value
	func (v Value) Bool() bool
	func (v Value) Bytes() []byte
	func (v Value) Float() float64
	...

	我们通过 `TypeOf`、`ValueOf `方法就可以将一个普通的变量转换成『反射』包中提供的 Type 和 Value，
使用反射提供的方法对这些类型进行复杂的操作。
 */

// 使用` reflect.TypeOf() `函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息
// 使用` reflect.ValueOf() `函数可以获得任意值的类型对象（reflect.Value）
func main101() {
	var a int
	// 通过 reflect.TypeOf() 取得变量 a 的类型对象 typeOfA，类型为 reflect.Type()。
	typeOfA := reflect.TypeOf(a)
	// 通过 typeOfA 类型对象的成员函数，可以分别获取到 typeOfA 变量的类型名为 int，种类（Kind）为 int。
	fmt.Println(typeOfA.Name(), typeOfA.Kind()) // int int


	valueOfA := reflect.ValueOf(a)
	fmt.Println(valueOfA.Kind()) // int

}


/*
类型 reflect.Value 有一个方法` Type()`，它会返回一个 reflect.Type 类型的对象。

Type 和 Value 都有一个名为` Kind `的方法，它会返回一个常量，表示底层数据的类型，常见值有：Uint、Float64、Slice 等。

Value 类型也有一些类似于 Int、Float 的方法，用来提取底层的数据：
	- Int 方法用来提取 int64
	- Float 方法用来提取 float64，示例代码如下：

反射对象的 Kind 方法描述的是基础类型，而不是静态类型。如果一个反射对象包含了用户定义类型的值
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
虽然变量 v 的静态类型是 MyInt，而不是 int，但 Kind 方法仍然会返回 reflect.Int。
换句话说 Kind 方法不会像 Type 方法一样区分 MyInt 和 int。
*/
func main105()  {

	var x float64 = 3.14
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type()) // type: float64
	// Kind()的方法，它会返回一个常量，表示底层数据的类型
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64) // kind is float64: true
	// Float() 方法用来提取 float64
	fmt.Println("value:", v.Float()) // value: 3.14

	var y uint8 = '6'
	v2 := reflect.ValueOf(y)
	fmt.Println("type:", v2.Type()) //type: uint8
	fmt.Println("kind is uint8: ", v2.Kind() == reflect.Uint8) // kind is uint8:  true
	//y = v2.Uint()  // cannot use v2.Uint() (type uint64) as type uint8 in assignment
	y = uint8(v2.Uint())
	fmt.Println("y: ", string(y)) // y:  6

	type myInt int
	var a myInt = 2
	va := reflect.ValueOf(a)
	fmt.Println("kind:", va.Kind()) //kind: int
	fmt.Println("type:", va.Type()) //type: main.myInt

}


/*******************反射种类（Kind）的定义********************/
/*
	Go语言程序中的类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type 关键字定义的类型，
这些类型的名称就是其类型本身的名称。例如使用 type A struct{} 定义结构体时，A 就是 struct{} 的类型。

种类（Kind）指的是对象归属的品种，在 reflect 包中有如下定义：
	type Kind uint
	const (
		Invalid Kind = iota  // 非法类型
		Bool                 // 布尔型
		Int                  // 有符号整型
		Int8                 // 有符号8位整型
		Int16                // 有符号16位整型
		Int32                // 有符号32位整型
		Int64                // 有符号64位整型
		Uint                 // 无符号整型
		Uint8                // 无符号8位整型
		Uint16               // 无符号16位整型
		Uint32               // 无符号32位整型
		Uint64               // 无符号64位整型
		Uintptr              // 指针
		Float32              // 单精度浮点数
		Float64              // 双精度浮点数
		Complex64            // 64位复数类型
		Complex128           // 128位复数类型
		Array                // 数组
		Chan                 // 通道
		Func                 // 函数
		Interface            // 接口
		Map                  // 映射
		Ptr                  // 指针
		Slice                // 切片
		String               // 字符串
		Struct               // 结构体
		UnsafePointer        // 底层指针
	)

	Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。
type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr。

 */


/******************类型对象中获取类型名称和种类*********************/
/*
	Go语言中的类型名称对应的反射获取方法是` reflect.Type `中的 `Name()` 方法，返回表示`类型名称的字符串`；
类型归属的`种类`（Kind）使用的是 reflect.Type 中的` Kind() `方法，返回 reflect.Kind 类型的常量。
 */

// 定义一个Enum类型
type Enum int

const zero Enum = 0

func main102()  {

	// 声明一个空结构体
	type cat struct {}
	// 将 cat 实例化，并且使用 reflect.TypeOf() 获取被实例化后的 cat 的反射类型对象。
	typeOfCat := reflect.TypeOf(cat{})

	//输出 cat 的类型名称和种类，类型名称就是 cat，而 cat 属于一种结构体种类，因此种类为 struct。
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind()) // cat struct

	// 获取Zero常量的反射类型对象
	typeOfEnum := reflect.TypeOf(zero)

	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfEnum.Name(), typeOfEnum.Kind()) //Enum int
}


/******************指针与指针指向的元素*********************/
/*
	Go语言程序中对指针获取反射对象时，可以通过` reflect.Elem() `方法获取这个指针指向的元素类型，
这个获取过程被称为取元素，等效于对指针类型变量做了一个`*`操作
*/
func main103()  {
	type Cat struct {

	}

	cat := &Cat{}

	// 对指针变量获取反射类型信息。
	typeOfCat := reflect.TypeOf(cat)

	// 输出指针变量的类型名称和种类。
	// 反射中对所有指针变量的种类都是 Ptr，但需要注意的是，指针变量的类型名称是空，不是 *cat。
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // name:'' kind:'ptr'

	// 取指针类型的元素类型
	// 也就是 cat 类型。这个操作不可逆，不可以通过一个非指针类型获取它的指针类型。
	typeOfCat = typeOfCat.Elem()

	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // name:'Cat' kind:'struct'
}


/******************反射法则*********************/
/*
Go语言是一门静态类型的语言，每个变量都有一个静态类型，类型在编译的时候确定下来。
	type MyInt in
	var i int
	var j MyInt
变量 i 的类型是 int，变量 j 的类型是 MyInt，虽然它们有着相同的基本类型，但静态类型却不一样，
在没有类型转换的情况下，它们之间无法互相赋值。

	接口是一个重要的类型，它意味着一个确定的方法集合，一个接口变量可以存储任何实现了接口的方法的具体值（除了接口本身）
例如 io.Reader 和 io.Writer：
	// Reader is the interface that wraps the basic Read method.
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	// Writer is the interface that wraps the basic Write method.
	type Writer interface {
		Write(p []byte) (n int, err error)
	}
	如果一个类型声明实现了 Reader（或 Writer）方法，那么它便实现了 io.Reader（或 io.Writer），
这意味着一个 io.Reader 的变量可以持有任何一个实现了 Read 方法的的类型的值。
	var r io.Reader
	r = os.Stdin
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
不管变量 r 中的具体值是什么，r 的类型永远是 io.Reader，由于Go语言是静态类型的，r 的静态类型就是 io.Reader。

在接口类型中有一个极为重要的例子——空接口：
	interface{}
它表示了一个空的方法集，一切值都可以满足它，因为它们都有零值或方法。

	有人说Go语言的接口是动态类型，这是错误的，它们都是静态类型，虽然在运行时中，接口变量存储的值也许会变，但接口变量的类型是不会变的。
我们必须精确地了解这些，因为反射与接口是密切相关的。

	运行时反射是程序在运行期间检查其自身结构的一种方式，它是元编程 的一种，但是它带来的灵活性也是一把双刃剑，
过量的使用反射会使我们的程序逻辑变得难以理解并且运行缓慢，我们在这一节中就会介绍 Go 语言反射的三大法则，这能够帮助我们更好地理解反射的作用。
	1. 从接口值可反射出反射对象；
	2. 从反射对象可反射出接口值；
	3. 要修改反射对象，其值必须可设置；

### 第一法则 --- 从接口到反射对象。
	反射的第一条法则就是，我们能够将 Go 语言中的接口类型变量转换成反射对象，
上面提到的reflect.TypeOf 和 reflect.ValueOf 就是完成这个转换的两个最重要方法，
如果我们认为 Go 语言中的类型和反射类型是两个不同『世界』的话，那么这两个方法就是连接这两个世界的桥梁。
	> 注：这里反射类型指 reflect.Type 和 reflect.Value。

 */
func main104()  {
	type myInterface interface {}

	var a myInterface = new(myInterface)
	var b interface{} = new(interface{})

	typeOfa := reflect.TypeOf(a)
	typeOfb := reflect.TypeOf(b)

	fmt.Printf("a: %v, %v\n", typeOfa.Name(), typeOfa.Kind()) // a: , ptr
	fmt.Printf("b: %v, %v\n", typeOfb.Name(), typeOfb.Kind()) // b: , ptr

	name := "ixfosa"
	//  其中 `TypeOf` 获取了变量 `name` 的类型也就是 `string` 而 `ValueOf` 获取了变量的值 `ixfosa`，
	// 如果我们知道了一个变量的类型和值，那么也就意味着我们知道了关于这个变量的全部信息。
	fmt.Println("TypeOf name:", reflect.TypeOf(name))   // TypeOf name: string
	fmt.Println("ValueOf name:", reflect.ValueOf(name)) // ValueOf name: ixfosa
}

/*
从变量的类型上我们可以获当前类型能够执行的方法 `Method` 以及当前类型实现的接口等信息
	- 对于结构体，可以获取字段的数量并通过下标和字段名获取字段 `StructField`；
	- 对于哈希表，可以获取哈希表的 `Key` 类型；
	- 对于函数或方法，可以获得入参和返回值的类型；
	- …

	总而言之，使用 `TypeOf` 和 `ValueOf` 能够将 Go 语言中的变量转换成反射对象，在这时我们能够获得几乎一切跟当前类型相关数据和操作，
然后就可以用这些运行时获取的结构动态的执行一些方法。

	为什么是从接口到反射对象，如果直接调用 `reflect.ValueOf(1)`，看起来是从基本类型 `int` 到反射类型，
但是 `TypeOf` 和 `ValueOf` 两个方法的入参其实是 `interface{}` 类型。
	Go 语言的函数调用都是值传递的，变量会在方法调用前进行类型转换，也就是 `int` 类型的基本变量会被转换成 `interface{}` 类型，
这也就是第一条法则介绍的是 从接口到反射对象。
 */


/*****************## 第二法则**********************/
/*
	我们既然能够将接口类型的变量转换成反射对象类型，那么也需要一些其他方法将反射对象还原成成接口类型的变量，
`reflect `中的` Interface`方法就能完成这项工作：
	然而调用 `Interface` 方法我们也只能获得 `interface{}` 类型的接口变量，如果想要将其还原成原本的类型还需要经过一次强制的类型转换，
如下所示：
	v := reflect.ValueOf(1)
	v.Interface().(int)
	这个过程就像从接口值到反射对象的镜面过程一样，从接口值到反射对象需要经过从基本类型到接口类型的类型转换和从接口类型到反射对象类型的转换	，
反过来的话，所有的反射对象也都需要先转换成接口类型，再通过强制类型转换变成原始类型;
	当然不是所有的变量都需要类型转换这一过程，如果本身就是 `interface{}` 类型的，那么它其实并不需要经过类型转换，对于大多数的变量来说，
类型转换这一过程很多时候都是隐式发生的，只有在我们需要将反射对象转换回基本类型时才需要做显示的转换操作。
 */
func main106()  {
	var x int = 1
	v := reflect.ValueOf(x)
	fmt.Printf("type: %T\n", v) // type: reflect.Value
	fmt.Printf("type: %v, kind: %v\n", v.Type(), v.Kind()) // type: int, kind: int

	// x = int(v.Int())

	vv := v.Interface().(int)

	fmt.Printf("type: %T, value: %v\n", vv, vv) // type: int, value: 1
}

/******************第三法则*********************/
/*
	Go语言反射的最后一条法则是与值是否可以被更改相关的，如果我们想要更新一个 `reflect.Value`，那么它持有的值一定是可以被更新的
 */
func main107()  {
	x := 6
	v := reflect.ValueOf(x)

	fmt.Println("获取值: ", v.Int()) //获取值:  6

	// 设置值
	//v.SetInt(10)  // panic: reflect: reflect.Value.SetInt using unaddressable value
	//fmt.Println("x: ", x)

	// 出错的原因:
	// Go 语言的 [函数调用]都是传值的，所以我们得到的反射对象其实跟最开始的变量没有任何关系，
	// 没有任何变量持有复制出来的值，所以直接对它修改会导致崩溃。

	//想要修改原有的变量我们只能通过如下所示的方法，首先通过 `reflect.ValueOf` 获取变量指针，
	//然后通过 `Elem` 方法获取指针指向的变量并调用 `SetInt` 方法更新变量的值：

	// 调用 reflect.ValueOf 函数获取变量指针；
	// 调用 reflect.Value.Elem 方法获取指针指向的变量；
	// 调用 reflect.Value.SetInt 方法更新变量的
	v2 := reflect.ValueOf(&x)
	v2.Elem().SetInt(99)
	fmt.Println("x: ", x) // x:  99

	// 这种获取指针对应的 `reflect.Value` 并通过 `Elem` 方法迂回的方式就能够获取到可以被设置的变量，
	// 这一复杂的过程主要也是因为 Go 语言的函数调用都是值传递的，我们可以将上述代码理解成：
	// i := 1
	// v := &i
	// *v = 10
}

/*******************类型和值*******************/
/*
	Go 语言的 interface{} 类型在语言内部是通过 emptyInterface 这个结体来表示的，
其中的 rtype 字段用于表示变量的类型，另一个 word 字段指向内部封装的数据：
	type emptyInterface struct {
		typ  *rtype
		word unsafe.Pointer
	}

用于获取变量类型的 reflect.TypeOf 函数将传入的变量隐式转换成 emptyInterface 类型并获取其中存储的类型信息 rtype：
	func TypeOf(i interface{}) Type {
		eface := *(*emptyInterface)(unsafe.Pointer(&i))
		return toType(eface.typ)
	}

	func toType(t *rtype) Type {
		if t == nil {
			return nil
		}
		return t
	}

rtype 就是一个实现了 Type 接口的结构体，我们能在 reflect 包中找到如下所示的 reflect.rtype.String 方法帮助我们获取当前类型的名称等信息：
	func (t *rtype) String() string {
		s := t.nameOff(t.str).name()
		if t.tflag&tflagExtraStar != 0 {
			return s[1:]
		}
		return s
	}
reflect.TypeOf 函数的实现原理是将一个 interface{} 变量转换成了内部的 emptyInterface 表示，然后从中获取相应的类型信息。

	用于获取接口值 Value 的函数 reflect.ValueOf 实现也非常简单，在该函数中我们先调用了 reflect.escapes 函数保证当前值逃逸到堆上，
然后通过 reflect.unpackEface 方法从接口中获取 Value 结构体：
	func ValueOf(i interface{}) Value {
		if i == nil {
			return Value{}
		}
		escapes(i)
		return unpackEface(i)
	}

	func unpackEface(i interface{}) Value {
		e := (*emptyInterface)(unsafe.Pointer(&i))
		t := e.typ
		if t == nil {
			return Value{}
		}
		f := flag(t.Kind())
		if ifaceIndir(t) {
			f |= flagIndir
		}
		return Value{t, e.word, f}
	}
reflect.unpackEface 函数会将传入的接口转换成 emptyInterface 结构体，然后将具体类型和指针包装成 Value 结构体并返回。


	当我们想要将一个变量转换成反射对象时，Go 语言会在编译期间完成类型转换的工作，将变量的类型和值转换成了 interface{}
并等待运行期间使用 reflect 包获取接口中存储的信息。
 */



/*******************更新变量********************/
/*
	当我们想要更新一个 reflect.Value，就需要调用 reflect.Value.Set 方法更新反射对象，该方法会调用
reflect.flag.mustBeAssignable和 reflect.flag.mustBeExported 分别检查当前反射对象是否是可以被设置的以及字段是否是对外公开的：
	func (v Value) Set(x Value) {
		v.mustBeAssignable()
		x.mustBeExported()
		var target unsafe.Pointer
		if v.kind() == Interface {
			target = v.ptr
		}
		x = x.assignTo("reflect.Set", v.typ, target)
		typedmemmove(v.typ, v.ptr, x.ptr)
	}

reflect.Value.Set 方法会调用 reflect.Value.assignTo 并返回一个新的反射对象，这个返回的反射对象指针就会直接覆盖原始的反射变量。
	func (v Value) assignTo(context string, dst *rtype, target unsafe.Pointer) Value {
		...
		switch {
		case directlyAssignable(dst, v.typ):
			...
			return Value{dst, v.ptr, fl}
		case implements(dst, v.typ):
			if v.Kind() == Interface && v.IsNil() {
				return Value{dst, nil, flag(Interface)}
			}
			x := valueInterface(v, false)
			if dst.NumMethod() == 0 {
				*(*interface{})(target) = x
			} else {
				ifaceE2I(dst, x, target)
			}
			return Value{dst, target, flagIndir | flag(Interface)}
		}
		panic(context + ": value of type " + v.typ.String() + " is not assignable to type " + dst.String())
	}

reflect.Value.assignTo 会根据当前和被设置的反射对象类型创建一个新的 Value 结构体：
	如果两个反射对象的类型是可以被直接替换，就会直接将目标反射对象返回；
	如果当前反射对象是接口并且目标对象实现了接口，就会将目标对象简单包装成接口值；
在变量更新的过程中，reflect.Value.assignTo 返回的 reflect.Value 中的指针会覆盖当前反射对象中的指针实现变量的更新。

 */


/*****************实现协议**********************/
/*
	reflect 包还为我们提供了 reflect.rtypes.Implements 方法可以用于判断某些类型是否遵循特定的接口。
在 Go 语言中获取结构体的反射类型 reflect.Type 还是比较容易的，但是想要获得接口的类型就需要通过以下方式：
	reflect.TypeOf((*<interface>)(nil)).Elem()

 reflect.rtypes.Implements 方法的工作原理：
	func (t *rtype) Implements(u Type) bool {
		if u == nil {
			panic("reflect: nil type passed to Type.Implements")
		}
		if u.Kind() != Interface {
			panic("reflect: non-interface type passed to Type.Implements")
		}
		return implements(u.(*rtype), t)
	}

	reflect.rtypes.Implements 方法会检查传入的类型是不是接口，如果不是接口或者是空值就会直接 panic 中止当前程序。
在参数没有问题的情况下，上述方法会调用私有函数 reflect.implements 判断类型之间是否有实现关系：
	func implements(T, V *rtype) bool {
		t := (*interfaceType)(unsafe.Pointer(T))
		if len(t.methods) == 0 {
			return true
		}
		...
		v := V.uncommon()
		i := 0
		vmethods := v.methods()
		for j := 0; j < int(v.mcount); j++ {
			tm := &t.methods[i]
			tmName := t.nameOff(tm.name)
			vm := vmethods[j]
			vmName := V.nameOff(vm.name)
			if vmName.name() == tmName.name() && V.typeOff(vm.mtyp) == t.typeOff(tm.typ) {
				if i++; i >= len(t.methods) {
					return true
				}
			}
		}
		return false
	}
如果接口中不包含任何方法，就意味着这是一个空的接口，任意类型都自动实现该接口，这时就会直接返回 true。

类型实现接口
	在其他情况下，由于方法都是按照字母序存储的，reflect.implements 会维护两个用于遍历接口和类型方法的索引 i 和 j 判断类型是否实现了接口，
因为最多只会进行 n 次比较（类型的方法数量），所以整个过程的时间复杂度是 O(n)。
 */

// 如何判断一个类型是否实现了某个接口。假设我们需要判断如下代码中的 CustomError 是否实现了 Go 语言标准库中的 error 接口：
type CustomError struct {

}

func (c *CustomError) Error() string {
	return ""
}
func main108()  {
	typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	customErrorPtr := reflect.TypeOf(&CustomError{})
	customError := reflect.TypeOf(CustomError{})

	fmt.Println(customErrorPtr.Implements(typeOfError)) // true
	fmt.Println(customError.Implements(typeOfError)) // false

	// CustomError 类型并没有实现 error 接口；
	// *CustomError 指针类型实现了 error 接口；
}


/******************* 方法调用********************/
/*
	作为一门静态语言，如果我们想要通过 reflect 包利用反射在运行期间执行方法不是一件容易的事情，
下面的十几行代码就使用反射来执行 Add(0, 1) 函数：

使用反射来调用方法非常复杂，原本只需要一行代码就能完成的工作，现在需要十几行代码才能完成，但这也是在静态语言中使用动态特性需要付出的成本。
	func (v Value) Call(in []Value) []Value {
		v.mustBe(Func)
		v.mustBeExported()
		return v.call("Call", in)
	}

	reflect.Value.Call 方法是运行时调用方法的入口，它通过两个 MustBe 开头的方法确定了当前反射对象的类型是函数以及可见性，
随后调用 reflect.Value.call 完成方法调用，这个私有方法的执行过程会分成以下的几个部分：
	1.检查输入参数以及类型的合法性；
	2.将传入的 reflect.Value 参数数组设置到栈上；
	3.通过函数指针和输入参数调用函数；
	4.从栈上获取函数的返回值；
我们将按照上面的顺序分析使用 reflect 进行函数调用的几个过程。

1.参数检查
	参数检查是通过反射调用方法的第一步，在参数检查期间我们会从反射对象中取出当前的函数指针 unsafe.Pointer，
如果该函数指针是方法，那么我们就会通过 reflect.methodReceiver 函数获取方法的接受者和函数指针。

	func (v Value) call(op string, in []Value) []Value {
		t := (*funcType)(unsafe.Pointer(v.typ))
		...
		if v.flag&flagMethod != 0 {
			rcvr = v
			rcvrtype, t, fn = methodReceiver(op, v, int(v.flag)>>flagMethodShift)
		} else {
			...
		}
		n := t.NumIn()
		if len(in) < n {
			panic("reflect: Call with too few input arguments")
		}
		if len(in) > n {
			panic("reflect: Call with too many input arguments")
		}
		for i := 0; i < n; i++ {
			if xt, targ := in[i].Type(), t.In(i); !xt.AssignableTo(targ) {
				panic("reflect: " + op + " using " + xt.String() + " as type " + targ.String())
		}
	}
在上述方法中，上述方法还会检查传入参数的个数以及参数的类型与函数签名中的类型是否可以匹配，任何参数的不匹配都会导致整个程序的崩溃中止。

2. 准备参数
当我们已经对当前方法的参数完成验证之后，就会进入函数调用的下一个阶段，为函数调用准备参数，
在前面的章节函数调用中我们已经介绍过 Go 语言的函数调用惯例，函数或者方法在调用时，所有的参数都会被依次放置到栈上。

	nout := t.NumOut()
	frametype, _, retOffset, _, framePool := funcLayout(t, rcvrtype)

	var args unsafe.Pointer
	if nout == 0 {
		args = framePool.Get().(unsafe.Pointer)
	} else {
		args = unsafe_New(frametype)
	}
	off := uintptr(0)
	if rcvrtype != nil {
		storeRcvr(rcvr, args)
		off = ptrSize
	}
	for i, v := range in {
		targ := t.In(i).(*rtype)
		a := uintptr(targ.align)
		off = (off + a - 1) &^ (a - 1)
		n := targ.size
		...
		addr := add(args, off, "n > 0")
		v = v.assignTo("reflect.Value.Call", targ, addr)
		*(*unsafe.Pointer)(addr) = v.ptr
		off += n
	}
	1.通过 reflect.funcLayout 函数计算当前函数需要的参数和返回值的栈布局，也就是每一个参数和返回值所占的空间大小；
	2.如果当前函数有返回值，需要为当前函数的参数和返回值分配一片内存空间 args；
	3.如果当前函数是方法，需要向将方法的接受者拷贝到 args 内存中；
	4.将所有函数的参数按照顺序依次拷贝到对应 args 内存中
		1.使用 reflect.funcLayout 返回的参数计算参数在内存中的位置；
		2.将参数拷贝到内存空间中；
准备参数的过程是计算各个参数和返回值占用的内存空间并将所有的参数都拷贝内存空间对应的位置的过程，
该过程会考虑函数和方法、返回值数量以及参数类型带来的差异。

3.调用函数
	准备好调用函数需要的全部参数之后，就会通过以下的代码执行函数指针了。
我们会向该函数传入栈类型、函数指针、参数和返回值的内存空间、栈的大小以及返回值的偏移量：
	call(frametype, fn, args, uint32(frametype.size), uint32(retOffset))
上述函数实际上并不存在，它会在编译期间被链接到 runtime.reflectcall 这个用汇编实现的函数上

4.处理返回值
	当函数调用结束之后，就会开始处理函数的返回值：

	如果函数没有任何返回值，会直接清空 args 中的全部内容来释放内存空间；
	如果当前函数有返回值；
		1.将 args 中与输入参数有关的内存空间清空；
		2.创建一个 nout 长度的切片用于保存由反射对象构成的返回值数组；
		3.从函数对象中获取返回值的类型和内存大小，将 args 内存中的数据转换成 reflect.Value 类型并存储到切片中；
			var ret []Value
			if nout == 0 {
				typedmemclr(frametype, args)
				framePool.Put(args)
			} else {
				typedmemclrpartial(frametype, args, 0, retOffset)
				ret = make([]Value, nout)
				off = retOffset
				for i := 0; i < nout; i++ {
					tv := t.Out(i)
					a := uintptr(tv.Align())
					off = (off + a - 1) &^ (a - 1)
					if tv.Size() != 0 {
						fl := flagIndir | flag(tv.Kind())
						ret[i] = Value{tv.common(), add(args, off, "tv.Size() != 0"), fl}
					} else {
						ret[i] = Zero(tv)
					}
					off += tv.Size()
				}
			}
			return ret
		}
由 reflect.Value 构成的 ret 数组会被返回到上层，到这里为止使用反射实现函数调用的过程就结束了。
 */
func Add(x, y int) int {
	return x + y
}
func main109() {
	// 通过 reflect.ValueOf 获取函数 Add 对应的反射对象；
	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	// 根据反射对象 reflect.rtype.NumIn 方法返回的参数个数创建 argv 数组；
	argv := make([]reflect.Value, t.NumIn())

	// 多次调用 reflect.ValueOf 函数逐一设置 argv 数组中的各个参数；
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i + 1)
	}

	// 调用反射对象 Add 的 reflect.Value.Call 方法并传入参数列表；
	res := v.Call(argv)

	// 获取返回值数组、验证数组的长度以及类型并打印其中的数据；
	if len(res) != 1 || res[0].Kind() != reflect.Int {
		return
	}


	fmt.Println(res[0].Int()) // 3
}


/******************空接口与反射*********************/
/*
	反射可以在运行时动态获取程序的各种详细信息
	反射获取interface类型信息
 */
func reflect_type(x interface{})  {
	t := reflect.TypeOf(x)
	fmt.Println("类型是：", t)
	// kind()可以获取具体类型
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Printf("x is float64\n")
	case reflect.String:
		fmt.Println("string")
	}
}
func main1010()  {
	x := 66.6
	reflect_type(x)
}


/****************反射获取interface值信息***********************/
func reflect_value(x interface{}){
	v := reflect.ValueOf(x)
		fmt.Println("值是：", v)
	// kind()可以获取具体类型
	k := v.Kind()
	fmt.Println(v)
	switch k {
	case reflect.Float64:
		fmt.Printf("x = %v\n", v.Float())
	case reflect.String:
		fmt.Println("string")
	}
}

func main1011()  {
	x := 66.6
	reflect_value(x)
	// 值是： 66.6
	// 66.6
	// x = 66.6
}


/******************反射修改值信息*********************/
/*
### 判定及获取元素的相关方法
	使用 reflect.Value 取元素、取地址及修改值的属性方法请参考下表。
| 方法名         | 备  注                                                       |
| -------------- | ------------------------------------------------------------ |
| Elem() Value   | 取值指向的元素值，类似于语言层`*`操作。当值类型不是指针或接口时发生宕 机，空指针时返回 nil 的 Value |
| Addr() Value   | 对可寻址的值返回其地址，类似于语言层`&`操作。当值不可寻址时发生宕机 |
| CanAddr() bool | 表示值是否可寻址                                             |
| CanSet() bool  | 返回值能否被修改。要求值可寻址且是导出的字段                    |


### 值修改相关方法
	使用 reflect.Value 修改值的相关方法如下表所示。

| Set(x Value)        | 将值设置为传入的反射值对象的值                               |
| ------------------- | ------------------------------------------------------------ |
| Setlnt(x int64)     | 使用 int64 设置值。当值的类型不是 int、int8、int16、 int32、int64 时会发生宕机 |
| SetUint(x uint64)   | 使用 uint64 设置值。当值的类型不是 uint、uint8、uint16、uint32、uint64 时会发生宕机 |
| SetFloat(x float64) | 使用 float64 设置值。当值的类型不是 float32、float64 时会发生宕机 |
| SetBool(x bool)     | 使用 bool 设置值。当值的类型不是 bod 时会发生宕机            |
| SetBytes(x []byte)  | 设置字节数组 []bytes值。当值的类型不是 []byte 时会发生宕机   |
| SetString(x string) | 设置字符串值。当值的类型不是 string 时会发生宕机             |

以上方法，在 reflect.Value 的 CanSet 返回 false 仍然修改值时会发生宕机。


### 值可修改条件之一：可被寻址
	通过反射修改变量值的前提条件之一：这个值必须可以被寻址。简单地说就是这个变量必须能被修改。
### 值可修改条件之一：被导出
	结构体成员中，如果字段没有被导出，即便不使用反射也可以被访问，但不能通过反射修改
	panic: reflect: reflect.Value.SetInt using value obtained using unexported field

值的修改从表面意义上叫`可寻址`，换一种说法就是值必须“可被设置”。那么，想修改变量值，一般的步骤是：
	1. 取这个变量的地址或者这个变量所在的结构体已经是指针类型。
	2. 使用 reflect.ValueOf 进行值包装。
	3. 通过 Value.Elem() 获得指针值指向的元素值对象（Value），因为值对象（Value）内部对象为指针时，使用 set 设置时会报出宕机错误。
	4. 使用 Value.Set 设置值。
*/
func reflect_set_value(val interface{})  {
	v := reflect.ValueOf(val)

	k := v.Kind()
	fmt.Println("kind is ", k) // kind is  ptr
	switch k {
	case reflect.Float64:
		// 反射修改值
		v.SetFloat(6.9)
		fmt.Println("a is ", v.Float())
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	}
}

func main1012()  {
	x := 66.6
	reflect_set_value(&x) // case: 7.9
	// 反射认为下面是指针类型，不是float类型
	fmt.Println("main:", x) // main: 7.9
}

/*****************判断反射值的空和有效性**********************/
/*
| 方 法           | 说 明
| -------------- | ------------------------------------------------------------ |
| IsNil() bool   | 返回值是否为 nil。如果值类型不是通道（channel）、函数、接口、map、指针或 切片时发生 panic，类似于语言层的`v== nil`操作
| IsValid() bool | 判断值是否有效。 当值本身非法时，返回 false，例如 reflect Value不包含任何值，值为 nil 等。

 */
// 反射值对象的零值和有效性判断：
func main1019()  {
	// *int的空指针
	var a *int
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil()) // var a *int: true

	// nil值
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid()) // nil: false

	// *int类型的空指针
	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid()) // (*int)(nil): false

	// 实例化一个结构体
	s := struct{}{}
	// 尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid()) //不存在的结构体成员: false
	// 尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(s).MethodByName("").IsValid()) //不存在的结构体方法: false

	// 实例化一个map
	m := map[int]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid()) //不存在的键： false
}


/******************结构体与反射*********************/
/*
### 使用反射获取结构体的成员类型
	任意值通过` reflect.TypeOf() `获得反射对象信息后，如果它的类型是结构体，
可以通过反射值对象 reflect.Type 的 `NumField() `和` Field() `方法获得结构体成员的详细信息。

与成员获取相关的 reflect.Type 的方法如下表所示。
| 方法                                                        | 说明                                                         |
| ----------------------------------------------------------- | ------------------------------------------------------------ |
| Field(i int) StructField                                    | 根据索引返回索引对应的结构体字段的信息，当值不是结构体或索引超界时发生宕机 |
| NumField() int                                              | 返回结构体成员字段数量，当类型不是结构体或索引超界时发生宕机 |
| FieldByName(name string) (StructField, bool)                | 根据给定字符串返回字符串对应的结构体字段的信息，没有找到时 bool 返回 false，当类型不是结构体或索引超界时发生宕机 |
| FieldByIndex(index []int) StructField                       | 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息，没有找到时返回零值。当类型不是结构体或索引超界时发生宕机 |
| FieldByNameFunc(match func(string) bool) (StructField,bool) | 根据匹配函数匹配需要的字段，当值不是结构体或索引超界时发生宕机 |

#### 结构体字段类型
	reflect.Type 的` Field() `方法返回` StructField `结构，这个结构描述结构体的成员信息，
通过这个信息可以获取成员与结构体的关系，如偏移、索引、是否为匿名字段、结构体标签（StructTag）等，
而且还可以通过 StructField 的 Type 字段进一步获取结构体成员的类型信息。

StructField 的结构如下：
	type StructField struct {
		Name string          // 字段名
		PkgPath string       // 字段路径
		Type      Type       // 字段反射类型对象
		Tag       StructTag  // 字段的结构体标签
		Offset    uintptr    // 字段在结构体中的相对偏移
		Index     []int      // Type.FieldByIndex中的返回的索引值
		Anonymous bool       // 是否为匿名字段
	}

字段说明如下：
	- Name：为字段名称。
	- PkgPath：字段在结构体中的路径。
	- Type：字段本身的反射类型对象，类型为 reflect.Type，可以进一步获取字段的类型信息。
	- Tag：结构体标签，为结构体字段标签的额外信息，可以单独提取。
	- Index：FieldByIndex 中的索引顺序。
	- Anonymous：表示该字段是否为匿名字段。
 */

// 查看类型、字段和方法
type User struct {
	Id int
	Name string
	Age int
}

// 绑方法
func (u User) Hello()  {
	fmt.Println("Hello")
}

func Poni(o interface{})  {
	t := reflect.TypeOf(o)
	fmt.Println("类型: ", t) // 类型:  main.User
	fmt.Println("字符串类型: ", t.Name()) //字符串类型:  User

	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println("v:", v) //v: {21 ixfosa 22}

	// 获取所有属性
	// 获取结构体字段个数 t.NumField
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v\n", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val :", val)
	}

	fmt.Println("=================方法====================")

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name) //Hello
		fmt.Println(m.Type) //func(main.User)
	}
}

func main1013()  {
	user := User{
		Id: 21,
		Name: "ixfosa",
		Age: 22,
	}
	Poni(user)
}

/******************查看匿名字段*********************/
// 匿名字段
type Boy struct {
	User
	Addr string
}

func main1014() {
	m := Boy{User{21, "ixfosa", 22}, "nc"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	// Anonymous：匿名
	fmt.Printf("%#v\n", t.Field(0))
	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0)) // main.User{Id:21, Name:"ixfosa", Age:22}
}

/*******************修改结构体的值********************/
func SetValue(o interface{})  {
	v := reflect.ValueOf(o)
	v = v.Elem()
	// 取字段
	f := v.FieldByName("Name")

	if f.Kind() == reflect.String {
		f.SetString("zhong")
	}
}
func main1015()  {
	user := &User{
		Id: 21,
		Name: "ixfosa",
		Age: 22,
	}
	SetValue(user)

	fmt.Println(user) // &{21 zhong 22}
}


/******************获取字段的tag*********************/
/*
	通过 reflect.Type 获取结构体成员信息 reflect.StructField 结构中的` Tag `被称为结构体标签（StructTag）。
结构体标签是对结构体字段的额外信息标签。

	JSON、BSON 等格式进行序列化及对象关系映射（Object Relational Mapping，简称 `ORM`）系统都会用到结构体标签，
这些系统使用标签设定字段在处理时应该具备的特殊属性和可能发生的行为。这些信息都是静态的，无须实例化结构体，可以通过反射获取到。

#### 结构体标签的格式
Tag 在结构体字段后方书写的格式如下：
	`key1:"value1" key2:"value2"`
结构体标签由一个或多个键值对组成；键与值使用冒号分隔，值用双引号括起来；键值对之间使用`一个空格分隔。`

#### 从结构体标签中获取值
StructTag 拥有一些方法，可以进行 Tag 信息的解析和提取，如下所示：

- func (tag StructTag) Get(key string) string：
	根据 Tag 中的键获取对应的值，例如``key1:"value1" key2:"value2"``的 Tag 中，可以传入“key1”获得“value1”。

- func (tag StructTag) Lookup(key string) (value string, ok bool)：
	根据 Tag 中的键，查询值是否存在。
 */

type Student struct {
	Name string `json:"name1" db:"name2"`
}

func main1017()  {
	var s Student
	v := reflect.ValueOf(&s)
	// 类型
	t := v.Type()
	// 获取字段
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json")) // name1
	fmt.Println(f.Tag.Get("db"))  // name2
}


/*****************反射调用函数**********************/
/*
	如果反射值对象（reflect.Value）中值的类型为`函数`时，可以通过 reflect.Value 调用该函数。
使用反射调用函数时，需要将参数使用反射值对象的切片` []reflect.Value` 构造后传入` Call() `方法中，
调用完成时，函数的返回值通过` []reflect.Value` 返回。

	下面的代码声明一个加法函数，传入两个整型值，返回两个整型值的和。将函数保存到反射值对象（reflect.Value）中，
然后将两个整型值构造为反射值对象的切片（[]reflect.Value），使用 Call() 方法进行调用。如果反射值对象（reflect.Value）
中值的类型为函数时，可以通过 reflect.Value 调用该函数。使用反射调用函数时，需要将参数使用反射值对象的切片
[]reflect.Value 构造后传入Call() 法中，调用完成时，函数的返回值通过` []reflect.Value` 返回。

	下面的代码声明一个加法函数，传入两个整型值，返回两个整型值的和。将函数保存到反射值对象（reflect.Value）中，
然后将两个整型值构造为反射值对象的切片（[]reflect.Value），使用 Call() 方法进行调用。

	反射调用函数的过程需要构造大量的 reflect.Value 和中间变量，对函数参数值进行逐一检查，还需要将调用参数复制到调用函数的参数内存中。
调用完毕后，还需要将返回值转换为 reflect.Value，用户还需要从中取出调用值。因此，反射调用函数的性能问题尤为突出，不建议大量使用反射函数调用。
 */

func add(a, b int) int {
	return a + b
}

func main1018()  {

	// 将函数包装为反射值对象
	funcVal := reflect.ValueOf(add)
	// 构造函数参数, 传入两个整型值
	args := []reflect.Value{reflect.ValueOf(6), reflect.ValueOf(9)}
	// 反射调用函数
	retList := funcVal.Call(args)

	// 获取第一个返回值, 取整数值
	fmt.Println(retList[0].Int()) //15
}

/*****************调用方法**********************/
/*
	整体与调用函数一致，额外的需要先通过对象的值反射获取对象方法的反射对象，再使用 Call() 调用
 */
func (u User) SayHello(name string) {
	fmt.Println(u.Name, "say: Hello：", name)
}
func main1016()  {
	user := User{
		Id: 21,
		Name: "ixfosa",
		Age: 22,
	}
	v := reflect.ValueOf(user)
	// 获取方法
	m := v.MethodByName("SayHello")
	args := []reflect.Value{reflect.ValueOf("zhong")}
	// 没参数的情况下：var args2 []reflect.Value
	// 调用方法，需要传入方法的参数
	m.Call(args) // ixfosa say: Hello： zhong
}


/******************通过类型信息创建实例*********************/
/*
	当已知 reflect.Type 时，可以动态地创建这个类型的实例，实例的类型为指针。
例如 reflect.Type 的类型为 int 时，创建 int 的指针，即 *int
 */
func main1020()  {
	var a int
	// 取变量a的反射类型对象
	typeOfA := reflect.TypeOf(a)
	// 根据反射类型对象创建类型实例
	aIns := reflect.New(typeOfA) //
	// 输出Value的类型和种类
	fmt.Println(aIns.Type(), aIns.Kind()) //代码输出如下：*int ptr //*int ptr
}

func main()  {
	type Person struct {
		Id int
		name string
		age int
	}

	//t := reflect.TypeOf(&Person{})
	t := reflect.TypeOf(Person{})
	// func New(typ Type) Value 创建这个类型的实例，实例的类型为指针。
	p := reflect.New(t)
	fmt.Println(p, p.Type(), p.Kind())

	person := p.Interface().(*Person)
	person.Id = 21
	fmt.Println(person)
	// &{0  0} *main.Person ptr
	// &{21  0}
}
