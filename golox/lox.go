package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexleyoung/golox/fe"
)

var hasError = false

func main() {
	args := os.Args
	switch len(args) {
	case 1:
		runPrompt()
	case 2:
		err := runFile(args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	}
}

func runFile(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	run(string(f))

	if hasError {
		os.Exit(65)
	}

	return nil
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("> ")

		in := scanner.Scan()
		if !in {
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
			break
		}

		line := scanner.Text()
		run(line)
		hasError = false
	}
}

func run(source string) {
	lexer := fe.NewLexer(source)
	tokens := lexer.ScanTokens()
	for _, token := range tokens {
		fmt.Println(token)
	}
}

func Error(line int, msg string) {
	report(line, "", msg)
}

func report(line int, where, msg string) {
	fmt.Println("[line ", line, "] Error", where, ": ", msg)
	hasError = true
}
