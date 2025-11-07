package main

import "fmt"

func main() {
	//slicing no starting index
	s := "hello"
	fmt.Println(s)
	//length of the string
	fmt.Println("Length of string:", len(s))
	//slicing with no starting index
	fmt.Println(s[2:]) // from index 2 to end
	//slicing no ending index
	fmt.Println(s[:3]) // from start to index 3 (not included)
	//slicing with both starting and ending index
	fmt.Println(s[1:4]) // from index 1 to index 4 (not included)
	//slicing with full range
	fmt.Println(s[:]) // from start to end
	//slicing single character
	fmt.Println(s[1:2]) // from index 1 to index 2 (not included)
	//slicing to get the last character
	fmt.Println(s[len(s)-1:]) // last character
	//slicing to get all but not the last character
	fmt.Println(s[:len(s)-1]) // all but not last character
	//slicing to get the middle part
	fmt.Println(s[1 : len(s)-1]) // from index 1 to len(s)-1 (not included)
	//slicing with out of range indices (will cause a runtime panic if uncommented)
	// fmt.Println(s[0:6]) // index out of range
	// fmt.Println(s[-1:3]) // index out of range
	// fmt.Println(s[3:2]) // invalid slice index
	// fmt.Println(s[5:5]) // valid but empty string
	// fmt.Println(s[3:10]) // index out of range
	// fmt.Println(s[-3:-1]) // index out of range
	//
	//slicing empty string
	empty := ""
	fmt.Println("Empty string slicing:", empty[:]) // should print an empty string
	// type of string empty slice
	fmt.Printf("Empty string slice: %s, (type %T) \n", empty[:], empty[:])
	//type of sliced string
	sliced := s[1:4]
	fmt.Printf("Sliced string: %s, (type %T) \n", sliced, sliced)

	//slicing single character string
	single := "A"
	fmt.Println("Single character string slicing:", single[:]) // should print "A"
	//slicing unicode string
	unicodeStr := "こんにちは"                                   // "Hello" in Japanese
	fmt.Println("Unicode string slicing:", unicodeStr[2:5]) // slicing part of the unicode string
	//slicing with multi-byte characters
	multiByteStr := "Go语言"
	fmt.Printf("Multi-byte string slicing %s and length %d , (type %T):\n", multiByteStr[2:5], len(multiByteStr), multiByteStr[2:3]) // slicing part of the multi-byte string
	//slicing with step is not directly supported in Go, but can be simulated
	// by manually constructing the result
	simulatedStep := ""
	for i := 0; i < len(s); i += 2 {
		simulatedStep += string(s[i])
	}
	fmt.Println("Simulated slicing with step of 2:", simulatedStep) // should print "hlo"
	//

}
