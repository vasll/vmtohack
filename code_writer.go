package jackvmt

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// TODO: List of optimizations that can be made
// 1. In WritePop and WritePush, on the 'local', 'argument', 'this', 'that' branch, if index is 0
//    there is no need to calculate address of @SEGMENT+index because it's simply @SEGMENT.
//    Same applies for the 'temp' branch
// 2. Add index limits for segments:
//    'temp' <= 8
//	  'pointer' <= 1
//    ...
// 3. Add comments to be optional into final output .asm file

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
		// TODO check if index is out of bounds
		return cw.writeStringAndFlush(fmt.Sprintf(
			"//push %s %d\n@%d\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n", segment, index, index,
		))
	} else if segment == "local" || segment == "argument" || segment == "this" || segment == "that" {
		// TODO check if index is out of bounds
		spname := getPointerNameFromSegment(segment) 	// SegmentPointer name like LCL, ARG, THIS, THAT
		return cw.writeStringAndFlush(fmt.Sprintf(
			"//push %s %d\n@%s\nD=M\n@%d\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n", segment, index, spname, index,
		))
	} else if segment == "temp" {
		// TODO check if index is out of bounds
		return cw.writeStringAndFlush(fmt.Sprintf(
			"//push temp %d\n@5\nD=M\n@%d\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n", index, index,
		))
	} else if segment == "pointer" {
		// TODO check if index is out of bounds
		return cw.writeStringAndFlush(fmt.Sprintf(
			"//push pointer %d\n@%d\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n", index, (3+index),
		))
	} else if segment == "static" {
		// TODO check if index is out of bounds
		return cw.writeStringAndFlush(fmt.Sprintf(
			"//push static %d\n@%s.%d\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n", index, getRawFileName(cw.file.Name()), index,
		))
	}

	return nil
}

/* Writes to the output file the assembly code that implements the given pop command */
func (cw *CodeWriter) WritePop(segment string, index int) error {
	if segment == "local" || segment == "argument" || segment == "this" || segment == "that" {
		// TODO check if index is out of bounds
		spname := getPointerNameFromSegment(segment) 	// SegmentPointer name like LCL, ARG, THIS, THAT
		return cw.writeStringAndFlush(fmt.Sprintf(
			"//pop %s %d\n@%s\nD=M\n@%d\nD=D+A\n@R13\nM=D\n@SP\nM=M-1\nA=M\nD=M\n@R13\nA=M\nM=D\n", segment, index, spname, index,
		))
	} else if segment == "temp" {
		// TODO check if index is out of bounds
		return cw.writeStringAndFlush(fmt.Sprintf(
			"//pop temp %d\n@5\nD=M\n@%d\nD=D+A\n@R13\nM=D\n@SP\nM=M-1\nA=M\nD=M\n@R13\nA=M\nM=D\n", index, index,
		))
	} else if segment == "pointer" {
		// TODO check if index is out of bounds
		return cw.writeStringAndFlush(fmt.Sprintf(
			"//pop pointer %d\n@SP\nM=M-1\nA=M\nD=M\n@%d\nM=D\n", index, (3+index),
		))
	} else if segment == "static" {
		// TODO check if index is out of bounds
		return cw.writeStringAndFlush(fmt.Sprintf(
			"//pop static %d\n@SP\nM=M-1\nA=M\nD=M\n@%s.%d\nM=D\n", index, getRawFileName(cw.file.Name()), index,
		))
	}

	return nil
}

/* Initialises the required pointers */
func (cw *CodeWriter) initPointers() error {
	// TODO pointer allocation
	return cw.writeStringAndFlush(
		"//init pointers\n@256\nD=A\n@SP\nM=D\n",
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

/* Returns the filename from a path without any extensions */
func getRawFileName(path string) string {
	filename := filepath.Base(path)
	extension := filepath.Ext(filename)
	return filename[0: len(filename) - len(extension)]
}

/* Closes the output file */
func (cw *CodeWriter) Close() error {
	return cw.file.Close()
}