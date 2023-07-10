package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/docopt/docopt.go"
	"github.com/vasll/jackvmt"
)

/* Docopt args usage */
// TODO make <outfile> param default as simply the input <file> with an .asm extension instead of the given one
const usage = `
Usage: 
  jackvmt <file> [<outfile>]
  jackvmt -h | --help

Options:
  file        path of the input .vm file
  outfile     path of the output .asm file (default: input file with .asm extension)
  -h, --help  show this help message and exit`

func main() {
	startTime := time.Now()	// Keep track of execution time

	// Load cli args
	args, _ := docopt.ParseDoc(usage)

	// Create Parser and CodeWriter from input/output files
	infile := args["<file>"].(string)
	parser, err := jackvmt.NewParser(infile)
	if err != nil { 
		fmt.Println("Error with input file")
		os.Exit(-1)
	}

	// If output file is not given, make it by adding a .asm extension to the original file name
	outfile := ""
	if args["<outfile>"] == nil {
		outfile = replaceExtension(infile, ".asm")
	}
	codeWriter, err := jackvmt.NewCodeWriter(outfile)
	if err != nil { 
		fmt.Println("Error with output file")
		os.Exit(-1)
	}

	for parser.HasMoreCommands() {
		if parser.CommandType == jackvmt.C_Arithmetic {
			codeWriter.WriteArithmetic(parser.Arg1)
		} else if parser.CommandType == jackvmt.C_Push {
			codeWriter.WritePush(parser.Arg1, parser.Arg2)
		} else if parser.CommandType == jackvmt.C_Pop {
			codeWriter.WritePop(parser.Arg1, parser.Arg2)
		}
		parser.Advance()
	}

	codeWriter.Close()

	fmt.Printf("Took %.7f seconds.\n", time.Since(startTime).Seconds())
}

func replaceExtension(path string, newExtension string) string {
	ext := filepath.Ext(path)
	base := strings.TrimSuffix(path, ext)
	return base + newExtension
}