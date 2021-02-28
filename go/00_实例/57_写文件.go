package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path := "./defer.txt"
	data := "hello goland..."
	err := ioutil.WriteFile(path, []byte(data), 0644)
	Check(err)

	// 对于更细粒度的写入，先打开一个文件。
	// file, err := os.OpenFile(path, os.O_RDWR, 0666)
	file, err := os.Create("./test")
	Check(err)

	// 打开文件后，习惯立即使用 defer 调用文件的 Close操作。
	defer file.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := file.Write(d2)
	Check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := file.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// 调用 Sync 来将缓冲区的信息写入磁盘。
	file.Sync()

	w := bufio.NewWriter(file)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// 使用 Flush 来确保所有缓存的操作已写入底层写入器。
	w.Flush()
}
