package main

import (
	"os"
	"ascii_art/functions"

)

func main() {


	input, match, inputLen := functions.InputArgs(os.Args)

	if inputLen == 0{
		return
	}

	asciiFields := functions.FileChoice(os.Args)

	functions.PrintWords(input, asciiFields, match)
}
