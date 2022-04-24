package coding

import (
	"fmt"
	"sync"
)

// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母
// 最终效果如下：12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func PrintNumLetter() {
	num, letter := make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		i := 1
		for range num {
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			letter <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		i := 'A'
		for range letter {
			if i > 'Z' {
				close(num)
				break
			}
			fmt.Print(string(i))
			i++
			fmt.Print(string(i))
			i++
			num <- struct{}{}
		}
	}()

	num <- struct{}{}
	wg.Wait()

}
