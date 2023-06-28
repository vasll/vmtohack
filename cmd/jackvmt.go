package main

import (
	"github.com/vasll/jackvmt"
)

func main() {
	parser, _ := jackvmt.NewParser("sample_files/test.vm")
	for parser.HasMoreCommands() {
		parser.Advance()
	}
}