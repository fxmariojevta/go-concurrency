package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	expected := "something"
	wg.Add(1)
	go updateMessage(expected)
	wg.Wait()

	if msg != expected {
		t.Errorf("output: %s, expected: %s", msg, expected)
	}
}

func Test_printMessage(t *testing.T) {
	expected := "Something"

	stdOut := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe failed")
	}
	os.Stdout = w

	wg.Add(1)
	go updateMessage(expected)
	wg.Wait()
	printMessage()

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
		t.Errorf("output: %s but expected %s", output, expected)
	}
}
