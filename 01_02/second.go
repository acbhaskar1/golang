package main

import (
	"fmt"
)

func main() {
	// var x int
	// var y int
	// x = 10
	// y = 20

	// x := 10
	// y := 20
	// x, y := 10, 20

	// x, y := 10.5, 20.5

	x, y := 10.0, 20.5

	fmt.Printf("The value of x is:%v and type is %T\n", x, x)
	fmt.Printf("The value of y is:%v and type is %T\n", y, y)

	mean := (x + y) / 2
	fmt.Printf("The mean value is:%v and type is %T\n", mean, mean)

}
