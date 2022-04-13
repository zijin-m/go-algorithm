package coding

import (
	"fmt"
	"sync"
)

// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母
// 最终效果如下：12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func PrintNumLetter() {
	var (
		num    = make(chan struct{})
		letter = make(chan struct{})
		wg     = sync.WaitGroup{}
	)
	wg.Add(2)
	go func() {
		i := 1
		for {
			if i > 28 {
				wg.Done()
				return
			}
			<-num
			fmt.Printf("%d%d", i, i+1)
			i = i + 2
			letter <- struct{}{}
		}
	}()

	go func() {
		c := 'A'
		for {
			<-letter
			if c >= 'Z' {
				wg.Done()
				return
			}
			fmt.Printf("%s%s", string(c), string(c+1))
			c = c + 2
			num <- struct{}{}
		}
	}()
	num <- struct{}{}
	wg.Wait()
}
