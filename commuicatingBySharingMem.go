package main

import (
	"fmt"
	"sync"
)

// Do not communicate by sharing memory; instead, share memory by communicating.

// Write(chan<- int)
// Read(<-chan int)
// ReadWrite(chan int)

// Results: 每次結果都拿到不同的值，因為在 func 內宣告 ints 是 []int，在 goroutine 內是共享這變數，但有可能在同一時間點對同位址 memory 進行讀寫
// 用 goroutine 讀寫變數，盡量不要用 share memory 來共享，出錯就蠻難 debug
// 於是使用 sync 解決
// 因為是放到 goroutine 方式來平行處理，所以需要使用 waitGroup 確保全部 goroutine 拿到資料後，才結束 func

func main()  {
	foo := addByShareMemory(10)
	fmt.Println(len(foo))
	fmt.Println(foo)
}

func addByShareMemory(n int) []int {
	var ints []int
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(n)
	for i := 0;i < n; i++ {
		go func(i int) {
			defer wg.Done()
			mux.Lock()
			ints = append(ints, i)
			mux.Unlock()
		}(i)
	}

	wg.Wait()
	return ints
}

