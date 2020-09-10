package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() { // calculate goroutine
		fmt.Println("Starts calculating")
		time.Sleep(time.Second)
		fmt.Println("Ends calculating")

		ch <- "END" // goroutine will be forced to wait

		fmt.Println("calculate goroutine finished")
	}()

	time.Sleep(2 * time.Second) // make main is slower than goroutine
	fmt.Println(<-ch)
	time.Sleep(time.Second)
	fmt.Println("main goroutine finished")
}
