package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter directory for files: ")

	scanner.Scan()
	dir := scanner.Text()

	info, err := os.Stat(dir)
	if nil != err {
		fmt.Printf("directory %s not exists", dir)
		return
	}

	if !info.IsDir() {
		fmt.Printf("directory %s isn't a directory", dir)
		return
	}

	for i := 0; i < 1000000; i++ {
		createFile(dir + "/" + strconv.Itoa(i))
	}
}

func createFile(filename string) {
	created, err := os.Create(filename)
	if nil != err {
		fmt.Println(err)
		return
	}
	defer func(created *os.File) {
		err := created.Close()
		if nil != err {
			fmt.Printf("Can't close file %s\n", filename)
		}
	}(created)
}
