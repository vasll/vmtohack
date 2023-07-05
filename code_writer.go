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
		_, err := cw.writer.WriteString("//add\n@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=D+M\n")	// from _implementations/add.asm
		if err != nil { return err }
		err = cw.writer.Flush()
		if err != nil { return err }
		return nil
	} else if command == "eq" {
		// TODO
		return nil
	}
	// TODO add other cases

	return fmt.Errorf("Arithmetic command '%s' not found", command)
}

/* Writes to the output file the assembly code that implements the given push command */
func (cw *CodeWriter) WritePush(segment string, index int) error {
	_, err := cw.writer.WriteString(
		fmt.Sprintf("//push %s %d\n@%d\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n", segment, index, index),	// from _implementations/push.asm
	)	
	if err != nil { return err }
	err = cw.writer.Flush()
	if err != nil { return err }
	return nil
}

/* Writes to he output file the assembly code that implements the given pop command */
func (cw *CodeWriter) WritePop(segment string, index int) error {
	// TODO
	return nil
}

/* Initialises the SP at RAM[0] with value 256*/
func (cw *CodeWriter) initPointers() error {
	_, err := cw.writer.WriteString("//@SP = 256\n@256\nD=A\n@SP\nM=D\n")
	if err != nil { return err }
	err = cw.writer.Flush()
	if err != nil { return err }
	return nil
}

/* Closes the output file */
func (cw *CodeWriter) Close() error {
	return cw.file.Close()
}