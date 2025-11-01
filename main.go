package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	msg = "Hello, world"

	words := []string{
		"universe",
		"cosmos",
		"world!",
	}

	for _, x := range words {
		wg.Add(1)
		go updateMessage(x)
		wg.Wait()
		printMessage()
	}
}
