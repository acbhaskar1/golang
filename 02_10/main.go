package main

import (
	"fmt"
	"strings"
)

func main() {

	// Strings Fields and Methods
	str := "Hello, Go Programming!"
	fmt.Println("Original String:", str)
	fmt.Println("Length of String:", len(str))

	upperStr := "HELLO, GO PROGRAMMING!"
	lowerStr := "hello, go programming!"
	fmt.Println("Uppercase String:", upperStr)
	fmt.Println("Lowercase String:", lowerStr)
	//to convert to upper case
	fmt.Println("To Upper:", strings.ToUpper(str))
	//to convert to lower case
	fmt.Println("To Lower:", strings.ToLower(str))
	//check if a substring is present
	substr := "Go"
	fmt.Printf("Does the string contain '%s'? %v\n", substr, strings.Contains(str, substr))
	//replace a substring
	replacedStr := strings.Replace(str, "Go", "Golang", 1)
	fmt.Println("Replaced String:", replacedStr)
	//split the string into a slice
	words := strings.Fields(str)
	fmt.Println("Split String into Words:", words)
	//join a slice of strings into a single string
	joinedStr := strings.Join(words, " - ")
	fmt.Println("Joined String:", joinedStr)

	someText := "Trim"
	somewords := strings.Fields(someText)
	fmt.Println("Some Words:", somewords)

	// Trim spaces
	textWithSpaces := "   Hello, Go!   "
	trimmedText := strings.TrimSpace(textWithSpaces)
	fmt.Printf("Original: '%s', Trimmed: '%s'\n", textWithSpaces, trimmedText)
	//left trim
	leftTrimmedText := strings.TrimLeft(textWithSpaces, " ")
	fmt.Printf("Left Trimmed: '%s'\n", leftTrimmedText)
	//right trim
	rightTrimmedText := strings.TrimRight(textWithSpaces, " ")
	fmt.Printf("Right Trimmed: '%s'\n", rightTrimmedText)
	//trim specific characters
	textWithChars := "***Hello, Go!***"
	trimmedCharsText := strings.Trim(textWithChars, "*")
	fmt.Printf("Original: '%s', Trimmed Characters: '%s'\n", textWithChars, trimmedCharsText)
	//check prefix
	prefix := "Hello"
	fmt.Printf("Does the string start with '%s'? %v\n", prefix, strings.HasPrefix(str, prefix))
	//check suffix
	suffix := "Programming!"
	fmt.Printf("Does the string end with '%s'? %v\n", suffix, strings.HasSuffix(str, suffix))
	//push a new line for better readability
	fmt.Println()
	// Strings Comparison
	str1 := "GoLang"
	str2 := "golang"
	fmt.Printf("String 1: '%s', String 2: '%s'\n", str1, str2)
	//case-sensitive comparison
	equalCaseSensitive := str1 == str2
	fmt.Printf("Are the strings equal (case-sensitive)? %v\n", equalCaseSensitive)
	//case-insensitive comparison
	equalCaseInsensitive := strings.EqualFold(str1, str2)
	fmt.Printf("Are the strings equal (case-insensitive)? %v\n", equalCaseInsensitive)
	//lexicographical comparison
	compareResult := strings.Compare(str1, str2)
	fmt.Printf("Lexicographical comparison result: %d\n", compareResult)
	//trim regular expressions set of characters
	textWithRegexChars := "!!!Welcome to Go!!!"
	trimmedRegexText := strings.Trim(textWithRegexChars, "!?")
	fmt.Printf("Original: '%s', Trimmed Regex Characters: '%s'\n", textWithRegexChars, trimmedRegexText)
	// trim newline characters, tabs, and spaces, etc.
	textWithWhitespace := "\n\t  Hello, Go!  \t\n"
	trimmedWhitespaceText := strings.TrimSpace(textWithWhitespace)
	fmt.Printf("Original: '%s', Trimmed Whitespace: '%s'\n", textWithWhitespace, trimmedWhitespaceText)
	//trim digitts, special characters, and letters,unnecessary characters
	textWithUnnecessaryChars := "123Hello@Go#456"
	trimmedUnnecessaryText := strings.Trim(textWithUnnecessaryChars, "123456@#")
	fmt.Printf("Original: '%s', Trimmed Unnecessary Characters: '%s'\n", textWithUnnecessaryChars, trimmedUnnecessaryText)
	//regular expression check for alphanumeric characters
	alphanumericText := "GoLang123"
	isAlphanumeric := true
	for _, char := range alphanumericText {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			isAlphanumeric = false
			break
		}
	}
	fmt.Printf("Is the string '%s' alphanumeric? %v\n", alphanumericText, isAlphanumeric)

}
