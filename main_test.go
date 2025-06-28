package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("Check main is called", func(t *testing.T) {
		origArgs := os.Args
		defer func() { os.Args = origArgs }()
		os.Args = []string{
			"test",
			"--num1=1",
			"--num2=2",
			"--oper=add",
		}

		// create a pipe to capture stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		main()
		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)

		output := buf.String()
		if !strings.Contains(output, "GO Calculator App") {
			t.Errorf("expected output to contain 'GO Calculator App', got %q", output)
		}
	})

	t.Run("Check main 1+2=3", func(t *testing.T) {
		origArgs := os.Args
		defer func() { os.Args = origArgs }()
		os.Args = []string{
			"test",
			"--num1=1",
			"--num2=2",
			"--oper=add",
		}

		// create a pipe to capture stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		main()
		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)

		output := buf.String()
		if !strings.Contains(output, "Output: 3") {
			t.Errorf("expected output to contain 'Output: 3', got %q", output)
		}
	})
	t.Run("Check main 1-2=-1", func(t *testing.T) {
		origArgs := os.Args
		defer func() { os.Args = origArgs }()
		os.Args = []string{
			"test",
			"--num1=1",
			"--num2=2",
			"--oper=sub",
		}

		// create a pipe to capture stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		main()
		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)

		output := buf.String()
		if !strings.Contains(output, "Output: -1") {
			t.Errorf("expected output to contain 'Output: -1', got %q", output)
		}
	})

	t.Run("Check main 2*2=4", func(t *testing.T) {
		origArgs := os.Args
		defer func() { os.Args = origArgs }()
		os.Args = []string{
			"test",
			"--num1=2",
			"--num2=2",
			"--oper=mult",
		}

		// create a pipe to capture stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		main()
		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)

		output := buf.String()
		if !strings.Contains(output, "Output: 4") {
			t.Errorf("expected output to contain 'Output: 4', got %q", output)
		}
	})

	t.Run("Check main 4/2=2", func(t *testing.T) {
		origArgs := os.Args
		defer func() { os.Args = origArgs }()
		os.Args = []string{
			"test",
			"--num1=4",
			"--num2=2",
			"--oper=div",
		}

		// create a pipe to capture stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		main()
		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)

		output := buf.String()
		if !strings.Contains(output, "Output: 2") {
			t.Errorf("expected output to contain 'Output: 2', got %q", output)
		}
	})
}

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
	t.Run("Add: 1 + 1", func(t *testing.T) {
		got := addInputs(1, 1)
		want := 2
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSub(t *testing.T) {
	t.Run("Sub: 1 - 1", func(t *testing.T) {
		got := subInputs(1, 1)
		want := 0
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestMult(t *testing.T) {
	t.Run("Mult: 1 * 1", func(t *testing.T) {
		got := multInputs(1, 1)
		want := 1
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestDiv(t *testing.T) {
	t.Run("Div: 1 / 1", func(t *testing.T) {
		got, _ := divInputs(1, 1)
		want := 1
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Div: 1 / 0", func(t *testing.T) {
		_, got := divInputs(1, 0)
		want := fmt.Errorf(ErrorDivideByZero)
		if got.Error() != want.Error() {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
