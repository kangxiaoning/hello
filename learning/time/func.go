package main

import (
	"time"
	"fmt"
)

func main() {
	// convert time to millisecond - with UnixNano()
	end := time.Now().UnixNano() / int64(time.Millisecond)
	start := time.Now().Add(time.Minute * -1).UnixNano() / int64(time.Millisecond)

	fmt.Println("start:", start, "end:", end)

	startTime := time.Unix(0, start*int64(time.Millisecond))
	endTime := time.Unix(0, end*int64(time.Millisecond))
	fmt.Println("startTime:", startTime, "endTime:", endTime)

	// convert millisecond to time - with time.Unix
	fmt.Println("start:", time.Unix(0,
		1521528912389*int64(time.Millisecond)).Format(
		"2006-01-02 15:04:05"))
	fmt.Println(time.Unix(0,
		1521528900000*int64(time.Millisecond)).Format(
		"2006-01-02 15:04:05"))
	fmt.Println(time.Unix(0,
		1521528960000*int64(time.Millisecond)))
	fmt.Println("end:", time.Unix(0,
		1521528972389*int64(time.Millisecond)).Format(
		"2006-01-02 15:04:05"))

	// 经验证到达60秒时，如下机制依然得到正确的时间
	tick := time.Tick(time.Minute)
	for {
		n := time.Now()
		sstart := time.Date(n.Year(),
			n.Month(),
			n.Day(),
			n.Hour(),
			n.Minute(),
			n.Second()-2,
			0,
			n.Location())
		send := time.Date(n.Year(),
			n.Month(),
			n.Day(),
			n.Hour(),
			n.Minute(),
			n.Second()-1,
			0,
			n.Location())

		time.Sleep(time.Second)

		select {
		case <-tick:
			return
		case <-time.After(time.Second):
			fmt.Println(sstart, send)
		}
	}

}
