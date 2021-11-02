package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type CustomError struct {
	err  error
	time time.Time
}

func (err *CustomError) Error() string {
	return fmt.Sprintf("critical error at %s: %s", err.time, err.err)
}

func newCustomError(err error) *CustomError {
	return &CustomError{
		err:  err,
		time: time.Now(),
	}
}

const (
	errMsg      = "We could not accept the order"
	tryAgainMsg = "Try again or call waiter"
)

var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	defer func() {
		if v := recover(); v != nil {
			printErrMsg(newCustomError(errors.New(fmt.Sprint(v))))
			orderFood()
		}
	}()

	orderFood()
}

func orderFood() {
	menu := []string{
		"Beef",
		"Chicken",
		"Fish",
	}

	fmt.Println("Choose a dish from the menu")
	for i, name := range menu {
		fmt.Printf("%d - %s\n", i+1, name)
	}

	fmt.Print("Enter dish number: ")

	scanner.Scan()
	dishNumber, err := strconv.ParseInt(scanner.Text(), 10, 0)
	if err != nil {
		printErrMsg(err)
		orderFood()
		return
	}

	fmt.Printf("You ordered %s\n", strings.ToLower(menu[dishNumber-1]))
}

func printErrMsg(msg interface{}) {
	fmt.Printf("%s: %s\n", errMsg, msg)
	fmt.Println(tryAgainMsg)
}
