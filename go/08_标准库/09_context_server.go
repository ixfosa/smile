package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

/******************客户端超时取消示例*********************/
/*
	调用服务端API时如何在客户端实现超时控制？
 */

func indexHandler(w http.ResponseWriter, r *http.Request)	  {
	number :=  rand.Intn(2)
	if number == 0 {
		time.Sleep(10 * time.Second)
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprintf(w, "quick response")
}
func main() {
	http.HandleFunc("/", indexHandler)
	err:= http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}

//new requestg failed, err:<nil>
//call server api success
//resp:quick response
