package main

func main() {
	
}

// 下面这段代码有什么缺陷
// func funcMui(x,y int)(sum int, error){ // syntax error: mixed named and unnamed function parameters
// 	return x+y,nil
// }

// 第二个返回值没有命名。

// 参考解析：
// 在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。
// 如果有多个返回值必须加上括号()；如果只有一个返回值且命名也必须加上括号()。
// 这里的第一个返回值有命名 sum，第二个没有命名，所以错误。


// true
func funcMui1(x,y int)(int, error){ // syntax error: mixed named and unnamed function parameters
	return x+y,nil
}

func funcMui13(x,y int)(sum int, err error){ // syntax error: mixed named and unnamed function parameters
	return x+y,nil
}