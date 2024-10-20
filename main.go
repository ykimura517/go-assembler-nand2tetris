package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

// convertToBinaryPadded takes a string representing an integer and converts it to
// a binary string padded to 15 characters with leading zeros.
func convertToBinaryPadded(s string) (string, error) {
	// Convert the string to an integer
	num, err := strconv.Atoi(s)
	if err != nil {
		return "", err // Handle the error if the conversion fails
	}

	// Convert the integer to a binary string
	binaryStr := strconv.FormatInt(int64(num), 2)

	// Pad the binary string with leading zeros to ensure it is 15 characters long
	paddedBinary := fmt.Sprintf("%015s", binaryStr)

	return paddedBinary, nil
}

// containsNonNumeric checks if a string contains any character that is not a digit.
func containsNonNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage: assembler <file.asm>")
	}
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}

	binaryCode := []string{}
	parser := NewParser(file)
	code := NewCode()
	symbolTable := NewSymbolTable()

	// First pass
	// Add labels to symbol table

	romAddress := 0
	for parser.hasMoreCommands() {
		if parser.commandType() == "C_COMMAND" {
			romAddress++
		} else if parser.commandType() == "A_COMMAND" {
			romAddress++
		} else if parser.commandType() == "L_COMMAND" {
			symbol := parser.symbol()
			if containsNonNumeric(symbol) && !symbolTable.contains(symbol) {
				// fmt.Println("Adding label to symbol table: ", symbol, romAddress)
				symbolTable.addEntry(symbol, romAddress)
			}
		}

		parser.advance()
	}

	// Second pass
	file, err = os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	parser = NewParser(file)
	definedVariables := 16
	for parser.hasMoreCommands() {
		if parser.commandType() == "C_COMMAND" {
			dest := code.dest(parser.dest())
			comp := code.comp(parser.comp())
			jump := code.jump(parser.jump())
			binaryCode = append(binaryCode, "111"+comp+dest+jump)
		} else if parser.commandType() == "A_COMMAND" {

			symbol := parser.symbol()

			if containsNonNumeric(symbol) {
				// if the symbol is in the symbol table,
				if symbolTable.contains(symbol) {
					address := symbolTable.getAddress(symbol)
					binaryPadded, err := convertToBinaryPadded(strconv.Itoa(address))
					if err != nil {
						log.Fatal(err)
					}
					binaryCode = append(binaryCode, "0"+binaryPadded)
				} else {
					// if the symbol is not in the symbol table, add it to the symbol table
					symbolTable.addEntry(symbol, definedVariables)
					binaryPadded, err := convertToBinaryPadded(strconv.Itoa(definedVariables))
					if err != nil {
						log.Fatal(err)
					}
					binaryCode = append(binaryCode, "0"+binaryPadded)
					definedVariables++

				}
			} else {
				// if the symbol only contains numbers, convert it to binary
				fmt.Println(symbol)
				binaryPadded, err := convertToBinaryPadded(symbol)
				if err != nil {
					log.Fatal(err)
				}
				binaryCode = append(binaryCode, "0"+binaryPadded)
			}
		}
		parser.advance()
	}
	fmt.Println(binaryCode)
	fmt.Println(symbolTable.symbolTable)
	// write to file
	outputFile, err := os.Create("output.hack")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()
	for _, line := range binaryCode {
		outputFile.WriteString(line + "\n")
	}
}
