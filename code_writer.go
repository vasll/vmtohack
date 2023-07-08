package jackvmt

import (
	"fmt"
	"os"
	"bufio"
)

/* CodeWriter implementation from nand2tetris */
type CodeWriter struct {
	file 	*os.File		// Output file to be written to
	writer	*bufio.Writer	// Output file writer
	eqcount int				// Count of how many eq labels are there
	ltcount int				// Count of how many lt labels are there
	gtcount	int				// Count of how many gt labels are there
}

/* Creates a new CodeWriter by opening an output file to be written to */
func NewCodeWriter(path string) (*CodeWriter, error) {
	// Open output file and check for errors
	file, err := os.Create(path)
	if err != nil { return nil, err }

	writer := bufio.NewWriter(file)
	codeWriter := &CodeWriter{ file, writer, 0, 0, 0}
	
	return codeWriter, codeWriter.initPointers()
}

/* Writes to the output file the assembly code that implements the given arithmetic command */
func (cw *CodeWriter) WriteArithmetic(command string) error {
	// The implementation of each command is in the _implementations folder
	if command == "add" {
		return cw.writeStringAndFlush("//add\n@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=D+M\n@SP\nM=M+1\n")
	} else if command == "sub" {
		return cw.writeStringAndFlush("//sub\n@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=M-D\n@SP\nM=M+1\n")
	} else if command == "neg" {
		return cw.writeStringAndFlush("//neg\n@SP\nM=M-1\nA=M\nM=-M\n@SP\nM=M+1\n")
	} else if command == "and" {
		return cw.writeStringAndFlush("//and\n@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=M&D\n@SP\nM=M+1\n")
	} else if command == "or" {
		return cw.writeStringAndFlush("//or\n@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=M|D\n@SP\nM=M+1\n")
	} else if command == "not" {
		return cw.writeStringAndFlush("//not\n@SP\nM=M-1\nA=M\nM=!M\n@SP\nM=M+1\n")
	} else if command == "eq" {
		err := cw.writeStringAndFlush(fmt.Sprintf(
			"//eq\n@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=D-M\nD=M\n@EQ_%d_T\nD;JEQ\n@SP\nA=M\nM=0\n"+
			"@EQ_%d_END\n0;JMP\n(EQ_%d_T)\n@SP\nA=M\nM=-1\n@EQ_%d_END\n0;JMP\n(EQ_%d_END)\n@SP\nM=M+1\n",
			cw.eqcount, cw.eqcount, cw.eqcount, cw.eqcount, cw.eqcount,
		))
		cw.eqcount++
		return err
	} else if command == "lt" {
		err := cw.writeStringAndFlush(fmt.Sprintf(
			"@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=D-M\nD=M\n@LT_%d_T\nD;JGT\n@SP\nA=M\nM=0\n"+
			"@LT_%d_END\n0;JMP\n(LT_%d_T)\n@SP\nA=M\nM=-1\n@LT_%d_END\n0;JMP\n(LT_%d_END)\n@SP\nM=M+1\n",
			cw.ltcount, cw.ltcount, cw.ltcount, cw.ltcount, cw.ltcount,
		))
		cw.ltcount++
		return err
	} else if command == "gt" {
		err := cw.writeStringAndFlush(fmt.Sprintf(
			"@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=D-M\nD=M\n@GT_%d_T\nD;JLT\n@SP\nA=M\nM=0\n"+
			"@GT_%d_END\n0;JMP\n(GT_%d_T)\n@SP\nA=M\nM=-1\n@GT_%d_END\n0;JMP\n(GT_%d_END)\n@SP\nM=M+1\n",
			cw.gtcount, cw.gtcount, cw.gtcount, cw.gtcount, cw.gtcount,
		))
		cw.gtcount++
		return err
	}

	return fmt.Errorf("arithmetic command '%s' not found", command)
}

/* Writes to the output file the assembly code that implements the given push command */
func (cw *CodeWriter) WritePush(segment string, index int) error {
	if segment == "constant" {
		return cw.writeStringAndFlush(fmt.Sprintf(
			"//push %s %d\n@%d\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n", segment, index, index,
		))
	} else if segment == "local" || segment == "argument" || segment == "this" || segment == "that" {
		// TODO implement (not working)
		spname := getPointerNameFromSegment(segment) 	// SegmentPointer name like LCL, ARG, THIS, THAT
		return cw.writeStringAndFlush(fmt.Sprintf(
			"//push %s %d\n@%s\nD=M\n\n@%d\nA=D+A\nD=M\n\n@SP\nA=M\nM=D\n\n@SP\nM=M+1\n", segment, index, spname, index,
		))
	} else if segment == "temp" {
		// TODO implement
	}

	return nil
}

/* Writes to the output file the assembly code that implements the given pop command */
func (cw *CodeWriter) WritePop(segment string, index int) error {
	if segment == "local" || segment == "argument" || segment == "this" || segment == "that" {
		// TODO implement (not working)
		spname := getPointerNameFromSegment(segment) 	// SegmentPointer name like LCL, ARG, THIS, THAT
		return cw.writeStringAndFlush(fmt.Sprintf(
			"//pop %s %d\n@%s\nD=M\n@%d\nD=D+A\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n@R13\nA=M\nM=D\n", segment, index, spname, index,
		))
	} else if segment == "temp" {
		// TODO implement
	}

	return nil
}

/* Initialises the required pointers */
func (cw *CodeWriter) initPointers() error {
	// TODO these pointers are here just for testing with BasicTest.vm
	return cw.writeStringAndFlush(
		"//init pointers\n@256\nD=A\n@SP\nM=D\n@300\nD=A\n@LCL\nM=D\n@400\nD=A\n@ARG\nM=D\n@3000\nD=A\n@THIS\nM=D\n@3010\nD=A\n@THAT\nM=D\n",
	)
}

/* Writes a string to a *CodeWriter.writer and flushes it */
func (cw *CodeWriter) writeStringAndFlush(s string) error {
	_, err := cw.writer.WriteString(s)	
	if err != nil { return err }
	err = cw.writer.Flush() 
	if err != nil { return err }
	return nil
}

/* Returns the pointer name from a string. Example: "local" -> "LCL" */
func getPointerNameFromSegment(s string) string {
	if s == "local" {
		return "LCL"
	} else if s == "argument" {
		return "ARG"
	} else if s == "this" {
		return "THIS"
	} else if s == "that" {
		return "THAT"
	}
	return "[SEGMENT_NOT_FOUND]"
}

/* Closes the output file */
func (cw *CodeWriter) Close() error {
	return cw.file.Close()
}