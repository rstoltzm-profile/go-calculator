package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("GO Calculator App")
	a, b, o, err := parseInputs()
	fmt.Println(a, b, o, err)
	addInputs(a, b)
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

func addInputs(a int, b int) int {
	return a + b
}

func subInputs(a int, b int) int {
	return a - b
}
