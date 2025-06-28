package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFlags(t *testing.T) {
	t.Run("Check for no --oper flag", func(t *testing.T) {
		origArgs := os.Args
		defer func() { os.Args = origArgs }()
		os.Args = []string{
			"test",
			"--num1=1",
			"--num2=2",
			"--oper=",
		}
		_, _, _, err := parseInputs()
		got := err
		want := fmt.Errorf("Error: --oper input cannot be empty")
		if got.Error() != want.Error() {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Check for invalid --oper flag", func(t *testing.T) {
		origArgs := os.Args
		defer func() { os.Args = origArgs }()
		os.Args = []string{
			"test",
			"--num1=1",
			"--num2=2",
			"--oper=bad",
		}
		_, _, _, err := parseInputs()
		got := err
		want := fmt.Errorf("Error: --oper is not add, sub, mult, div")
		if err == nil {
			t.Fatalf("expected error but got nil")
		}

		if got.Error() != want.Error() {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add 1 + 1", func(t *testing.T) {
		got := addInputs(1, 1)
		want := 2
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
