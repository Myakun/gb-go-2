package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number: ")
	scanner.Scan()
	number, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if nil != err {
		fmt.Println("Error:", err)
		return
	}

	var (
		counter int
		mutex   sync.Mutex
		ch      = make(chan struct{}, number)
	)

	for i := 0; i < int(number); i += 1 {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			counter += 1
			ch <- struct{}{}
		}()
	}

	time.Sleep(2 * time.Second)
	close(ch)

	i := 0
	for range ch {
		i += 1
	}
}
