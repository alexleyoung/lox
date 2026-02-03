package main

import (
	"bufio"
	"fmt"
	"os"
)

var reporter = NewErrorReporter()

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

	err = run(string(f))
	if err != nil {
		if err.Error() == "lexical error" {
			os.Exit(65)
		}
		if err.Error() == "parse error" {
			os.Exit(65)
		}
		if err.Error() == "runtime error" {
			os.Exit(70)
		}
		return err
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
		reporter.Reset()
	}
}

func run(source string) error {
	lexer := NewLexer(source)
	tokens, lexErrors := lexer.ScanTokens()

	for _, err := range lexErrors {
		reporter.Report(err)
	}

	if reporter.HadError() {
		return fmt.Errorf("lexical error")
	}

	parser := NewParser(tokens)
	statements, _ := parser.Parse()

	if parser.HadError() {
		return fmt.Errorf("parse error")
	}

	interpreter := NewInterpreter()
	err := interpreter.Interpret(statements)
	if err != nil {
		reporter.Report(err)
		if runtimeErr, ok := err.(RuntimeError); ok {
			fmt.Println(runtimeErr.Error())
			return fmt.Errorf("runtime error")
		}
		return err
	}

	return nil
}
