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
	reply, err := conn.Do("MSet", "name1", "ixfosa", "name2", "long", "name3", "zhong")
	if err != nil {
		fmt.Println("conn.Do err: ", err)
		return
	}
	fmt.Println("reply: ", reply)

	strings, err := redis.Strings(conn.Do("MGet", "name1", "name2", "name3"))

	fmt.Println(strings)
}
