package jackvmt

import (
	"fmt"
	"os"
	"bufio"
)

/* CodeWriter implementation from nand2tetris */
type CodeWriter struct {
	file 	*os.File
	writer	*bufio.Writer
}

/* Creates a new CodeWriter by opening an output file to be written to */
func NewCodeWriter(path string) (*CodeWriter, error) {
	// Open output file and check for errors
	file, err := os.Create(path)
	if err != nil { return nil, err }

	writer := bufio.NewWriter(file)
	codeWriter := &CodeWriter{ file, writer }
	
	return codeWriter, codeWriter.initPointers()
}

/* Writes to the output file the assembly code that implements the given arithmetic command */
func (cw *CodeWriter) WriteArithmetic(command string) error {
	if command == "add" {
		// from _implementations/add.asm
		return cw.writeStringAndFlush("//add\n@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=D+M\n@SP\nM=M+1\n")
	} else if command == "sub" {
		// TODO implement
	} else if command == "neg" {
		// TODO implement
	} else if command == "and" {
		// TODO implement
	} else if command == "or" {
		// TODO implement
	} else if command == "not" {
		// TODO implement
	} else if command == "eq" {
		// TODO implement
	} else if command == "lt" {
		// TODO implement
	} else if command == "gt" {
		// TODO implement
	}

	return fmt.Errorf("Arithmetic command '%s' not found", command)
}

/* Writes to the output file the assembly code that implements the given push command */
func (cw *CodeWriter) WritePush(segment string, index int) error {
	// from _implementations/push.asm
	return cw.writeStringAndFlush(
		fmt.Sprintf("//push %s %d\n@%d\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n", segment, index, index),
	)
}

/* Writes to he output file the assembly code that implements the given pop command */
func (cw *CodeWriter) WritePop(segment string, index int) error {
	// TODO implement
	return nil
}

/* Initialises the SP at RAM[0] with value 256 */
func (cw *CodeWriter) initPointers() error {
	// TODO implement all the other pointers like LCL, ARG, THIS etc...
	return cw.writeStringAndFlush("//@SP = 256\n@256\nD=A\n@SP\nM=D\n")
}

/* Writes a string to a *CodeWriter.writer and flushes it */
func (cw *CodeWriter) writeStringAndFlush(s string) error {
	_, err := cw.writer.WriteString(s)	
	if err != nil { return err }
	err = cw.writer.Flush()
	if err != nil { return err }
	return nil
}

/* Closes the output file */
func (cw *CodeWriter) Close() error {
	return cw.file.Close()
}