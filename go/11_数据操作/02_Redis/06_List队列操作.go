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

	reply, err := conn.Do("lpush", "list", "node3", "node2", "node1")
	if err != nil {
		fmt.Println("conn.Do lpush err: ", err)
		return
	}
	fmt.Println("reply: ", reply)

	r, err := redis.String(conn.Do("lpop", "list"))
	if err != nil {
		fmt.Println("redis.String lpop err: ", err)
		return
	}
	fmt.Println("r: ", r)

	res, err := redis.String(conn.Do("lindex", "list", 0))
	if err != nil {
		fmt.Println("redis.String lindex err: ", err)
		return
	}
	fmt.Println("res: ", res)
}

/*
127.0.0.1:6379> lindex list 0
"node1"
127.0.0.1:6379> lindex list 1
"node2"
127.0.0.1:6379> lindex list 2


*/
