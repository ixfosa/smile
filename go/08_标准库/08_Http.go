package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*****************HTTP客户端**********************/
/*
基本的HTTP/HTTPS请求 Get、Head、Post和PostForm函数发出HTTP/HTTPS请求。
	resp, err := http.Get("http://5lmh.com/")
	...
	resp, err := http.Post("http://5lmh.com/upload", "image/jpeg", &buf)
	...
	resp, err := http.PostForm("http://5lmh.com/form", url.Values{"key": {"Value"}, "id": {"123"}})


程序在使用完response后必须关闭回复的主体。
	resp, err := http.Get("http://5lmh.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
 */


/****************GET请求示例***********************/

func main081() {
	resp, err := http.Get("https://www.jxut.edu.cn/")
	if err != nil {
		fmt.Println("http.Get err:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err:", err)
		return
	}
	fmt.Println(string(body))
}

/******************带参数的GET请求示例*********************/
/*
	GET请求的参数需要使用Go语言内置的net/url这个标准库来处理。

func main() {
    apiUrl := "http://127.0.0.1:9090/get"
    // URL param
    data := url.Values{}
    data.Set("name", "枯藤")
    data.Set("age", "18")
    u, err := url.ParseRequestURI(apiUrl)
    if err != nil {
        fmt.Printf("parse url requestUrl failed,err:%v\n", err)
    }
    u.RawQuery = data.Encode() // URL encode
    fmt.Println(u.String())
    resp, err := http.Get(u.String())
    if err != nil {
        fmt.Println("post failed, err:%v\n", err)
        return
    }
    defer resp.Body.Close()
    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("get resp failed,err:%v\n", err)
        return
    }
    fmt.Println(string(b))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    data := r.URL.Query()
    fmt.Println(data.Get("name"))
    fmt.Println(data.Get("age"))
    answer := `{"status": "ok"}`
    w.Write([]byte(answer))
}
 */
func main082()  {
	apiUrl := "https://tieba.baidu.com/f"
	data := url.Values{}
	data.Set("kw", "go语言")
	data.Set("ie", "utf-8")
	data.Set("pn", "50")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println("url.ParseRequestURI:", err)
		return
	}

	u.RawQuery = data.Encode()

	fmt.Println(u.String())

	resp, err := http.Get(u.String())
	//resp, err := http.Get("https://tieba.baidu.com/f?kw=go%E8%AF%AD%E8%A8%80&ie=utf-8&pn=50")
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("ioutil.ReadAll err:", err)
		return
	}
	fmt.Println(string(body))
}

/******************Post请求示例*********************/
func main083()  {
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=ixfosa&age=22"
	// json
	contentType := "application/json"
	data := `{"name":"ixfosa", "age":"22"}`

	// func Post(url string, contentType string, body io.Reader) (resp *Response, err error)
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("http.Post err:", err)
		return
	}
	defer resp.Body.Close()

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}


/******************自定义Client*********************/
/*
管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client：
	client := &http.Client{
    CheckRedirect: redirectPolicyFunc,
	}
	resp, err := client.Get("http://5lmh.com")
	// ...
	req, err := http.NewRequest("GET", "http://5lmh.com", nil)
	// ...
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	// ...
 */



/******************自定义Transport*********************/
/*
要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport：
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://5lmh.com")

Client和Transport类型都可以安全的被多个go程同时使用。出于效率考虑，应该一次建立、尽量重用。
 */

/******************默认的Server*********************/
/*
ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，这表示采用包变量DefaultServeMux作为处理器。

Handle和HandleFunc函数可以向DefaultServeMux添加处理器。
	http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
 */

//默认的Server示例
//使用Go语言中的net/http包来编写一个简单的接收HTTP请求的Server端示例，net/http包是对net包的进一步封装，专门用来处理HTTP协议的数据。
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello ixfosa！")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}

/*
自定义Server
要管理服务端的行为，可以创建一个自定义的Server：
	s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
*/