package main

import (
	"os"
	"fmt"
	"ascii_art_color/functions"
)

func main() {
	input, match, inputLen := functions.InputArgs(os.Args)
	// fmt.Println(input)
	// fmt.Println(match)

	if inputLen == 0 {
		return
	}

	asciiFields := functions.FileChoice(os.Args)
	if len(asciiFields) == 0 {
		return
	}

	output := functions.PrintWords(input, asciiFields, match)
	if output == ""{
		return
	}else{
		fmt.Println(output)
	}
}
