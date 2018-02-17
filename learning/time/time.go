package main

import (
	"fmt"
	"time"
)

func main() {
	//初始化定时器
	t := time.NewTimer(2 * time.Second)
	//当前时间
	now := time.Now()
	fmt.Printf("Now time : %v.\n", now)

	// C是一个chan time.Time类型的缓冲通道，一旦timer到期
	// 定时器就会向自己的C字段发送一个time.Time类型的元素值
	expire := <-t.C
	fmt.Printf("Expiration time: %v.\n", expire)

	// time.After
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	// 注释下面一行则等待2秒打印 "Timed out"
	ch1 <- 1

	select {
	case e1 := <-ch1:
		//如果ch1通道成功读取数据，则执行该case处理语句
		fmt.Printf("1th case is selected. e1=%v\n", e1)
	case e2 := <-ch2:
		//如果ch2通道成功读取数据，则执行该case处理语句
		fmt.Printf("2th case is selected. e2=%v\n", e2)
	case <-time.After(2 * time.Second):
		fmt.Println("Timed out")
	}

	//
	var tt *time.Timer

	f := func() {
		fmt.Printf("Expiration time : %v.\n", time.Now())
		fmt.Printf("C`s len: %d\n", len(tt.C))
	}

	tt = time.AfterFunc(1*time.Second, f)
	// 让当前Goroutine 睡眠2s，确保大于内容的完整
	// 这样做原因是，time.AfterFunc的调用不会被阻塞
	// 它会以异步的方式在到期事件来临执行我们自定义函数f。
	time.Sleep(2 * time.Second)
}
