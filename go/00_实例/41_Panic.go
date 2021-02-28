package main

import "fmt"

func java()  {
	fmt.Println("Java...")
}

func python()  {
	panic("python error...")
	fmt.Println("Python...")
}
func main() {
	java()
	python()
}
// Java...
// panic: python error...
