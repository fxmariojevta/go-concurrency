package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	p := PizzaOrder{
		pizzaNumber: pizzaNumber,
	}

	if pizzaNumber > NumberOfPizzas {
		return &p
	}

	fmt.Printf("Received order #%d!\n", pizzaNumber)
	localRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	delay := localRand.Intn(5) + 1
	fmt.Printf("Making pizza #%d. It will take %d seconds....\n", pizzaNumber, delay)

	time.Sleep(time.Duration(delay) * time.Second)

	rnd := localRand.Intn(12) + 1
	var msg string
	var success bool

	if rnd <= 2 {
		msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		pizzasFailed++
	} else if rnd <= 4 {
		msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		pizzasFailed++
	} else {
		msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		pizzasMade++
		success = true
	}
	total++

	p.message = msg
	p.success = success

	return &p
}

func pizzeria(pizzaMaker *Producer) {
	i := 0

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("Order #%d could not be completed", i.pizzaNumber)
			}
		} else {
			color.Cyan("Done making pizzas....")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("Error closing the pizza job", err)
			}
		}
	}

	// print out the ending message
	color.Cyan("----------------------------------")
	color.Cyan("The Pizzeria is closing for the day!")

	color.Cyan("Pizzas made: %d", pizzasMade)
	color.Cyan("Pizzas failed: %d", pizzasFailed)
	color.Cyan("Total orders: %d", total)

	switch {
	case pizzasFailed > 9:
		color.Red("The pizzeria is going out of business!")
	case pizzasFailed >= 6:
		color.Red("The pizzeria is in trouble!")
	case pizzasFailed >= 4:
		color.Yellow("The pizzeria is doing okay.")
	case pizzasFailed >= 2:
		color.Green("The pizzeria is doing good!")
	default:
		color.Green("The pizzeria is doing great!")
	}
}
