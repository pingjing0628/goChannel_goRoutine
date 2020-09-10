package main

import "fmt"

// select + default 確保 channel 是否已滿
func main()  {
	ch := make(chan int, 2)

	ch <- 1

	select {
	case ch <- 2:
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)
	default:
		fmt.Println("channel blocking")
	}
}
// buffer size 改為 2，就可以繼續塞 value 進 channel