package main

import (
	"fmt"
	"github.com/myakun/numbers"
	numbersV2 "github.com/myakun/numbers/v2"
	"time"
)

func main() {
	// V1
	now := time.Now()
	fib := numbers.Fib(30)
	fmt.Printf("Fibonacci v1 result is %d in %d microsec\n", fib, time.Now().Sub(now).Microseconds())

	// V2
	now = time.Now()
	fib = numbersV2.Fib(30)
	fmt.Printf("Fibonacci v2 result is %d in %d microsec\n", fib, time.Now().Sub(now).Microseconds())
}
