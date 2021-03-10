package main

// A、B、C、D 哪些选项有语法错误？
type S struct {
	
}

func f421(f interface{})  {

}

func g421(f *interface{})  {

}

func main() {
	s := S{}
	p := &s

	f421(s) // A
	// g421(s) // B
	f421(p) // C
	// g421(p) // D
}
// BD。函数参数为 interface{} 时可以接收任何类型的参数，包括用户自定义类型等，
// 即使是接收指针类型也用 interface{}，而不是使用 *interface{}。
