package main

import "fmt"

// Unbuffered Channel 同步
//func main()  {
//
//	// 建立 channel
//	c := make(chan bool)
//	// go routine
//	go func() {
//		fmt.Println("GO GO GO")
//
//		// 於 go routine 裡面寫入一個值，代表結束
//		c <- true
//	}()
//
//	// 箭頭指向左邊代表 讀出
//	// 代表需要等待讀出一個 channel 值，main 函數才會結束，此時就達到跟用 Sleep 一樣的效果
//	<- c
//	// 裡面寫的 true，外面讀到了，才會結束，才有同步的效果
//}

// Buffered Channel 異步，寫進去不用去等待他讀出來，沒有輸出東西
func main()  {

	// 建立 channel，1代表這個 channel 只有一個
	c := make(chan bool, 1)
	// go routine
	go func() {
		fmt.Println("GO GO GO")
		<- c
	}()
	c <- true
}
