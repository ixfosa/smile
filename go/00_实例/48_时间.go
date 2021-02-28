package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now: ", now) // now:  2021-02-21 13:15:13.1989872 +0800 CST m=+0.002998901

		then := time.Date(1998, 2, 19, 0, 0, 0, 0, time.Local)
		fmt.Println("then: ", then) // then:  1998-02-19 00:00:00 +0800 CST

		fmt.Println("year: ", then.Year()) // year:  1998
		fmt.Println("month: ", then.Month()) // month:  February
		fmt.Println("day: ", then.Day()) // day:  19
		fmt.Println("hour: ", then.Hour()) // hour:  0
		fmt.Println("minute: ", then.Minute()) // minute:  0
		fmt.Println("second:", then.Second()) // second: 0
		fmt.Println(then.Location()) // Local

		fmt.Println("weekday: ", then.Weekday()) // weekday:  Thursday

		fmt.Println(then.Before(now)) //true
		fmt.Println(then.After(now)) // false
		fmt.Println(then.Equal(now)) // false

		diff := now.Sub(then)
		fmt.Println("diff: ", diff)// diff:  201685h23m14.9694138s

		fmt.Println(diff.Hours() / 24 / 365) // 23.023449866442895

		fmt.Println(then.Add(diff)) // 2021-02-21 13:25:47.4360034 +0800 CST
		fmt.Println(now.Add(-diff)) // 1998-02-19 00:00:00 +0800 CST m=-726067589.622945699
}
