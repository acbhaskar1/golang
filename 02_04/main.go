package main

import "fmt"

func main() {
	//using a byte slice to modify a string
	s := "hello"
	fmt.Println(s)
	bs := []byte(s)
	for i := range bs {
		switch i {
		case 0:
			bs[i] = 'H'
		case 2:
			bs[i] = 'L'
		case 4:
			bs[i] = 'O'
		}
	}
	fmt.Printf("Byte slice: %v, (type %T) \n", bs, bs)
	s = string(bs)
	fmt.Println(s)
}
