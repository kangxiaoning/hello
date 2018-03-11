package main

import (
	"fmt"
	"math/rand"
	"time"
)

func onceSelect() {
	// 从nil的channel拿不到数据，select会进入default分支
	// 非阻塞方式 select + default
	var c1, c2 chan int
	select {
	case n := <-c1:
		fmt.Println("Received from c1", n)
	case n := <-c2:
		fmt.Println("Received from c2", n)
		// 没有default会导致deadlock
	default:
		fmt.Println("No value received")
	}
}

func forLoopSelect() {
	// 循环执行select
	for {
		var c1, c2 chan int
		select {
		case n := <-c1:
			fmt.Println("Received from c1", n)
		case n := <-c2:
			fmt.Println("Received from c2", n)
		// 超时退出
		case <-time.After(time.Microsecond):
			fmt.Println("forLoopSlect timeout,exit")
			return
		default:
			fmt.Println("No value received")
		}
	}
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func doTasks() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int
	// done and tick写在for外层
	done := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeValue = values[0]
			activeWorker = worker
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		// 超时执行任务
		case <-time.After(500 * time.Millisecond):
			fmt.Println("done")
		// 定期执行任务
		case <-tick:
			fmt.Println(time.Now(), "queue len =", len(values))
		// 运行指定时间结束
		case <-done:
			fmt.Println("bye")
			return
		}
	}
}

func main() {
	onceSelect()
	forLoopSelect()
	doTasks()
}
