package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func main() {
	// func ReadFile(filename string) ([]byte, error)
	data, err := ioutil.ReadFile("./defer.txt")
	check(err)
	fmt.Println(string(data)) // hello go...

	file, err := os.Open("./defer.txt")
	check(err)

	buf := make([]byte, 5)
	// func (f *File) Read(b []byte) (n int, err error)
	n, err := file.Read(buf)
	check(err)
	fmt.Println(n, " ", string(buf[:n])) // 5   hello

	//  Seek 到一个文件中已知的位置并从这个位置开始进行读取。
	ret2, err := file.Seek(6, 0)
	check(err)
	buf2 := make([]byte, 2)
	n2, err := file.Read(buf2)
	check(err)
	fmt.Println("ret2:", ret2, " n2:", n2, " ", string(buf2[:n2])) // ret2: 6  n2: 3   go


	ret3, err := file.Seek(6, 0)
	check(err)
	buf3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(file, buf3, 2)
	check(err)
	fmt.Println("ret3:", ret3, " n3:", n3, " ", string(buf3[:n3])) // ret3: 6  n3: 2   go

	// 没有内置的回转支持，但是使用 Seek(0, 0) 实现。
	_, err = file.Seek(0, 0)
	check(err)


	// bufio 包实现了带缓冲的读取，这不仅对有很多小的读取操作的能提升性能，也提供了很多附加的读取函数。
	reader := bufio.NewReader(file)
	bytes, err := reader.Peek(5)
	check(err)
	fmt.Println(string(bytes)) // hello

	file.Close()
}
