package main

import (
	"fmt"
	"strings"
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

	fmt.Println(sum(1, 2, 3, 4, 5))

	var fullname string = "Toan Nguyen Phuc Thanh"
	var firstname, lastname = determineFirstNameAndLastName(fullname)
	fmt.Printf("First name: %s, Last name: %s", firstname, lastname)
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

func sum(values ...int) int {
	var result int = 0
	for value := range values {
		result += value
	}

	return result
}

func determineFirstNameAndLastName(fullName string) (firstname string, lastname string) {
	var values []string = strings.Split(fullName, " ")
	if values != nil && len(values) > 0 {
		firstname = values[0]
		for i, item := range values {
			if i > 0 {
				lastname += item + " "
			}
		}
	}
	return firstname, strings.TrimSpace(lastname)
}
