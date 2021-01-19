package main

import "fmt"

/*******************匿名字段********************/
/*
	Golang匿名字段 ：可以像字段成员那样访问匿名字段方法，编译器负责查找。
 */
type User02 struct {
	id int
	name string
}

type Manager struct {
	User02
}

func (self *User02) ToString() string {
	return fmt.Sprintf("User：%p, %v\n", self, self)
}

func main02() {
	m := Manager{User02{1, "ixfosa"}}
	fmt.Printf("Manager:%p \n", &m)
	fmt.Println(m.ToString())
	//Manager:0xc0420023e0
	//User：0xc0420023e0, &{1 ixfosa}
}

/***************************************/
/*
通过匿名字段，可获得和继承类似的复用能力。依据编译器查找次序，只需在外层定义同名方法，就可以实现 "override"。
 */
type Manager02 struct {
	User02
	title string
}

func (self *Manager02) ToString() string {
	return fmt.Sprintf("User：%p, %v\n", self, self)
}

func main() {
	m := Manager02{User02{1, "ixfosa"}, "go"}
	fmt.Println(m.ToString())
	fmt.Println(m.User02.ToString())
	//User：0xc042076060, &{{1 ixfosa} go}
	//User：0xc042076060, &{1 ixfosa}
}
/***************************************/

/*

 */


/***************************************/

/*

 */