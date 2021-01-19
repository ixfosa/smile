package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost: 6379")

	if err != nil {
		fmt.Println("redis.Dial err: ", err)
	}

	fmt.Println("redis conn success")

	defer conn.Close()
}
