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
	flagType := os.Args[1]

	flag.StringVar(&colorValue, "color", "white", "The color of the output")

	flag.Parse()
	arguments := flag.Args()

	// This checks for wrong flag input the user might pass in through the terminal
	if flagType == "--color="+colorValue && colorValue != "white" {
		fmt.Println(asciiart.ApplyColor(colorValue, arguments))
		return
	} else if strings.HasPrefix(flagType, "-color=") || strings.HasPrefix(flagType, "color=") || strings.HasPrefix(flagType, ".color=") {
		fmt.Println(`Usage: go run . [OPTION] [STRING]
EX: go run . --color=<color> <substring to be colored> "something"`)
		return
	} else {
		fmt.Println(asciiart.FormatPrinter(arguments[len(arguments)-1])) // this allows for validation of the base ascii-art project
	}
}
