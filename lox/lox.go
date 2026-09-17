package lox

import (
	"bufio"
	"errors"
	"fmt"
	interpreter2 "golox/lox/interpreter"
	parser2 "golox/lox/parser"
	scanner2 "golox/lox/scanner"
	"golox/lox/token"
	"io"
	"os"
)

var (
	GlobalLox = NewLox()
)

type Lox struct {
	hadError        bool
	hadRuntimeError bool
	interpreter     interpreter2.Interpreter
}

func NewLox() *Lox {
	return &Lox{
		interpreter: interpreter2.Interpreter{},
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
	scanner := scanner2.NewScanner(source)
	tokens, err := scanner.ScanTokens()
	for _, e := range err {
		l.LineError(e.Line, e.Msg)
	}
	fmt.Println("##### tokens #####")
	fmt.Println(token.PrintTokens(tokens))

	p := parser2.NewParser(tokens)
	statements, _err := p.Parse()
	if _err != nil {
		l.TokenError(_err.Token, _err.Msg)
	}
	if l.hadError {
		return
	}
	fmt.Println("##### Output #####")
	if err := l.interpreter.Interpret(statements); err != nil {
		l.RuntimeError(err)
	}

}

func (l *Lox) LineError(line int, msg string) {
	l.report(line, "", msg)
}
func (l *Lox) TokenError(t token.Token, msg string) {
	if t.Type == token.EOF {
		l.report(t.Line, "at end", msg)
	} else {
		l.report(t.Line, "at '"+t.Lexeme+"'", msg)
	}
}
func (l *Lox) RuntimeError(err *interpreter2.RuntimeError) {
	fmt.Fprintln(os.Stderr, err.Error(), "\n[line ", err.Token.Line, "]")
	l.hadRuntimeError = true
}

// func (l *Lox)
func (l *Lox) report(line int, where string, msg string) {
	fmt.Fprintln(os.Stderr, "[line ", line, "] Error", where, ": ", msg)
	l.hadError = true
}
