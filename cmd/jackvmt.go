package main

import (
	"github.com/vasll/jackvmt"
)

func main() {
	parser, _ := jackvmt.NewParser("../_sample_files/BasicTest.vm")
	for parser.HasMoreCommands() {
		parser.Advance()
		// TODO
	}
}