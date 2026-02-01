package main

import (
	"bufio"
	"fmt"
	"os"
)

var interpreter = Interpreter{}
var hadError = false

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

	if hadError {
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
		hadError = false
	}
}

func run(source string) {
	lexer := NewLexer(source)
	tokens := lexer.ScanTokens()
	parser := NewParser(tokens)
	statements, _ := parser.Parse()

	if hadError {
		return
	}

	// fmt.Println((&AstPrinter{}).Print(expr))
	interpreter.Interpret(statements)
}

func LexError(line int, msg string) {
	report(line, "", msg)
}

func ParseError(tok Token, msg string) {
	if tok.Type == EOF {
		report(tok.Line, " at end", msg)
	} else {
		report(tok.Line, " at '"+tok.Lexeme+"'", msg)
	}
}

func report(line int, where, msg string) {
	fmt.Println("[line ", line, "] Error", where, ": ", msg)
	hadError = true
}
