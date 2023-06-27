package main

import (
	"fmt"
	"github.com/vasll/jackvmt"
)

func main() {
	parser, _ := jackvmt.NewParser("")
	fmt.Println(parser)
}