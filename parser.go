package jackvmt

import (
	"bufio"
	"os"
)

// CommandType Enum
type CommandType uint8
const (
	C_Arithmetic CommandType = iota
	C_Push
	C_Pop
	C_Label
	C_Goto
	C_If
	C_Function
	C_Return
	C_Call
)

type Parser struct {
	file		*os.File
	scanner 	*bufio.Scanner
	line 		string
	command 	CommandType
	arg1		string
	arg2		int
}

func NewParser(path string) (*Parser, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	parser := &Parser{
		file: file, 
		scanner: scanner, 
		line: "", 
		command: 0, arg1: "", arg2: 0,
	}
	parser.Advance()
	return parser, nil
}

func (p *Parser) Advance() {
	if p.scanner.Scan() {
		p.line = p.scanner.Text()

		// TODO parse the line and change values of command, arg1, arg2

	} else {
		p.line = ""
	}
}

func (p *Parser) HasMoreCommands() bool {
	return p.line != ""
}

func (p *Parser) Close() error {
	return p.file.Close()
}
