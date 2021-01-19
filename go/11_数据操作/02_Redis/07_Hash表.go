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

	// hmset key field1 value1 [field2 field2......]
	//reply, err := conn.Do("hmset", "role", "name", "ixfosa", "age", "22", "sex", "男")

	// hset key filed value
	_, err = conn.Do("hset", "role", "name", "ixfosa")
	if err != nil {
		fmt.Println("conn.Do hset err: ", err)
		return
	}
	_, err = conn.Do("hset", "role", "age", 22)
	if err != nil {
		fmt.Println("conn.Do hset err: ", err)
		return
	}
	_, err = conn.Do("hset", "role", "sex", "男")
	if err != nil {
		fmt.Println("conn.Do hset err: ", err)
		return
	}

	r, err := redis.String(conn.Do("hget", "role", "sex"))
	if err != nil {
		fmt.Println("redis.String hget err: ", err)
		return
	}
	fmt.Println("r: ", r)
	// redis 中文存储乱码问题
	// redis-cli 后面加上 --raw
}
