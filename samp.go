package main

import (
	"fmt"
	"os"
	"strings"
)

var colors = map[string]string{
	"red":    "\033[31m",
	"green":  "\033[32m",
	"blue":   "\033[34m",
	"reset":  "\033[0m",
}

func mainn() {
	if len(os.Args) < 4 {
		printUsage()
		return
	}

	option := os.Args[1]
	if !strings.HasPrefix(option, "--color=") {
		printUsage()
		return
	}

	color := strings.TrimPrefix(option, "--color=")
	colorCode, exists := colors[color]
	if !exists {
		printUsage()
		return
	}

	substring := os.Args[2]
	text := os.Args[3]

	if substring == "" {
		fmt.Println(colorCode + text + colors["reset"])
		return
	}

	coloredText := highlightSubstring(text, substring, colorCode)
	fmt.Println(coloredText)
}

func highlightSubstring(text, substring, colorCode string) string {
	var result strings.Builder

	index := 0
	for {
		start := strings.Index(text[index:], substring)
		if start == -1 {
			result.WriteString(text[index:])
			break
		}

		start += index
		result.WriteString(text[index:start])
		result.WriteString(colorCode + substring + colors["reset"])
		index = start + len(substring)
	}

	return result.String()
}

func printUsage() {
	fmt.Println("Usage: go run . --color=<color> <substring to be colored> \"<string>\"")
}
