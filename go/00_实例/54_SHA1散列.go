package main

import (
	"crypto/sha1"
	"fmt"
)

// SHA1 散列经常用生成二进制文件或者文本块的短标识
// Go 在多个 crypto/* 包中实现了一系列散列函数。
func main() {
	str := "hello go"

	// 产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。
	hash := sha1.New()

	// 写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	hash.Write([]byte(str))

	// 得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bytes := hash.Sum(nil)

	fmt.Println(str) // hello go

	// SHA1 值经常以 16 进制输出
	fmt.Printf("%x\n", bytes) // 4e5ab3a6ecc24cd5b413a6d4db5db02fb4907018

	fmt.Println(len("4e5ab3a6ecc24cd5b413a6d4db5db02fb4907018")) // 40

	// 可以使用和上面相似的方式来计算其他形式的散列值。例如，计算 MD5 散列，引入 crypto/md5 并使用 md5.New()方法。
}
