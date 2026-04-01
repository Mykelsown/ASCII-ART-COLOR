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

	// If only one argument is passed (program name + 1 argument),
	// run the base ascii-art functionality with a default string
	if len(os.Args) == 2 {
		fmt.Println(asciiart.FormatPrinter("something")) // base project output
		return
	}

	// Get the first CLI argument (expected to be the flag)
	flagAll := os.Args[1]

	// Split the argument at "=" to isolate the flag name (e.g., "--color")
	beforeEqualTo, _, _ := strings.Cut(flagAll, "=")

	// Define the --color flag with a default value of "white"
	flag.StringVar(&colorValue, "color", "white", "stores the color")

	// Parse all flags from command-line input
	flag.Parse()

	// Get remaining arguments after flag parsing
	// These are expected to be: substring + main string
	arguments := flag.Args()

	// If the flag is exactly "--color", apply coloring logic
	if beforeEqualTo == "--color" {

		// Call function that applies color to the given substring within the string
		fmt.Println(asciiart.ApplyColor(colorValue, arguments))

	} else if (
		// Check for invalid flag formats like:
		// "-color", ".color", "color", or anything not exactly "--color"
		strings.HasPrefix(beforeEqualTo, "-color") ||
		strings.HasPrefix(beforeEqualTo, ".color") ||
		strings.HasPrefix(beforeEqualTo, "color")) ||

		// Also reject anything whose length is not equal to "--color" (7 chars)
		len(beforeEqualTo) != 7 {

		// Print usage instructions for incorrect input
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")
		return
	}
}