package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339)) // 2021-02-21T15:02:47+08:00

	str := "1998-02-19T00:00:00+08:00"
	t1, _ := time.Parse(time.RFC3339, str)
	fmt.Println(t1) // 1998-02-19 00:00:00 +0800 CST
	fmt.Println(t1.Format("2006-01-02 15:04")) // 1998-02-19 00:00
	fmt.Printf("%T \n", t1) // time.Time

	fmt.Println(now.Format("3:04PM"))
	fmt.Println(now.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(now.Format("2006-01-02 15:04")) // 2021-02-21 15:14
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	fmt.Println(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	fmt.Println(e)
}
