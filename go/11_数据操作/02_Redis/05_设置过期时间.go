package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println("redis Dial err: ", err)
		return
	}

	defer conn.Close()

	reply, err := conn.Do("Set", "msg", "hello world")
	if err != nil {
		fmt.Println("conn.Do Set err: ", err)
		return
	}
	fmt.Println("reply: ", reply)

	do, err := conn.Do("expire", "msg", 10)
	if err != nil {
		fmt.Println("conn.Do err expire: ", err)
		return
	}
	fmt.Println("do: ", do)
}
