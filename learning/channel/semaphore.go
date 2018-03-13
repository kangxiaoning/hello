package main

import (
	"runtime"
	"sync"
	"time"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	sem := make(chan struct{}, 2) // 最多允许2个并发

	for i := 0; i < 10; i++ {
		wg.Add(1)

		// 最多有2个goroutine能获取信号量
		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}        // acquire sem
			defer func() { <-sem }() // release sem

			time.Sleep(time.Second * 2)
			fmt.Println(id, time.Now())
		}(i)
	}

	wg.Wait()

}
