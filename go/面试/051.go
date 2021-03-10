package main

// 下面这段代码输出什么？为什么？
	// A. speak
	// B. compilation error
/*
type People051 interface {
	Speak(string) string
}

type Student051 struct {

}

func (stu *Student051) Speak(text string) (talk string) {
	if text == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People051 = Student051{}
	text := "speak"
	fmt.Println(peo.Speak(text))
}
 */
// cannot use Student051 literal (type Student051) as type People051 in assignment:
// Student051 does not implement People051 (Speak method has pointer receiver)

// 值类型 Student 没有实现接口的 Speak() 方法，而是指针类型 *Student 实现该方法。