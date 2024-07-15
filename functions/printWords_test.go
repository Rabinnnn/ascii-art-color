package functions

import (
	
	"testing"
//	"ascii_art_color/functions"
	"fmt"
	"os"
	"strings"
	
)
func TestPrintWords(t *testing.T) {
	
	params := []string{"main.go","--color=red", "kit", "a kit has kitten"}
	
	//fmt.Println(len(params))
	input, match, _ := InputArgs(params)
	// fmt.Println(input)
	// fmt.Println(match)

	//match := "kit"

	file, err := os.ReadFile("standard.txt")
	if err != nil{
		return
	}
	asciiFields := strings.Split(string(file), "\n")

	//create a map containing the input text as key and the path to its expected output as value
	testCases := map[string]string{
		input: "testCases/expectedOutput1.txt",
		
	}
	for input, expectedFilePath := range testCases {
		//subtest to get the content from the file in the specified filepath
		t.Run(input, func(t *testing.T) {
			expectedContent, err := os.ReadFile(expectedFilePath)
			if err != nil {
				t.Fatalf("Failed to read expected output file '%s' for input '%s': %v", expectedFilePath, input, err)
			}
			//convert content read from the file to string
			expectedContentStr := string(expectedContent)
			result := PrintWords(input, asciiFields, match)
			fmt.Println(result)
			if result != expectedContentStr {
				t.Errorf("For input:\n'%s'\nExpected:\n%s\n but got:\n%s", input, expectedContentStr, result)
			}
		})
	}
}