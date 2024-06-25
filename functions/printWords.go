package functions
import(
	"fmt"
	"strings"
)
// printing of the characters
func PrintWords(input []string, asciiFields []string, match string) {
	for _, word := range input {
		if word == "" {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				for j, char := range word {
					if !validChar(char) {
						return
					}
					startPoint := Start(int(char))
					if j >= strings.Index(word, match) && j < strings.Index(word, match)+len(match) && strings.ContainsRune(match, char) && strings.Contains(word, match) {
						fmt.Print("\033[31m" + asciiFields[startPoint+i] + "\033[0m")
					} else {
						fmt.Print(asciiFields[startPoint+i])
					}
				}
				fmt.Println()
			}
		}
	}
}

// starting positions of the characters
func Start(s int) int {
	pos := int(s-32)*9 + 1
	return pos
}

func validChar(s rune) bool {
	if !(s >= ' ' && s <= '~') {
		fmt.Println("Error:" + string(s) + " " + "is not valid character")
		return false
	}
	return true
}
