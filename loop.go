package main

import (
	"fmt"
	"time"
)

func main()  {
	i := 0
	ch := make(chan string, 0)

	defer func() {
		close(ch)
	}()

	// 當沒有值送進來，就會停在 select 區塊
	// 結束 for or select 都需要透過 break
	// 但在 select 區間直接結束掉 for 迴圈，只能使用 break variable 來結束
	go func() {
		LOOP:
			for {
				time.Sleep(1 * time.Second)
				fmt.Println(time.Now().Unix())
				i++

				select {
				case m := <-ch:
					println(m)
					break LOOP
				default:
				}
			}
	}()

	time.Sleep(time.Second * 4)
	ch <- "stop"
}
