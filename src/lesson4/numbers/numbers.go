package numbers

import (
	"errors"
)

func Increment(initialValue int, incBy int) (int, error) {
	if incBy < 0 {
		return 0, errors.New("incBy need to more than 0")
	}

	var workers = make(chan int, incBy)

	for i := 0; i < incBy; i++ {
		go func(channel chan int) {
			channel <- 1
		}(workers)
	}

	for i := 0; i < incBy; i++ {
		initialValue += <-workers
	}

	return initialValue, nil
}
