package main

type SymbolTable struct {
	symbolTable map[string]int
}

func (s *SymbolTable) addEntry(symbol string, address int) {
	s.symbolTable[symbol] = address
}

func (s *SymbolTable) contains(symbol string) bool {
	_, ok := s.symbolTable[symbol]
	return ok
}

func (s *SymbolTable) getAddress(symbol string) int {
	return s.symbolTable[symbol]
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{}
}
