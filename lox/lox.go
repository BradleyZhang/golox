package lox

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	GlobalLox = NewLox()
)

type Lox struct {
	hadError        bool
	hadRuntimeError bool
	interpreter     Interpreter
}

func NewLox() *Lox {
	return &Lox{
		interpreter: Interpreter{},
	}
}

func (l *Lox) Main(args []string) {
	if len(args) > 1 {
		fmt.Println("Usage: golox [script]")
	} else if len(args) == 1 {
		err := l.runFile(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1) //TODO 错误码
		}
	} else {
		l.runPrompt()
	}
}

func (l *Lox) runFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	l.run(string(data))
	if l.hadError {
		os.Exit(65)
	}
	if l.hadRuntimeError {
		os.Exit(70)
	}
	return nil
}

func (l *Lox) runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return err
			}
		}
		l.run(line)
		l.hadError = false
	}
	return nil
}
func (l *Lox) run(source string) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens()
	fmt.Println("##### tokens #####")
	fmt.Println(PrintTokens(tokens))

	fmt.Println("##### ST #####")
	parser := NewParser(tokens)
	expression := parser.Parse()
	if l.hadError {
		return
	}
	astPrinter := AstPrinter{}
	fmt.Print(astPrinter.Print(expression))
	fmt.Println()

	fmt.Println("##### Value #####")
	l.interpreter.Interpret(expression)

}

func (l *Lox) LineError(line int, msg string) {
	l.report(line, "", msg)
}
func (l *Lox) TokenError(token Token, msg string) {
	if token.Type == EOF {
		l.report(token.Line, "at end", msg)
	} else {
		l.report(token.Line, "at '"+token.Lexeme+"'", msg)
	}
}
func (l *Lox) RuntimeError(err *RuntimeError) {
	fmt.Fprintln(os.Stderr, err.Error(), "\n[line ", err.Token.Line, "]")
	l.hadRuntimeError = true
}

// func (l *Lox)
func (l *Lox) report(line int, where string, msg string) {
	fmt.Fprintln(os.Stderr, "[line ", line, "] Error", where, ": ", msg)
	l.hadError = true
}
