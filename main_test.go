package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestPrintSomething(t *testing.T) {
	expected := "Something"

	stdOut := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe failed")
	}
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go printSomething(expected, &wg)
	wg.Wait()

	err = w.Close()
	if err != nil {
		t.Errorf("w.close error")
	}

	result, err := io.ReadAll(r)
	if err != nil {
		t.Errorf("io.ReadAll failed")
	}
	output := string(result)
	output = strings.TrimSpace(output)

	os.Stdout = stdOut

	if output != expected {
		t.Errorf("Output: %s but Expected %s", output, expected)
	}
}
