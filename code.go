package main

type Code struct {
}

func (c *Code) dest(mnemonic string) string {
	return ""
}

func (c *Code) comp(mnemonic string) string {
	return ""
}

func (c *Code) jump(mnemonic string) string {
	return ""
}

func NewCode() *Code {
	return &Code{}
}
