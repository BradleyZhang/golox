package main

import (
	"golox/lox"
	"os"
)

func main() {
	lox := lox.Lox{}
	args := os.Args[1:]
	lox.Main(args)
}
