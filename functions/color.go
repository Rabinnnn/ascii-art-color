package functions

import (
	"flag"
	//"fmt"
)

func Color()string{
	var colors = map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"thing": "\x1b[0m",
		"reset":   "\033[0m",
	}

	result := "'"
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "color" {
			result =  string(f.Value.String())
		
		}
	})
	return (colors[result])
}