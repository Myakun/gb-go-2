package main

import (
	"lesson6/set"
	"sync"
)

func main() {
	number := 100
	wg := sync.WaitGroup{}
	wg.Add(number * 2)

	for i := 0; i < number; i++ {
		for j := 0; j < 2; j++ {
			go func(n int) {
				set.WriteData(n, float64(n))
				set.ReadData(n)
				wg.Done()
			}(i)
		}
	}

	wg.Wait()
}
