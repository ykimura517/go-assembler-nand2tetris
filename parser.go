package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Parser struct {
	currentCommandIndex int
	allLines            []string
}

func (p *Parser) hasMoreCommands() bool {
	return p.currentCommandIndex < len(p.allLines)
}

func (p *Parser) advance() {
	p.currentCommandIndex++
}

func (p *Parser) commandType() string {
	if p.allLines[p.currentCommandIndex][0] == '@' {
		return "A_COMMAND"
	} else if p.allLines[p.currentCommandIndex][0] == '(' {
		return "L_COMMAND"
	} else {
		return "C_COMMAND"
	}
}

func (p *Parser) symbol() string {
	if !(p.commandType() == "A_COMMAND" || p.commandType() == "L_COMMAND") {
		log.Fatal("Cannot call symbol() on a non-AorL_COMMAND")
	}
	if p.commandType() == "A_COMMAND" {
		return p.allLines[p.currentCommandIndex][1:]
	} else {
		return p.allLines[p.currentCommandIndex][1 : len(p.allLines[p.currentCommandIndex])-1]
	}

}

func (p *Parser) dest() string {
	if p.commandType() != "C_COMMAND" {
		log.Fatal("Cannot call dest() on a non-C_COMMAND")
	}
	currentCommand := p.allLines[p.currentCommandIndex]
	// Check if there is a dest mnemonic
	if strings.Contains(currentCommand, "=") {
		return strings.Split(currentCommand, "=")[0]
	} else {
		return "null"
	}
}

func (p *Parser) comp() string {
	if p.commandType() != "C_COMMAND" {
		log.Fatal("Cannot call comp() on a non-C_COMMAND")
	}
	currentCommand := p.allLines[p.currentCommandIndex]

	if strings.Contains(currentCommand, "=") {
		// If there is a dest mnemonic, split by "=" and take the second part
		if strings.Contains(currentCommand, ";") {
			// dest=comp;jump
			cj := strings.Split(currentCommand, "=")[1]
			return strings.Split(cj, ";")[0]
		} else {
			// dest=comp
			return strings.Split(currentCommand, "=")[1]
		}

	} else {
		return strings.Split(currentCommand, ";")[0]
	}
}

func (p *Parser) jump() string {
	if p.commandType() != "C_COMMAND" {
		log.Fatal("Cannot call jump() on a non-C_COMMAND")
	}
	currentCommand := p.allLines[p.currentCommandIndex]
	// Check if there is a jump mnemonic
	if strings.Contains(currentCommand, ";") {
		return strings.Split(currentCommand, ";")[1]
	} else {
		return "null"
	}

}

func NewParser(inputFileStream *os.File) *Parser {
	reader := bufio.NewReader(inputFileStream) // Initialize the reader once outside the loop
	allLines := make([]string, 0)
	for {
		line, _, err := reader.ReadLine() // Use the same reader to continue where the last ReadLine left off
		if err == io.EOF {
			break
		}

		toAppendRaw := string(line)
		// Remove comments and spaces
		toAppend := strings.Split(toAppendRaw, "//")[0]
		toAppend = strings.TrimSpace(toAppend)

		// Trim whitespace and check if line is empty
		if len(toAppend) == 0 {
			continue
		}
		allLines = append(allLines, toAppend)
	}
	return &Parser{
		currentCommandIndex: 0,
		allLines:            allLines,
	}
}
