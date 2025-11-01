package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int
	var muBalance sync.Mutex

	fmt.Printf("Initial bank balance: $%d.00\n", bankBalance)

	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gift", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investment", Amount: 100},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				muBalance.Lock()
				bankBalance += income.Amount
				muBalance.Unlock()
				// In windows, executing below line will ended up with infinite loop when running test
				// fmt.Printf("On week %d, you have earned $%d.00 form %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	fmt.Printf("Final bank balance: $%d.00", bankBalance)
}
