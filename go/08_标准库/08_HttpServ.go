package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()

	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm)  // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err:", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status":"OK"}`
	w.Write([]byte(answer))
}
func main() {
	http.HandleFunc("/post", postHandler)

	err := http.ListenAndServe("127.0.0.1:9090", nil)

	if err != nil {
		fmt.Println("http.ListenAndServe err:", err)
		return
	}
}
