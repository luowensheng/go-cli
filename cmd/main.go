package main

import (
	"fmt"
	"log"

	args "github.com/luowensheng/go-cli/arg_parse"
)

func main() {

	parser := args.NewArgParser(`
	Simple CLI Calculator
	`)

	choice := parser.GetArgument(args.Arg{
		Name:    "operation",
		Help:    "Operation to perform",
		Choices: []any{"add", "subtract", "multiply", "divide"},
	}).Value

	number1, err := parser.GetArgument(args.Arg{
		Name:     "num1",
		Required: true,
		Help:     "First number",
	}).IntoInt()

	if err != nil {
		log.Fatal(err)
	}

	number2, err := parser.GetArgument(args.Arg{
		Name:     "num2",
		Required: true,
		Help:     "Second number",
	}).IntoInt()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(choice, number1, number2)
}
