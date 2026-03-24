package main

import (
	asciiart "asciiartcolor/MethodsAndTesting"
	"flag"
	"fmt"
)

func main() {
	var colorValue string

	flag.StringVar(&colorValue, "color", "white", "The color of the output")
	flag.Parse()
	arguments := flag.Args()

	fmt.Println(asciiart.ApplyColor(colorValue, arguments))
}
