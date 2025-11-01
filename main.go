package main

import "fmt"

func printSomething() {
	fmt.Println("Something")
}

func main() {
	go printSomething()

	fmt.Println("Something in main")
}
