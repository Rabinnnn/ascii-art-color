package functions
import(
	"strings"
	"fmt"
	"os"
)

// FileChoice function selects the appropriate ASCII art file to read
func FileChoice(osArgs []string) []string {
	var file []byte
	var asciiFields []string
	// split the file
	if len(osArgs) == 3 && osArgs[2] == "-t" {
		filec, err := os.ReadFile("thinkertoy.txt")
		file = filec
		asciiField := strings.Split(string(file), "\r\n")
		asciiFields = asciiField
		if err != nil {
			fmt.Println(err)
			return []string{}
		}
	} else if len(osArgs) == 3 && osArgs[2] == "-sh" {
		filec, err := os.ReadFile("shadow.txt")
		file = filec
		asciiField := strings.Split(string(file), "\n")
		asciiFields = asciiField
		if err != nil {
			fmt.Println(err)
			return []string{}
		}
	} else if len(osArgs) == 2 {
		filec, err := os.ReadFile("standard.txt")
		file = filec
		asciiField := strings.Split(string(file), "\n")
		asciiFields = asciiField
		if err != nil {
			fmt.Println(err)
			return []string{}
		}
	}
	if len(string(file)) == 0 {
		fmt.Println("Empty file")
		return []string{}
	}
	if len(asciiFields) != 856 {
		fmt.Println("The File Banner used has been tampered with")
		return []string{}
	}
	return asciiFields
}
