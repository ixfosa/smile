package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis.Dial err: ", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("set", "name", "ixfosa")
	//	_, err = conn.Do("set", "age", 22)
	if err != nil {
		fmt.Println("conn.Do err: ", err)
		return
	}

	r, err := redis.String(conn.Do("get", "name"))
	//r, err := redis.Int(conn.Do("get", "age"))
	if err != nil {
		fmt.Println("redis.Int err: ", err)
		return
	}
	fmt.Println(r)
}
