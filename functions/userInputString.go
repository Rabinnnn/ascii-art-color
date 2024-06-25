package functions
import (
	"fmt"
	//"os"
	"flag"
	"strings"
)
var args string
// InputArgs function processes command-line arguments
func InputArgs(osArgs []string) ([]string, int) {
	// Variable to track if the flag was set
	var nameSet bool
	var flagSet bool = false
	// Enforce the flag format to be used to be --output=<filename.txt>
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "output" {
			flagSet = true
			result := strings.Replace(os.Args[1], *output, "", 1)
			if !(result == "--output=") {
				nameSet = true
			}
		}
	})
	
	if len(flag.Args()) > 2 || len(flag.Args()) < 1{
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")
		return []string{}, 0
	} else if len(flag.Args()) == 2{
		args = flag.Args()[1]
	} else if len(flag.Args()) && fla== 1{
		args = osArgs[0]
	}
	args = strings.ReplaceAll(args, "\n", "\\n")
	args = strings.ReplaceAll(args, "\\t", " ")
	input := strings.Split(args, "\\n")
	if input[0] == "" && len(input) == 1 {
		return []string{}, 0
	} else if input[0] == "" && input[1] == "" && len(input) == 2 {
		fmt.Println()
		return []string{}, 0
	}
	return input, len(osArgs)
}
