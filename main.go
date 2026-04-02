package main

import (
	asciiart "asciiartcolor/MethodsAndTesting"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var colorValue string

	// Base case: only one argument provided
	if len(os.Args) == 2 {
		fmt.Println(asciiart.FormatPrinter(os.Args[1]))
		return
	}

	// Define flag
	flag.StringVar(&colorValue, "color", "white", "stores the color")
	flag.Parse()

	// Get non-flag arguments
	arguments := flag.Args()

	// Ensure there are enough arguments for processing
	if len(arguments) == 0 {
		printUsage()
		return
	}

	// Validate the flag format using raw input
	flagAll := os.Args[1]
	beforeEqualTo, _, withEqualTo := strings.Cut(flagAll, "=")

	if !withEqualTo {
		printUsage()
		return
	}

	// Valid flag case
	if beforeEqualTo == "--color" {
		fmt.Print(asciiart.ApplyColor(colorValue, arguments))
		return
	}

	// Invalid flag cases
	if strings.HasPrefix(beforeEqualTo, "-color") ||
		strings.HasPrefix(beforeEqualTo, ".color") ||
		strings.HasPrefix(beforeEqualTo, "color") ||
		beforeEqualTo != "--color" {

		printUsage()
		return
	}
}

// Extracted usage message for reusability and cleaner main()
func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> \"something\"")
}