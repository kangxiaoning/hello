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
}
