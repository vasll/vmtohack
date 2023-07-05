package main

import (
	"time"
	"fmt"
	"os"
	"github.com/docopt/docopt.go"
	"github.com/vasll/jackvmt"
)

/* Docopt args usage */
// TODO make <outfile> default = "out.asm" and optional
const usage = `
Usage: 
  jackvmt <file> <outfile>
  jackvmt -h | --help

Options:
  file        path of the input .vm file
  -h, --help  show this help message and exit`

func main() {
	startTime := time.Now()	// Keep track of execution time

	// Load cli args
	args, _ := docopt.ParseDoc(usage)

	// Create Parser and CodeWriter from input/outupt files
	parser, err := jackvmt.NewParser(args["<file>"].(string))
	if err != nil { 
		fmt.Println("Error with input file")
		os.Exit(-1)
	}
	codeWriter, err := jackvmt.NewCodeWriter(args["<outfile>"].(string))
	if err != nil { 
		fmt.Println("Error with output file")
		os.Exit(-1)
	}

	for parser.HasMoreCommands() {
		// fmt.Printf("%d, %s, %d\n", parser.CommandType, parser.Arg1, parser.Arg2)	// For debugging

		if parser.CommandType == jackvmt.C_Arithmetic {
			codeWriter.WriteArithmetic(parser.Arg1)
		} else if parser.CommandType == jackvmt.C_Push {
			codeWriter.WritePush(parser.Arg1, parser.Arg2)
		} else if parser.CommandType == jackvmt.C_Pop {
			codeWriter.WritePop(parser.Arg1, parser.Arg2)
		}
		// TODO add other cases

		parser.Advance()
	}

	codeWriter.Close()

	fmt.Printf("Took %.7f seconds.\n", time.Since(startTime).Seconds())
}