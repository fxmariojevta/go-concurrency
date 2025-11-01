package main

import (
	"fmt"
	"time"
)

func printSomething() {
	fmt.Println("Something")
}

func main() {
	go printSomething()

	time.Sleep(1 * time.Second)

	fmt.Println("Something in main")
}
