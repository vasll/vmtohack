package jackvmt

import "os"

/* CodeWriter implementation from nand2tetris */
type CodeWriter struct {
	file *os.File
}

/* Creates a new CodeWriter by opening an output file to be written to */
func NewCodeWriter(path string) (*CodeWriter, error) {
	// TODO
}

/* Writes to the output file the assembly code that implements the given arithmetic command */
func WriteArithmetic(command string) error {
	// TODO
}

/* Writes to he output file the assembly code that implements the given push command */
func WritePush(command string) error {
	// TODO
}

/* Writes to he output file the assembly code that implements the given pop command */
func WritePop(command string) error {
	// TODO
}

/* Closes the output file */
func Close() error {
	// TODO
}