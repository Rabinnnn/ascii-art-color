package main

import (
	"os"
	"ascii_art/functions"
	//"flags"
)

func main() {
	input, _ := functions.InputArgs(os.Args)
	asciiFields := functions.FileChoice(os.Args)
	functions.PrintWords(input, asciiFields, os.Args[2])
}
