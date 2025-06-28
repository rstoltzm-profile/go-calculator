package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("GO Calculator App")
	// Parse Inputs
	a, b, o, err := parseInputs()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Oper:   ", o)
	fmt.Println("Inputs: ", a, b)

	// Operate and get output
	output, err := operInputs(a, b, o)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Output:", output)
}

func parseInputs() (int, int, string, error) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	// Get all the flags and parse
	var validOper = [4]string{"add", "sub", "mult", "div"}
	var a = fs.Int("num1", 1, "first number")
	var b = fs.Int("num2", 1, "second number")
	var o = fs.String("oper", "", "operation: add, sub, mult, div")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return 0, 0, "", err
	}

	if len(*o) == 0 {
		return *a, *b, *o, fmt.Errorf("Error: --oper input cannot be empty")
	}

	found := false
	for _, v := range validOper {
		if *o == v {
			found = true
			break
		}
	}

	if !found {
		return *a, *b, *o, fmt.Errorf("Error: --oper is not add, sub, mult, div")
	}

	return *a, *b, *o, nil

}

func operInputs(a int, b int, oper string) (int, error) {
	var result int
	var err error
	err = nil
	switch oper {
	case "add":
		result = addInputs(a, b)
	case "sub":
		result = subInputs(a, b)
	case "mult":
		result = multInputs(a, b)
	case "div":
		result, err = divInputs(a, b)
	}
	return result, err
}

func addInputs(a int, b int) int {
	return a + b
}

func subInputs(a int, b int) int {
	return a - b
}

func multInputs(a int, b int) int {
	return a * b
}

const ErrorDivideByZero = "Error: cannot divide by 0"

func divInputs(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf(ErrorDivideByZero)
	}
	return a / b, nil
}
