package main

func main()  {
	foo := addByShareCommunicate(10)
	fmt.Println(len(foo))
	fmt.Println(foo)
}

// 此 channel 確保只能寫入，不能輸出
func addByShareCommunicate(n int) []int {
	var ints []int
	channel := make(chan int, n)

	for i := 0; i < n; i++ {
		go func(channel chan <- int, order int) {
			channel <- order
		}(channel, i)
	}

	for i := range channel{
		ints = append(ints, i)

		if len(ints) == n {
			break
		}
	}

	close(channel)

	return ints
}