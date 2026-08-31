package lox

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	lox = Lox{}
)

type Lox struct {
	hadError bool
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
	for _, token := range tokens {
		fmt.Println(token.ToString())
	}
}

func (l *Lox) error(line int, msg string) {
	l.report(line, "", msg)
}
func (l *Lox) report(line int, where string, msg string) {
	fmt.Fprintln(os.Stderr, "[line ", line, "] Error", where, ": ", msg)
	l.hadError = true
}
