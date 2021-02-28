package main

import (
	"fmt"
	"os"
)

// Defer 被用来确保一个函数调用在程序执行结束前执行。
// 同样用来执行一些清理工作。 defer 用在像其他语言中的ensure 和 finally用到的地方。
func main() {
	file := createFile("./defer.txt")
	defer closeFile(file)
	writeFile(file)
}

func createFile(path string) *os.File {
	fmt.Println("creating")
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func writeFile(file *os.File) {
	fmt.Println("writing")
	fmt.Fprint(file, "hello go...")
}


func closeFile(file *os.File) {
	fmt.Println("closing")
	file.Close()
}

