package main

import "fmt"

// 下面这段代码能否编译通过？如果通过，输出什么？
	// 参考答案及解析：能，输出m1 m2，
	// 第 2 行代码，第 3 行代码是，
type User struct {

}

type User1 User     // 基于类型 User 创建了新类型 User1
type User2 = User   // 创建了 User 的类型别名 User2， 注意使用 = 定义类型别名。

func (u User1) m1()  {
	fmt.Println("m1")
}

func (u User) m2()  {
	fmt.Println("m2")
}


func main() {
	var u1 User1
	var u2 User2  // 因为 User2 是别名，完全等价于 User，所以 User2 具有 User 所有的方法。
	u1.m1()
	// u2.m2()    // // 但是 u1.m2() 是不能执行的，因为 User1 没有定义该方法。
	u2.m2()
}
