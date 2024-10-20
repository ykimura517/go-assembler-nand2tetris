package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage: assembler <file.asm>")
	}
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	parser := NewParser(file)
	// symbolTable := NewSymbolTable()
	// code := NewCode()
	fmt.Println(parser)
	fmt.Println(parser.commandType())
	fmt.Println(parser.symbol())

}
