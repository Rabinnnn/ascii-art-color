package functions

import (
	//"ascii_art/functions"
	"fmt"
	"strings"

)

// Prints the respective Ascii Art characters
func PrintWords(input []string, asciiFields []string, match string) {
	for _, word := range input {
		if word == "" {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				j := 0
				for j < len(word) {
					char := rune(word[j])
					if !validChar(char) {
						return
					}
					startPoint := Start(int(char))
					if strings.HasPrefix(word[j:], match) {
						for k := 0; k < len(match); k++ {
							fmt.Print(Color() + asciiFields[Start(int(word[j+k]))+i] + "\033[0m")
						}
						j += len(match)
					} else {
						fmt.Print(asciiFields[startPoint+i])
						j++
					}
				}
				fmt.Println()
			}
		}
	}
}
// Determines starting position of the character
func Start(s int) int {
	pos := int(s-32)*9 + 1
	return pos
}

//Checks the validity of the character
func validChar(s rune) bool {
	if !(s >= ' ' && s <= '~') {
		fmt.Println("Error:" + string(s) + " " + "is not valid character")
		return false
	}
	return true
}
