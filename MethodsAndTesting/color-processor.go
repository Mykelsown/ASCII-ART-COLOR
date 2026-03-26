// package asciiart

// import (
//     "strings"
// )

// var colorStorage = map[string]string{
//     "red":     "\033[31m",
//     "green":   "\033[32m",
//     "yellow":  "\033[33m",
//     "blue":    "\033[34m",
//     "magenta": "\033[35m",
//     "cyan":    "\033[36m",
//     "white":   "\033[37m",
//     "black":   "\033[30m",

//     // Bright variants
//     "bright_red":     "\033[91m",
//     "bright_green":   "\033[92m",
//     "bright_yellow":  "\033[93m",
//     "bright_blue":    "\033[94m",
//     "bright_magenta": "\033[95m",
//     "bright_cyan":    "\033[96m",
//     "bright_white":   "\033[97m",
// }

// func ApplyColor(colorType string, arguments []string) string {
//     mainString := arguments[len(arguments)-1] // gets the main string in the argument
//     options := arguments[:len(arguments)-1]   // gets the substrings(s) that needs to be colored in the main string from the argument

//     formatedMainString := FormatPrinter(mainString) //This gets the ascii-art of the main string
//     // This is the main string that will store the final output
//     var newFormattedMain strings.Builder

//     // The main logic that will check for each sub-string in the main string and then apply the color to them
//     for _, v := range options {
//         formatedOption := FormatPrinter(v) // This stores the sub-string ascii-art
//         // Checks if the substring passed from the terminal really exist in the main string
//         contains := strings.Contains(mainString, v)

//         splittedFormattedMainString := strings.Split(formatedMainString, "\n")
//         splittedFormattedOption := strings.Split(formatedOption, "\n")

//         var slicedModifiedMainString = make([]string, len(splittedFormattedMainString))
//         if contains {
//             for i, asciiChars := range splittedFormattedMainString {
//                 splittedColoredFormattedOption := colorStorage[colorType] + splittedFormattedOption[i] + "\033[0m"
//                 newAsciiChars := strings.ReplaceAll(asciiChars, splittedFormattedOption[i], splittedColoredFormattedOption)
//                 slicedModifiedMainString[i] = newAsciiChars
//             }
//         }
//         newFormattedMain.WriteString(strings.Join(slicedModifiedMainString, "\n"))
//     }

//     return newFormattedMain.String()
// }

package asciiart

import (
	"strings"
)

var colorStorage = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"black":   "\033[30m",

	"bright_red":     "\033[91m",
	"bright_green":   "\033[92m",
	"bright_yellow":  "\033[93m",
	"bright_blue":    "\033[94m",
	"bright_magenta": "\033[95m",
	"bright_cyan":    "\033[96m",
	"bright_white":   "\033[97m",
}

func ApplyColor(colorType string, arguments []string) string {
	mainString := arguments[len(arguments)-1]
	options := arguments[:len(arguments)-1]

	formattedMain := strings.Split(FormatPrinter(mainString), "\n")

	color := colorStorage[colorType]
	reset := "\033[0m"

	// If no substring → color everything
	if len(options) == 0 {
		var result strings.Builder
		for _, line := range formattedMain {
			result.WriteString(color + line + reset + "\n")
		}
		return result.String()
	}

	// Prepare result lines
	resultLines := make([]string, len(formattedMain))

	// Walk through string (instead of Cut)
	for i := 0; i < len(mainString); {
		matched := false

		for _, sub := range options {
			if sub != "" && strings.HasPrefix(mainString[i:], sub) {
				asciiSub := strings.Split(FormatPrinter(sub), "\n")

				for line := range resultLines {
					resultLines[line] += color + asciiSub[line] + reset
				}

				i += len(sub)
				matched = true
				break
			}
		}

		if !matched {
			char := string(mainString[i])
			asciiChar := strings.Split(FormatPrinter(char), "\n")

			for line := range resultLines {
				resultLines[line] += asciiChar[line]
			}
			i++
		}
	}

	return strings.Join(resultLines, "\n")
}