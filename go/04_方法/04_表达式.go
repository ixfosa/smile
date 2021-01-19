package main

import "fmt"


/******************表达式*********************/
/*
Golang 表达式 ：根据调用者不同，方法分为两种表现形式:
    instance.method(args...) ---> <type>.func(instance, args...)
前者称为 method value，后者 method expression。
两者都可像普通函数那样赋值和传参，区别在于 method value 绑定实例，而 method expression 则须显式传参。
 */
type User04 struct {
	id int
	name string
}

func (self *User04) test041() {
	fmt.Printf("%p, %v\n", self, self)
}

func (self User04) test042() {
	fmt.Printf("%p, %v\n", &self, self)
}

func main41() {
	u := User04{1, "ixfosa"}
	u.test041() //0xc0420023e0, &{1 ixfosa}

	mValue := u.test041
	mValue() // 隐式传递 receiver

	mExpression := (*User04).test041
	mExpression(&u) // 显式传递 receiver
	//0xc0420023e0, &{1 ixfosa}
	//0xc0420023e0, &{1 ixfosa}

	mValue2 := u.test042
	mValue2()

	mExpression2 := (User04).test042
	mExpression2(u)
	//0xc042002480, {1 ixfosa}
	//0xc0420024c0, {1 ixfosa}

	//需要注意，method value 会复制 receiver。
	u2 := User04{1, "long"}
	mValue3 := u2.test042// 立即复制 receiver，因为不是指针类型，不受后续修改影响。

	u2.id, u2.name = 99, "zhong"
	mValue3() //0xc042002500, {1 long}
}

/***************************************/
type Person struct {
	name string
}

func (p Person) setVName(name string) {
	p.name = name
}

func (p *Person) setPName(name string) {
	p.name = name
}

func main() {
	p1 := &Person{"ixfosa"}
	testV1 := (Person).setVName
	testV1(*p1, "long")
	fmt.Println("testV1: ", p1)  //testV1:  {ixfosa}

	/*testP1 := (Person).setPName
	testP1(*p1, "zhong") // invalid method expression Person.setPName (needs pointer receiver: (*Person).setPName)
	fmt.Println("testP1: ", p1)  //testP1:  {zhong}*/

	fmt.Println()
	testV2 := (*Person).setVName
	testV2(p1, "long")
	fmt.Println("testV2: ", p1)  //testV2:  {ixfosa}

	testP2 := (*Person).setPName
	testP2(p1, "zhong")
	fmt.Println("testP2: ", p1)  //testP2:  {zhong}
}
func main44()  {
	p1 := Person{"ixfosa"}
	testV1 := (Person).setVName
	testV1(p1, "long")
	fmt.Println("testV1: ", p1)  //testV1:  {ixfosa}

	//testP1 := (Person).setPName
	//testP1(p1, "zhong") //invalid method expression Person.setPName (needs pointer receiver: (*Person).setPName)
	//testP1(&p1, "zhong")
	//fmt.Println("testP1: ", p1)  //testP1:  {zhong}

	fmt.Println()
	testV2 := (*Person).setVName
	testV2(&p1, "long")
	fmt.Println("testV1: ", p1)  //testV1:  {ixfosa}

	testP2 := (*Person).setPName
	testP2(&p1, "zhong")
	fmt.Println("testP1: ", p1)  //testP1:  {zhong}

}
func main43() {
	p1 := Person{"ixfosa"}
	testV1 := p1.setVName
	testV1("long")
	fmt.Println("testV1: ", p1)  //testV1:  {ixfosa}

	testP1 := p1.setPName
	testP1("zhong")
	fmt.Println("testP1: ", p1)  //testP1:  {zhong}

	fmt.Println()

	p2 := &Person{"ixfosa"}
	testV2 := p2.setVName
	testV2("long")
	fmt.Println("testV2: ", p2)  //testV2:  &{ixfosa}

	testP2 := p2.setPName
	testP2("long")
	fmt.Println("testP2: ", p2)  //testP2:  &{long}


}


func main42() {
	p1 := Person{"ixfosa"}
	fmt.Println("p1: ", p1) //p1:  {ixfosa}

	p1.setVName("long")
	fmt.Println("p1.setVName: ", p1)  //p1.setVName:  {ixfosa}

	p1.setPName("zhong")
	fmt.Println("p1.setPName: ", p1) //p1.setPName:  {zhong}


	fmt.Println()

	p2 := &Person{"ixfosa"}
	fmt.Println("p2: ", p2) //p2:  &{ixfosa}

	p2.setVName("long")
	fmt.Println("p2.setVName: ", p2) //p2.setVName:  &{ixfosa}

	p2.setPName("zhong")
	fmt.Println("p1.setPName: ", p2) //p1.setPName:  &{zhong}

	fmt.Println()
}

/***************************************/

/*

 */


/***************************************/

/*

 */