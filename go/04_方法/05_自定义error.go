package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"time"
)

/******************系统抛*********************/
func test051()  {
	a := [5]int{0, 1, 2, 3, 4}
	a[1] = 123
	fmt.Println(a)
	
	//a[10] = 10
	idx := 10
	a[idx] = 100
	fmt.Println(a)
}

func getCircleArea(radius float32) (area float32){
	if radius < 0 {
		panic("半径不能为负")
	}

	return math.Pi * radius * radius
}

func test052()  {
	getCircleArea(-5)
}

func test053()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getCircleArea(-5)
	fmt.Println("这里有没有执行")
}

func test054()  {
	test053()
	fmt.Println("test054")
}
func main051() {
	test054()
	//半径不能为负
	//test054
}


/****************** 返回异常*********************/
func zero(x, y int) (z int, err error) {
	if y == 0 {
		err = errors.New("除数不能为零")
		return
	}
	z = x / y
	return z, nil
}
func main052() {
	v, err := zero(9, 0)
	if err != nil {
		fmt.Println(err) //除数不能为零
	} else {
		fmt.Println(v)
	}

}


/******************自定义error：*********************/
type PathError struct {
	path string
	op string
	createTime string
	message string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path, p.op, p.createTime, p.message)
}

func Open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path: filename,
			op: "read",
			message: err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}
	
	defer file.Close()
	return nil
}

func main()  {
	err := Open("hello.txt")

	switch v := err.(type) {
		case *PathError:
			fmt.Println("get path error,", v)
		default:
	}
	//get path error, path=hello.txt
	//op=read
	//createTime=2020-11-01 17:28:48.7086171 +0800 CST m=+0.008021001
	//message=open hello.txt: The system cannot find the file specified.

}
