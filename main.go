package main

import (
	asciiart "asciiartcolor/MethodsAndTesting"
	"flag"
	"fmt"
	"os"
)

func main() {
	var colorValue string
	flagType := os.Args[1]

	flag.StringVar(&colorValue, "color", "white", "The color of the output")

	flag.Parse()
	arguments := flag.Args()
	// This checks for wrong flag input the user might pass in through the terminal
	if flagType != "--color="+colorValue {
		fmt.Println(`Usage: go run . [OPTION] [STRING]
EX: go run . --color=<color> <substring to be colored> "something"`)
		return
	}
	fmt.Println(asciiart.ApplyColor(colorValue, arguments))
}
