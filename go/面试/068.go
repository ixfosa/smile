package main

import "fmt"

/*
1.关于 bool 变量 b 的赋值，下面错误的用法是？
	A. b = true
	B. b = 1
	C. b = bool(1)
	D. b = (1 == 2)
参考答案及解析：BC。


关于变量的自增和自减操作，下面语句正确的是？
	A.
		i := 1
		i++

	B.
		i := 1
		j = i++

	C.
		i := 1
		++i

	D.
		i := 1
		i--
参考答案及解析：AD。知识点：自增自减操作。

i++ 和 i–- 在 Go 语言中是语句，不是表达式，因此不能赋值给另外的变量。此外没有 ++i 和 –-i。



3.关于GetPodAction定义，下面赋值正确的是
	type Fragment interface {
			Exec(transInfo *TransInfo) error
	}

	type GetPodAction struct {
	}

	func (g GetPodAction) Exec(transInfo *TransInfo) error {
			...
			return nil
	}

	A. var fragment Fragment = new(GetPodAction)

	B. var fragment Fragment = GetPodAction

	C. var fragment Fragment = &GetPodAction{}

	D. var fragment Fragment = GetPodAction{}

参考答案及解析：ACD。
 */


type Fragment interface {
	Exec() error
}

type GetPodAction struct {

}

func (g GetPodAction) Exec() error {
	return nil
}

func main() {
	// var fragment Fragment = new(GetPodAction) // &{}

	// var fragment Fragment = GetPodAction{} // {}

	var fragment Fragment = &GetPodAction{} // &{}
	fmt.Println(fragment)

	//  cannot use GetPodAction literal (type *GetPodAction) as type GetPodAction in assignment
	// var getPodAction GetPodAction = &GetPodAction{}
}
