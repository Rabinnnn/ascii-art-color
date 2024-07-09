package functions

import (
	"fmt"
	"strings"
)

// Prints the respective Ascii Art characters
func PrintWords(input []string, asciiFields []string, match string) {
	for a, word := range input {
		if word == "" {
			fmt.Println()
		}

		arr := make([]string, 8)

		c := 0
		matchs := 0
		z := 0

		for j := 0; j < len(word); j++ {

			char := rune(word[j])
			for i := 0; i < 8; i++ {

				if !validChar(char) {
					return
				}
				startPoint := Start(int(rune(char)))
				matchs = strings.Index(input[a], match) + z

				if j >= matchs && j < matchs+len(match) && matchs != -1 && strings.ContainsRune(match, char) && strings.Contains(input[a], match) {
					arr[i] += (Color() + asciiFields[startPoint+i] + "\033[0m")
					c++
				} else if matchs == -1 {
					arr[i] += (Color() + asciiFields[startPoint+i] + "\033[0m")
				} else {
					arr[i] += (asciiFields[startPoint+i])
				}
			}
			// key kit kitten ,bye bruyne

			if c/8 == len(match) {
				if j < len(word)-1 {
					input[a] = word[j+1:]
				}
				z = j + 1
				c = 0

			}
			fmt.Println(c/8, matchs, j, len(input[a]), len(word), word)

		}
		fmt.Println(strings.Join(arr, "\n"))
	}
}

// Determines starting position of the character
func Start(s int) int {
	pos := int(s-32)*9 + 1
	return pos
}

// Checks the validity of the character
func validChar(s rune) bool {
	if !(s >= ' ' && s <= '~') {
		fmt.Println("Error:" + string(s) + " " + "is not valid character")
		return false
	}
	return true
}
