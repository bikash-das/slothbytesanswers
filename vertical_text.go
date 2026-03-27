/**

Vertical Text: Create a function that converts a string into a matrix of characters that can be read vertically.
Add spaces when characters are missing.

Examples:
vertical_txt("Holy bananas")
output = [
	["h", "b"],
	["o", "a"],
	["l", "n"],
	["y", "a"],
	["",  "n"],
	["",  "a"],
	["",  "s"],
]
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "The Quick Brown Fox Jumps Over The Lazy Dog...."
	var splittedStr []string = strings.Split(name, " ")
	var totalWords int = len(splittedStr)

	var lenOfBiggestWord int
	for _, word := range splittedStr {
		if len(word) > lenOfBiggestWord {
			lenOfBiggestWord = len(word)
		}
	}

	var matrix [][]string = make([][]string, lenOfBiggestWord)
	for i := range matrix {
		matrix[i] = make([]string, totalWords)
	}

	for i := 0; i < totalWords; i++ {
		word := splittedStr[i]
		for j := 0; j < lenOfBiggestWord; j++ {
			if len(word) > j {
				matrix[j][i] = string(word[j])
			} else {
				matrix[j][i] = " "
			}
		}
	}
	for _, row := range matrix {
		fmt.Println(strings.Join(row, " "))
	}
}
