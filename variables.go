package main

import (
	"fmt"
)

// Day 1: learn declare variables, if else and declare functions in GoLang
func main() {
	var height float32 = 1.0
	if (height > 1.6 && height < 2.0) || height == 1 {
		fmt.Println("You are tall enough to join basketball team")
	} else if height >= 2 {
		fmt.Println("You are too tall to join basketball team")
	} else {
		fmt.Println("You are not tall enough to join basketball team")
	}

	var myMail string = "toannpt@gmail.com"
	if lenght := len(myMail); lenght > 20 {
		fmt.Println("Email is too long")
	} else {
		fmt.Println("Email is short enough")
	}

	fullName := concatByDelimiter(" ", "Toan", "NPT")
	fmt.Println(fullName)
}

func concatByDelimiter(delimiter string, values ...string) string {
	var result string
	for i, value := range values {
		result += value
		if i < len(values)-1 {
			result += delimiter
		}
	}
	return result
}
