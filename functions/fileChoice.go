package functions
import(
	"strings"
	"fmt"
	"os"
)

// FileChoice function reads the standard.txt banner file and returns its contents
func FileChoice(osArgs []string) []string {

		file, err := os.ReadFile("standard.txt")
		asciiFields := strings.Split(string(file), "\n")
		if err != nil {
			fmt.Println(err)
			return []string{}
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
