package coding

import (
	"fmt"
	"sync"
)

func worker(id int, task chan int, wg *sync.WaitGroup) chan int {
	next := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for d := range task {
			if d >= 10 {
				break
			}
			fmt.Printf("goroutine: %d recevie %d\n", id, d)
			next <- d + 1
		}
		close(next)
		fmt.Println("goroutine: finish", id)
	}()
	return next
}

func Print(n int) {
	var wg sync.WaitGroup
	var first = make(chan int)
	last := first
	for i := 0; i < n; i++ {
		last = worker(i, last, &wg)
	}
	// 写入到第一个，开启循环
	first <- 0
	// 读取最后一个goroutine的数据，并写入到第一个, 达到循环的效果
	for v := range last {
		first <- v
	}
	// 关闭第一个
	close(first)
	wg.Wait()
}
