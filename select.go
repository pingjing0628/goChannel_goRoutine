package main

import "fmt"

// select 只能接 channel ，否則會出錯
// default 會直接執行，所以沒有 default 的 select會遇到 blocking
func main()  {
	ch := make(chan int, 1)

	// 若是沒有送 value 進來，channel 就會 panic
	//ch <- 1 // 移除就會造成 deadlock，主程式爆炸

	// 有時會拿到01有時拿到02
	// case 是隨機選取
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	default:
		fmt.Println("exit") // 解決 deadlock
	}
}