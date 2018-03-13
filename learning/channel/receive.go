package main

import (
	"fmt"
)

func main() {

	// 方式一： ok-idom
	fmt.Println("方式一： ok-idom - ok 判断channel是否close")
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done) // 确保发出结束通知

		for {
			x, ok := <-c
			// 判断channel是否close
			if !ok {
				return
			}

			fmt.Println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3

	close(c)

	<-done // 等待receive结束

	// 方式二： range
	fmt.Println("方式二： range - 循环获取消息，知道通道被关闭")
	done1 := make(chan struct{})
	c1 := make(chan int)

	go func() {
		defer close(done1)

		// 循环获取消息，知道通道被关闭
		for x := range c1 {
			fmt.Println(x)
		}
	}()

	c1 <- 4
	c1 <- 5
	c1 <- 6
	close(c1)

	<-done1
}
