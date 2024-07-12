package functions

import (
	"fmt"
	"strings"
)

// Prints the respective Ascii Art characters
func PrintWords(input1 string, asciiFields []string, match string) {
	input2 := strings.Split(input1, "\\n")

	if isSliceEmpty(input2) {
		for i := 0; i < len(input2)-1; i++ {
			if input2[i] == "" {
				fmt.Println()
				continue
			}
		}

		return

	}

	input := strings.ReplaceAll(input1, "\\n", "\n")
	match = strings.ReplaceAll(match, "\\n", "\n")
	// fmt.Println(match)

	// fmt.Println(input)
	// j := 0
	// for _, word := range input1 {
	// 	if word == '\n' {
	// 		fmt.Println()
	// 		return
	// 	}
	// }
	word := input
	arr := make([]string, 8)

	c := 0
	matchs := 0
	z := 0

	for j, char := range word {

		if char == '\n' {
			if isSliceEmpty(arr) {
				fmt.Println()
			} else {
				fmt.Println(strings.Join(arr, "\n"))
				arr = make([]string, 8)
			}
			continue
		}
		if !validChar(char) {
			return
		}
		for i := 0; i < 8; i++ {
			startPoint := Start(int(char))
			matchs = strings.Index(word, match) + z

			if j >= matchs && j < matchs+len(match) && matchs != -1 && strings.ContainsRune(match, char) && strings.Contains(word, match) {
				arr[i] += (Color() + asciiFields[startPoint+i] + "\033[0m")
				c++
			} else if matchs == -1 {
				arr[i] += (Color() + asciiFields[startPoint+i] + "\033[0m")
			} else {
				arr[i] += (asciiFields[startPoint+i])
			}
		}

		if (c/8 == len(match) && !strings.Contains(match, "\n")) || (c/8 == len(match)-1 && strings.Contains(match, "\n")) {
			if j < len(word)-1 {
				word = word[j+1:]
			}
			z = j + 1
			c = 0

		}
		// fmt.Println(c/8, matchs, j, len(input1), len(word), word, match, input)

	}
	fmt.Println(strings.Join(arr, "\n"))
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

func isSliceEmpty(arr []string) bool {
	for _, v := range arr {
		if v != "" {
			return false
		}
	}
	return true
}
