package main

import "fmt"

/*
go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段
*/

//人
type Person struct {
	name string
	age int8
	sex string
}

type Student struct {
	Person
	id string
	addr string
}

func main011() {
	//初始化
	stu := Student{Person{"ixfosa", 22, "man"}, "202099770821", "NanChang"}
	fmt.Println(stu) //{{ixfosa 22 man} 202099770821 NanChang}
}

type Student2 struct {
	Person
	id string
	addr string
	//同名字段
	name string
}

func main012() {

	var s Student2
	s.name = "ixfosa" // 给自己字段赋值了
	fmt.Println(s) //{{ 0 }   ixfosa}

	s.Person.name = "long" //给父类同名字段赋值
	fmt.Println(s) //{{long 0 }   ixfosa}
}



/*
有的内置类型和自定义类型都是可以作为匿名字段去使用
 */

// 自定义类型
type mystr string

// 学生
type Student3 struct {
	Person
	int
	mystr
}

func main013()  {
	s := Student3{Person{"long", 21, "woman"}, 9, "hello"}
	fmt.Println(s) //{{long 21 woman} 9 hello}
}



/*
指针类型匿名字段
 */
// 学生
type Student4 struct {
	*Person
	id int
	addr string
}

func main()  {
	s := Student4{&Person{"ixfosa", 22, "man"}, 9, "nanchang"}
	fmt.Println(s)
	fmt.Println(s.name)
	fmt.Println(s.Person.name)
	//{0xc042078060 9 nanchang}
	//ixfosa
	//ixfosa
}


/***************************************/

/*

 */