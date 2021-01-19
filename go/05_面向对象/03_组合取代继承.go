package main

import "fmt"

/*****************组合取代继承*********************/
/*
golang 中的继承是通过结构体中的匿名字段来实现
 */

//通过嵌套结构体进行组合
type author  struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title     string
	content   string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.author.fullName())
	fmt.Println("Bio: ", p.author.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main()  {

	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}

	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}

	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}

	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		posts: []post{post1, post2, post3},
	}
	w.contents()
}
func main031()  {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post1.details()
	post1.fullName()
}


