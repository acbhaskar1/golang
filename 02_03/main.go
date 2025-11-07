package main

import "fmt"

func main() {
	//using a byte slice to modify a string
	s := "hello"
	fmt.Println(s)
	bs := []byte(s)
	bs[0] = 'H'
	bs[2] = 'L'
	bs[4] = 'O'
	fmt.Printf("Byte slice: %v, (type %T) \n", bs, bs)
	s = string(bs)
	fmt.Println(s)
}
