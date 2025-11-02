package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", 0 * time.Second},
		{"quarter second delay", 250 * time.Millisecond},
		{"half second delay", 500 * time.Millisecond},
	}

	for _, e := range theTests {
		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay

		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slice; expected : 5 but got %d", len(orderFinished))
		}
	}
}
