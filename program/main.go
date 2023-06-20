package main

import (
	"asciiart"
	"fmt"
	"os"
	"strings"
)

func main() {
	str := ""
	subStr := ""
	fileName := "standard.txt"
	option := []string{"", "", ""} // Option[0] for Output , Option[1] for color , Option[2] for align

	// Check the number of command-line arguments
	if len(os.Args) < 2 || len(os.Args) > 6 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}

	// set the str, subStr, fileName and options
	for _, arg := range os.Args[1:] {
		argLower := strings.ToLower(arg)
		if argLower == "standard" || argLower == "shadow" || argLower == "thinkertoy" || argLower == "extrafont" {
			fileName = argLower + ".txt"
		} else if strings.HasPrefix(argLower, "--output=") {
			option[0] = strings.TrimPrefix(argLower, "--output=")
			if option[0] == "" {
				fmt.Println("Enter a valid name for the output file")
				return
			} else if !strings.HasSuffix(option[0], ".txt") {
				option[0] += ".txt"
			}
		} else if strings.HasPrefix(argLower, "--color=") {
			option[1] = asciiart.ColorSelector(strings.TrimPrefix(argLower, "--color="))
			if option[1] == "" {
				fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
				return
			}
		} else if strings.HasPrefix(argLower, "--align=") {
			align := strings.TrimPrefix(argLower, "--align=")
			if align == "center" || align == "left" || align == "right" || align == "justify" {
				option[2] = align
			} else {
				fmt.Println("You can align only left, right, center or justify")
				return
			}
		} else if str == "" {
			str = arg
		} else {
			if len(str) > len(arg) {
				subStr = arg
			} else {
				subStr = str
				str = arg
			}
			if !strings.Contains(str, subStr) {
				fmt.Println("String does not contain the substring")
				return
			}
		}
	}
	if str == "" {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	} else if str != "" && subStr != "" && option[1] == "" {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
	}

	// If there is an align option the width of the terminal wil be calculated
	width, err := asciiart.CalcWidth()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Validate characters in the input text
	for _, letter := range str {
		if letter < ' ' || letter > '~' {
			fmt.Println("Use English chars only")
			os.Exit(0)
		}
	}
	asciiArr, err := asciiart.AsciiArtArray(fileName)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	printArr := asciiart.MakeTheArt(str, subStr, asciiArr, option, width)
	if option[0] != "" { // output option
		err := asciiart.PrintArrayToFile(printArr, width, option)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	} else { // Print the Art
		for _, asciiline := range printArr {
			if asciiart.GetLength(asciiline) <= width || option[2] == "" {
				finalLine := asciiart.AlignString(asciiline, width, option[2])
				fmt.Println(finalLine)
			} else {
				fmt.Println("Error: The text does not fit in terminal")
				break
			}
		}
	}
}
