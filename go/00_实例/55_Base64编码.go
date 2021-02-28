package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 将要编解码的字符串
	data := "hello go"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc) // aGVsbG8gZ28=
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec)) // hello go

	fmt.Println()

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc) // aGVsbG8gZ28=
	uDec, _ := base64.URLEncoding.DecodeString(sEnc)
	fmt.Println(string(uDec)) // hello go
}
