package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

/*-----------------自定义类型 -------------------*/
/*
在Go语言中有一些基本的数据类型，如string、整型、浮点型、布尔等数据类型，Go语言中可以使用type关键字来定义自定义类型。
自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。例如：
    //将MyInt定义为int类型
    type MyInt int
通过Type关键字的定义，MyInt就是一种新的类型，它具有int的特性。
 */


/*-----------------类型别名-------------------*/
/*
类型别名是Go1.9版本添加的新功能。
类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。
就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。
	 type TypeAlias = Type

我们之前见过的rune和byte就是类型别名，他们的定义如下：
	type byte = uint8
	type rune = int32
 */


/*------------------类型定义和类型别名的区别------------------*/
//类型别名与类型定义表面上看只有一个等号的差异，我们通过下面的这段代码来理解它们之间的区别。
func main151() {
	type NewInt int
	type MyInt = int
	var a NewInt
	var b MyInt
	fmt.Printf("a type:%T\n", a) //a type:main.NewInt
	fmt.Printf("b type:%T\n", b) //b type:int
	//结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。
	//b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。
}


/*-------------------结构体-----------------*/
/*
Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，
这时候再用单一的基本数据类型明显就无法满足需求了，Go语言提供了一种自定义数据类型，
可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct。 也就是我们可以通过struct来定义自己的类型了。

Go语言中通过struct来实现面向对象。
 */


/*-----------------结构体的定义-------------------*/
/*
使用type和struct关键字来定义结构体，具体代码格式如下：

    type 类型名 struct {
        字段名 字段类型
        字段名 字段类型
        …
    }
其中：
    1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
    2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
    3.字段类型：表示结构体字段的具体类型。

举个例子，我们定义一个Person（人）结构体，代码如下：
    type person struct {
        name string
        city string
        age  int8
    }

同样类型的字段也可以写在一行，
    type person1 struct {
        name, city string
        age        int8
    }
这样我们就拥有了一个person的自定义类型，它有name、city、age三个字段，分别表示姓名、城市和年龄。
这样我们使用这个person结构体就能够很方便的在程序中表示和存储人信息了。

语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的。比如一个人有名字、年龄和居住城市等，本质上是一种聚合型的数据类型
 */


/*------------------结构体实例化------------------*/
/*
只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。
    var 结构体实例 结构体类型
 */


/*-------------------基本实例化-----------------*/
func main152() {
	type person struct {
		name string
		city string
		age int8
	}
	var p person
	fmt.Println(p) //{  0}
	p.name = "ixfosa"
	p.city = "NanChang"
	p.age = 22
	fmt.Printf("p=%v\n", p)  //p={ixfosa NanChang 22}
	fmt.Printf("p=%#v\n", p) //p=main.person{name:"ixfosa", city:"NanChang", age:22}
}


/*----------------匿名结构体	--------------------*/
func main153() {
	//在定义一些临时数据结构等场景下还可以使用匿名结构体。
	var user struct{name string; age int8}
	user.name = "ixfosa"
	user.age = 22
	fmt.Println(user) //{ixfosa 22}
}



/*--------------------创建指针类型结构体----------------*/
/*
通过使用new关键字对结构体进行实例化，得到的是结构体的地址。 格式如下：
 	var p = new(person)
    fmt.Printf("%T\n", p)     //*main.person
    fmt.Printf("p=%#v\n", p) //p=&main.person{name:"", city:"", age:0}
从打印的结果中我们可以看出p是一个结构体指针。
需要注意的是在Go语言中支持对结构体指针直接使用.来访问结构体的成员。
 */
func main154()  {
	type person struct {
		name string
		city string
	}
	var p = new(person)
	fmt.Printf("p=%#v, type:%T\n", p, p) //p=&main.person{name:"", city:""}, type:*main.person

	p.name = "long"
	p.city = "NanChang"
	fmt.Printf("p=%#v, type:%T\n", p, p) //p=&main.person{name:"long", city:"NanChang"}, type:*main.person
}




/*-----------------取结构体的地址实例化-------------------*/
func main155() {
	//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	type person struct {
		name string
		city string
	}

	var p = &person{}
	fmt.Printf("%T\n", p)     //*main.person
	fmt.Printf("p=%#v\n", p) //p=&main.person{name:"", city:""}
	p.name = "zhong"   //p.name = "zhong" 其实在底层是(*p).name = "zhong"，这是Go语言帮我们实现的语法糖。 指针解引用
	p.city = "NanChang"
	fmt.Printf("p=%#v\n", p) //p=&main.person{name:"zhng", city:"NanChang"}
}



/*------------------结构体初始化------------------*/
func main156() {
	type person struct {
		name string
		city string
	}
	var p person
	fmt.Println(p) //{ }

	//使用键值对初始化
	//使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
	p1 := person{
		name: "ixfosa",
		city: "NanChang",
	}
	fmt.Println("用键值对初始化 p1", p1) //用键值对初始化 p1 {ixfosa NanChang}

	//也可以对结构体指针进行键值对初始化，例如：
	p2 := &person{
		name: "long",
		city: "NanChang",
	}
	fmt.Println("对结构体指针进行键值对初始化 p2", p2) //对结构体指针进行键值对初始化 p2 &{long NanChang}

	//当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	p3 := &person{
		name: "Zhong",
	}
	fmt.Println("对结构体指针某个值初始化 p3", p3) //对结构体指针某个值初始化 p3 &{Zhong }
}


/*-----------------使用值的列表初始化-------------------*/
/*
使用这种格式初始化时，需要注意：
    1.必须初始化结构体的所有字段。
    2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
    3.该方式不能和键值初始化方式混用。
*/
func main158()  {
	//初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：
	type person struct {
		name string
		city string
	}

	p := &person{
		"Long",
		"NanChang",
	}
	fmt.Printf("使用值的列表初始化 p: %#v", p) //使用值的列表初始化 p: &main.person{name:"Long", city:"NanChang"}
}



/*-----------------结构体内存布局-------------------*/
func mai159n() {
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := &test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a: %p\n", &n.a)
	fmt.Printf("n.b: %p\n", &n.b)
	fmt.Printf("n.c: %p\n", &n.c)
	fmt.Printf("n.d: %p\n", &n.d)

	//n.a: 0xc04205e080
	//n.b: 0xc04205e081
	//n.c: 0xc04205e082
	//n.d: 0xc04205e083

	type test1 struct {
		a int
		b int
		c int
		d int
	}
	nn := &test1{
		1, 2, 3, 4,
	}

	fmt.Printf("nn.a: %p\n", &nn.a)
	fmt.Printf("nn.b: %p\n", &nn.b)
	fmt.Printf("nn.c: %p\n", &nn.c)
	fmt.Printf("nn.d: %p\n", &nn.d)
	//nn.a: 0xc042010340
	//nn.b: 0xc042010348
	//nn.c: 0xc042010350
	//nn.d: 0xc042010358
}

/*------------------面试题------------------*/
func main1510() {
	type student struct {
		name string
		age int
	}

	m := make(map[string]*student)
	stus := []student{
		{"ixfosa", 22},
		{"Long", 21},
		{"Zhong", 19},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}

	for k, v := range m	{
		fmt.Println(k, "=>", v.name)
	}
	//Zhong => Zhong
	//ixfosa => Zhong
	//Long => Zhong
}



/*----------------构造函数--------------------*/
/*
Go语言的结构体没有构造函数，我们可以自己实现。
例如，下方的代码就实现了一个person的构造函数。
因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
 */
type person struct {
	name string
	city string
	age int
}

func NewPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age: age,
	}
}

func main1511() {
	p := NewPerson("ixfosa", "NanChang", 22)
	fmt.Printf("p= %#v", p)  //p= &main.person{name:"ixfosa", city:"NanChang", age:22}
}


/*------------------方法和接收者------------------*/
/*
Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。
接收者的概念就类似于其他语言中的 this 或者 self。
方法的定义格式如下：
    func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
        函数体
    }

其中，
    1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
    2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
    3.方法名、参数列表、返回参数：具体格式与函数定义相同。
 */
func (p person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}
//方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
func main1512() {
	p := NewPerson("ixfosa", "NanChang", 22)
	p.Dream() //ixfosa的梦想是学好Go语言！
}

/*------------------指针类型的接收者------------------*/
/*
指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。
这种方式就十分接近于其他语言中面向对象中的this或者self。 例如我们为Person添加一个SetAge方法，来修改实例变量的年龄。

	// SetAge 设置p的年龄
    // 使用指针接收者
    func (p *Person) SetAge(newAge int8) {
        p.age = newAge
    }
 */
func (p *person) setAge(age int)  {
	p.age = age
}
func main1513() {
	p := NewPerson("ixfosa", "NanChang", 22)
	p.setAge(9)
	fmt.Printf("p= %#v\n", p)  //p= &main.person{name:"ixfosa", city:"NanChang", age:9}
	p.setAge(6)
	fmt.Printf("p= %#v\n", p)  //p= &main.person{name:"ixfosa", city:"NanChang", age:6}
}




/*-------------------值类型的接收者-----------------*/
/*
当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
 */
func (p person) setName(name string)  {
	p.name = name
}

func main1514() {
	p := NewPerson("ixfosa", "NanChang", 22) //p= &main.person{name:"ixfosa", city:"NanChang", age:22}
	fmt.Printf("p= %#v\n", p)

	p.setName("Long")
	fmt.Printf("set after p= %#v\n", p) //set after p= &main.person{name:"ixfosa", city:"NanChang", age:22}
}

/*-----------------什么时候应该使用指针类型接收者-------------------*/
/*
 	1.需要修改接收者中的值
    2.接收者是拷贝代价比较大的大对象
    3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
 */



/*----------------任意类型添加方法--------------------*/
/*
在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。

 */
//举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

//MyInt 将int定义为自定义MyInt类型
type MyInt int

func (myInt MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

func (myInt MyInt) printValue() {
	fmt.Println("value = ", myInt)
}

func main1515() {
	var i MyInt
	i.SayHello()
	i.printValue() //value =  0

	i = 6
	i.printValue() //value =  6
	fmt.Printf("%#v  %T\n", i, i) //100  main.MyInt
}




/*-----------------结构体的匿名字段	-------------------*/
/*
结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。

匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
 */
func main1516() {
	type Person struct {
		string
		int
	}
	p := &Person{
		"ixfosa",
		22,
	}
	fmt.Println(p.string , "-", p.int) //ixfosa - 22
	p.string = "smile"
	fmt.Println(p.string , "-", p.int) //smile - 22
}


/*----------------嵌套结构体--------------------*/
func main1517() {

	type Address struct {
		Province string
		City string
	}

	type Person struct {
		Name string
		Age int
		Addr Address
	}

	p := Person{
		Name: "ixfosa",
		Age: 22,
		Addr: Address{
			Province: "江西",
			City: "南昌",
		},
	}
	fmt.Printf("p= %#v\n", p) //p= main.Person{Name:"ixfosa", Age:22, Addr:main.Address{Province:"江西", City:"南昌"}}
}


/*----------------嵌套匿名结构体--------------------*/
func main1518() {
	type Address struct {
		Province string
		City string
	}

	type Person struct {
		Name string
		Age int
		Address  //匿名结构体
	}

	var p Person
	p.Name = "ixfosa"
	p.Age = 22
	//当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。
	p.Province = "江西"
	p.Address.City = "南昌"
	fmt.Printf("p= %#v\n", p) //p= main.Person{Name:"ixfosa", Age:22, Address:main.Address{Province:"江西", City:"南昌"}}
}


/*----------------嵌套结构体的字段名冲突--------------------*/
/*
嵌套结构体内部可能存在相同的字段名。这个时候为了避免歧义需要指定具体的内嵌结构体的字段。
 */
func main1519() {
	type Addr struct {
		Province string
		City string
		CreateTime string
	}

	type Eamil struct {
		Account string
		CreateTime string
	}

	type User struct {
		Name, Gender string
		Addr
		Eamil
	}

	var user User
	user.Name = "FO"
	user.Gender = "man"
	user.Province = "江西"
	user.Addr.City = "南昌"
	user.Addr.CreateTime = "2020" // //指定Addr结构体中的CreateTime

	user.Account = "ixfosa"

	user.Eamil.CreateTime = "2021" //指定Email结构体中的CreateTime

	//user = main.User{Name:"FO", Gender:"man", Addr:main.Addr{Province:"江西", City:"南昌", CreateTime:"2020"}, Eamil:main.Eamil{Account:"ixfosa", CreateTime:"2021"}}
	fmt.Printf("user = %#v", user)
}


/*-----------------结构体的“继承”-------------------*/

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) Run()  {
	fmt.Printf("%s会动！\n", a.name)
}

type Dog struct {
	feet int8
	*Animal
}

func (d *Dog) Sleep()  {
	fmt.Printf("%s睡觉！\n", d.name)
}

func main1520()  {
	dog := &Dog{
		feet: 4,
		Animal: &Animal{
			name: "老黑",
		},
	}
	fmt.Printf("dog = %#v\n", dog) //dog = &main.Dog{feet:4, Animal:(*main.Animal)(0xc0420441d0)}

	dog.Run() //老黑会动！
	dog.Animal.Run() //老黑会动！
	dog.Sleep() //老黑睡觉！

}


/*-----------------结构体字段的可见性-------------------*/
/*
结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
 */


/*------------------- 结构体与JSON序列化-----------------*/
/*
JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式。
JSON键值对是用来保存JS对象的一种方式，键/值对组合中的键名写在前面并用双引号""包裹，使用冒号:分隔，然后紧接着值；多个键值之间使用英文,分隔。
 */

/*type Student struct {
	ID int
	Name, Gender string
}

type Class struct {
	Title string
	Students []*Student
}

func main()  {
	c := Class{
		Title: "18本软件5班",
		Students: make([]*Student, 0, 100),
	}
	for i := 1; i <= 3; i++ {
		stu := &Student{
			ID: i,
			Name: fmt.Sprintf("stu%02d", i),
			Gender: "男",
		}
		c.Students = append(c.Students, stu)
	}

	data, err := json.Marshal(c)

	if err != nil {
		fmt.Println("json.Marshal failed")

	}

	// //JSON序列化：结构体-->JSON格式的字符串
	//json: {"Title":"18本软件5班","Students":[{"ID":1,"Name":"stu01","Gender":"男"},{"ID":2,"Name":"stu02","Gender":"男"},{"ID":3,"Name":"stu03","Gender":"男"}]}
	fmt.Printf("json: %s \n", data)

	// //JSON序列化：结构体-->JSON格式的字符串
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json.Unmarshal failed")
		return
	}
	fmt.Printf("%#v\n", c1)
}
*/

/*--------------- 结构体标签（Tag）---------------------*/
/*
Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。
Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
    `key1:"value1" key2:"value2"`
结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。
键值对之间使用一个空格分隔。 注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。
结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。
	例如不要在key和value之间添加空格。
	例如我们为Student结构体的每个字段定义json序列化时使用的Tag：
 */
type Student struct {
	ID int `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name string //私有不能被json包访问
}

func main1521() {
	stu := &Student{
		ID: 66666,
		Gender: "nan",
		name: "ixfosa",
	}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json.Marshal failed")
		return
	}
	fmt.Printf("data:%s", data) //data:{"id":66666,"Gender":"nan"}
}


/*----------------小练习：--------------------*/
type student struct {
	id int
	name string
	age int
}

func demo(s []student)  {
	//切片是引用传递，是可以改变值的
	s[0].id = 66
}
func main1522() {
	var s []student //定义一个切片类型的结构体
	s = []student{
		{1, "long", 22},
		{1, "zhong", 20},
	}
	//s: []main.student{main.student{id:1, name:"long", age:22}, main.student{id:1, name:"zhong", age:20}}
	fmt.Printf("s: %#v\n", s)
	demo(s)
	//s: []main.student{main.student{id:66, name:"long", age:22}, main.student{id:1, name:"zhong", age:20}}
	fmt.Printf("s: %#v\n", s)
}


/*--------------- 删除map类型的结构体---------------------*/
func main1523() {
	s := make(map[int]student, 5)
	s[0] = student{
		id: 1,
		name: "Long",
		age: 22,
	}

	s[1] = student{
		id: 2,
		name: "Zhong",
		age: 20,
	}

	s[2] = student{
		id: 3,
		name: "ixfosa",
		age: 22,
	}

	fmt.Println(s) //map[0:{1 Long 22} 1:{2 Zhong 20} 2:{3 ixfosa 22}]

	delete(s, 2)

	fmt.Println(s) //map[0:{1 Long 22} 1:{2 Zhong 20}]
}



/*------------------实现map有序输出------------------*/
func main() {
	data := make(map[int]string, 5)
	data[0] = "java"
	data[1]  = "python"
	data[2] = "c/c++"
	data[3] = "goland"
	data[4] = "rust"
	data[5] = "sql"
	fmt.Println(data) //map[3:go 4:rust 5:sql 0:java 1:python 2:c/c++]

	keys := []int{}
	for k := range data {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		println(key, "=>", data[key])
	}
	//0 => java
	//1 => python
	//2 => c/c++
	//3 => go
	//4 => rust
	//5 => sql
}


