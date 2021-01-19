package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/*
type Pool struct {
	Dial func() (Conn, error)
	TestOnBorrow func(c Conn, t time.Time) error
	MaxIdle int
	IdleTimeout time.Duration
	Wait bool
	MaxConnLifetime time.Duration
	mu     sync.Mutex    // mu protects the following fields
	closed bool          // set to true when the pool is closed.
	active int           // the number of open connections in the pool
	ch     chan struct{} // limits open connections when p.Wait is true
	idle   idleList      // idle connections
}
*/

var pool *redis.Pool


func init()  {
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive: 0, //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）    
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func main() {
	conn := pool.Get() //从连接池，取一个链接
	defer conn.Close()  //函数运行结束 ，把连接放回连接池

	reply, err := conn.Do("set", "msg", "hello world")
	if err != nil {
		fmt.Println(" conn.Do set err: ", err)
		return
	}
	fmt.Println("reply: ", reply)

	r, err := redis.String(conn.Do("get", "msg"))
	if err != nil {
		fmt.Println(" redis.String get err: ", err)
		return
	}
	fmt.Println("r: ", r)
}
