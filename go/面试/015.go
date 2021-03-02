package main

// 下面这段代码能否编译通过？如果可以，输出什么？
func GetValue() int {
	return 1
}
/*
func main() {
	i := GetValue()

	switch i.(type) {  //  cannot type switch on non-interface value i (type int)
	case int:
		println("int")
	case string:
		fmt.Println("string")
	case interface{}:
		fmt.Println("interface{}")
	default:
		println("unknown")
	}
}
*/

// 参考答案及解析：编译失败。考点：类型选择，类型选择的语法形如：i.(type)，其中 i 是接口，type 是固定关键字，
// 需要注意的是，只有接口类型才可以使用类型选择。