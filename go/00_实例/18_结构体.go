package main

import "fmt"

// Go 的结构体 是各个字段字段的类型的集合。
type person struct {
	name string
	age int
}
func main() {
	fmt.Println(person{"Zi Qing", 18}) // {Zi Qing 18}

	fmt.Println(person{name: "long", age: 22}) // {long 22}

	// & 前缀生成一个结构体指针
	fmt.Println(&person{name: "zhong", age: 21}) // &{zhong 21}

	p := person{name: "fo", age: 22}
	fmt.Println("p = ", p) // p =  {fo 22}

	p2 := p
	// 使用点来访问结构体字段。
	p2.age = 23
	fmt.Println("p = ", p) // p =  {fo 22}
	fmt.Println("p2 = ", p2) // p2 =  {fo 23}

	// 可以对结构体指针使用. - 指针会被自动解引用
	p3 := &p
	p3.name = "Zi Qing"
	fmt.Println("p = ", p) // p =  {Zi Qing 22}
	fmt.Println("p3 = ", p3) // p =  &{Zi Qing 22}
}
