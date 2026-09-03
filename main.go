package main

import (
	"golox/lox"
	"os"
)

func main() {
	args := os.Args[1:]
	lox.GlobalLox.Main(args)
}
