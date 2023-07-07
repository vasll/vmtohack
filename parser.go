package jackvmt

import (
	"bufio"
	"os"
	"strconv"
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

/* List of arithmetic symbols */
var arithmeticSymbols []string = []string {
	"add", "sub", "net", "eq", "gt", "lt", "and", "or", "not",
}

/* VM Parser implementation from nand2tetris */
type Parser struct {
	file		*os.File		// Input file
	scanner 	*bufio.Scanner	// Scanner of input file
	line 		string			// Content of current line in textfile, uncommented
	CommandType CommandType		// Current CommandType
	Arg1		string			// Current Arg1 (if any)
	Arg2		int				// Current Arg2 (if any)
}

const END_COMMANDS = "$END_COMMANDS$" // Signals the end of commands in the file

/* Returns true if array s contains string str, false otherwise */
func arrayContains(s []string, str string) bool {
	for _, v := range s {
		if v == str { return true }
	}
	return false
}

/* Removes comments from a string and applies strings.TrimSpace() to it */
func removeComments(s string) string{
	if strings.ContainsAny(s, "//") {
		return strings.TrimSpace(strings.Split(s, "//")[0])
	}
	return strings.TrimSpace(s)
}

/* Creates a new Parser by opening an input file for reading */
func NewParser(path string) (*Parser, error) {
	// Open input file and check for errors
	file, err := os.Open(path)
	if err != nil { return nil, err }

	scanner := bufio.NewScanner(file)
	parser := &Parser{
		file: file, 
		scanner: scanner, 
		line: "", 
		CommandType: 0, Arg1: "", Arg2: 0,
	}
	parser.Advance()	// Read the first command
	return parser, nil
}

/* Reads the next command from the input file and makes it the current command */
func (p *Parser) Advance() {
	for p.scanner.Scan() {	// Read the file until a commands is found
		p.line = removeComments(p.scanner.Text())
		if len(p.line) == 0 { continue }  // Skip empty lines, go to next line

		fields := strings.Split(p.line, " ")	// Split command into its fields

		// Check command type
		if arrayContains(arithmeticSymbols, fields[0]) {	// Arithmetic command
			p.CommandType = C_Arithmetic
			p.Arg1 = fields[0]
		}else if fields[0] == "push" {	// Push commnand [push segment number]
			p.CommandType = C_Push	// push
			p.Arg1 = fields[1]		// segment

			number, err := strconv.Atoi(fields[2])
			if err != nil {
				// TODO handle conversion error
				os.Exit(-1)
			}
			p.Arg2 = number			// number
		}else if fields[0] == "pop" {	// Pop command [pop segment number]
			p.CommandType = C_Pop	// pop
			p.Arg1 = fields[1]		// segment

			number, err := strconv.Atoi(fields[2])
			if err != nil {
				// TODO handle conversion error
				os.Exit(-1)
			}
			p.Arg2 = number			// number
		}

		return
	}

	p.line = END_COMMANDS // If there are no more commands, turn p.line into END_COMMANDS conts
}

/* Checks if there are more commands in the input file */
func (p *Parser) HasMoreCommands() bool {
	return p.line != END_COMMANDS
}

/* Closes the input file */
func (p *Parser) Close() error {
	return p.file.Close()
}
