package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()

	fmt.Println("now: ", now) // now:  2021-02-21 13:33:04.6666417 +0800 CST m=+0.003988201

	millis := nanos / 1000000
	fmt.Println(secs) // 1613885584
	fmt.Println(millis) // 1613885584666
	fmt.Println(nanos) // 1613885584666641700

	fmt.Println(time.Unix(secs, 0)) // 2021-02-21 13:33:04 +0800 CST
	fmt.Println(time.Unix(0, nanos)) // 2021-02-21 13:33:04.6666417 +0800 CST

}
