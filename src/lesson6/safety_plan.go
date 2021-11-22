package main

import (
	"lesson6/set"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	err := trace.Start(os.Stderr)
	if nil != err {
		return
	}
	defer trace.Stop()

	number := 100
	wg := sync.WaitGroup{}
	wg.Add(number)

	for i := 0; i < number; i++ {
		go func(n int) {
			set.WriteData(n, float64(n))
			wg.Done()
		}(i)
	}

	wg.Wait()
}
