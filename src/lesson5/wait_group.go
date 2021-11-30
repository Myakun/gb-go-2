package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
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

	wg := sync.WaitGroup{}
	wg.Add(int(number))

	for i := 0; i < int(number); i += 1 {
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
