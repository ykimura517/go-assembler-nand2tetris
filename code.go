package main

type Code struct {
}

func (c *Code) dest(mnemonic string) string {
	if mnemonic == "null" {
		return "000"
	} else if mnemonic == "M" {
		return "001"
	} else if mnemonic == "D" {
		return "010"
	} else if mnemonic == "MD" {
		return "011"
	} else if mnemonic == "A" {
		return "100"
	} else if mnemonic == "AM" {
		return "101"
	} else if mnemonic == "AD" {
		return "110"
	} else if mnemonic == "AMD" {
		return "111"
	}
	panic("Invalid mnemonic: " + mnemonic)
}

func (c *Code) comp(mnemonic string) string {
	if mnemonic == "0" {
		return "0101010"
	} else if mnemonic == "1" {
		return "0111111"
	} else if mnemonic == "-1" {
		return "0111010"
	} else if mnemonic == "D" {
		return "0001100"
	} else if mnemonic == "A" {
		return "0110000"
	} else if mnemonic == "!D" {
		return "0001101"
	} else if mnemonic == "!A" {
		return "0110001"
	} else if mnemonic == "-D" {
		return "0001111"
	} else if mnemonic == "-A" {
		return "0110011"
	} else if mnemonic == "D+1" {
		return "0011111"
	} else if mnemonic == "A+1" {
		return "0110111"
	} else if mnemonic == "D-1" {
		return "0001110"
	} else if mnemonic == "A-1" {
		return "0110010"
	} else if mnemonic == "D+A" {
		return "0000010"
	} else if mnemonic == "D-A" {
		return "0010011"
	} else if mnemonic == "A-D" {
		return "0000111"
	} else if mnemonic == "D&A" {
		return "0000000"
	} else if mnemonic == "D|A" {
		return "0010101"
	} else if mnemonic == "M" {
		return "1110000"
	} else if mnemonic == "!M" {
		return "1110001"
	} else if mnemonic == "-M" {
		return "1110011"
	} else if mnemonic == "M+1" {
		return "1110111"
	} else if mnemonic == "M-1" {
		return "1110010"
	} else if mnemonic == "D+M" {
		return "1000010"
	} else if mnemonic == "D-M" {
		return "1010011"
	} else if mnemonic == "M-D" {
		return "1000111"
	} else if mnemonic == "D&M" {
		return "1000000"
	} else if mnemonic == "D|M" {
		return "1010101"
	}
	panic("Invalid mnemonic: " + mnemonic)
}

func (c *Code) jump(mnemonic string) string {
	if mnemonic == "null" {
		return "000"
	} else if mnemonic == "JGT" {
		return "001"
	} else if mnemonic == "JEQ" {
		return "010"
	} else if mnemonic == "JGE" {
		return "011"
	} else if mnemonic == "JLT" {
		return "100"
	} else if mnemonic == "JNE" {
		return "101"
	} else if mnemonic == "JLE" {
		return "110"
	} else if mnemonic == "JMP" {
		return "111"
	}
	panic("Invalid mnemonic: " + mnemonic)
}

func NewCode() *Code {
	return &Code{}
}
