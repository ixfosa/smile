package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

/****************** web工作流程*********************/
/*
Web服务器的工作原理可以简单地归纳为
	客户机通过TCP/IP协议建立到服务器的TCP连接
	客户端向服务器发送HTTP协议请求包，请求服务器里的资源文档
	服务器向客户机发送HTTP协议应答包，如果请求的资源包含有动态语言的内容，那么服务器会调用动态语言的解释引擎负责处理“动态内容”，
		并将处理得到的数据返回给客户端
	客户机与服务器断开。由客户端解释HTML文档，在客户端屏幕上渲染图形结果
 */

/********************HTTP协议*******************/
/*
超文本传输协议(HTTP，HyperText Transfer Protocol)是互联网上应用最为广泛的一种网络协议，
它详细规定了浏览器和万维网服务器之间互相通信的规则，通过因特网传送万维网文档的数据传送协议

HTTP协议通常承载于TCP协议之上
 */
func myHandler(resp http.ResponseWriter, req *http.Request)  {

	fmt.Println(req.RemoteAddr, "连接成功")

	fmt.Printf("%#v", resp)

	resp.Write([]byte(time.Now().String()))

}

func main031() {

	//http://127.0.0.1:8000/go
	http.HandleFunc("/go", myHandler)

	// addr：监听的地址
	// handler：回调函数
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func main() {
	resp, err := http.Get("http://www.topgoer.com")
	if err !=nil {
		fmt.Println("http.Get err:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("%v", resp)

	for {
		buf := [4096]byte{}
		n, err :=resp.Body.Read(buf[:])

		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			res := string(buf[:n])
			fmt.Print(res)
			break
		}
	}
}