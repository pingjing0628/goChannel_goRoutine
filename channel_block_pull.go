package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan string)

	go func() {
		fmt.Println("Starts calculating")
		time.Sleep(time.Second)
		fmt.Println("Ends calculating")

		ch <- "ENDS"

		fmt.Println("calculate goroutine finished")
	}()

	fmt.Println("main goroutine is waiting for channel to receive value")
	fmt.Println(<-ch)
	fmt.Println("main goroutine finished")
}
