package jackvmt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* CommandType enum implementation */
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

/* VM Parser implementation from nand2tetris */
type Parser struct {
	file		*os.File
	scanner 	*bufio.Scanner
	line 		string
	CommandType CommandType
	Arg1		string
	Arg2		int
}

/* Creates a new Parser by opening an input file for reading */
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
		CommandType: 0, Arg1: "", Arg2: 0,
	}
	parser.Advance()
	return parser, nil
}

/* Reads the next command from the input file and makes it the current command */
func (p *Parser) Advance() {
	if p.scanner.Scan() {
		p.line = p.scanner.Text()

		// TODO parse the line and change values of command, arg1, arg2
		fields := strings.Split(p.line, " ")
		for i := range fields {
			fmt.Println(fields[i])
		}
	} else {
		p.line = ""
	}
}

/* Checks if there are more commands in the input file */
func (p *Parser) HasMoreCommands() bool {
	return p.line != ""
}

/* Closes the input file */
func (p *Parser) Close() error {
	return p.file.Close()
}
