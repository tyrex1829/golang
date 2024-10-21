package main

import (
	"fmt"
)

// Defining a struct
type rect struct {
	width int32
	height int32
}

func area(r rect) int32 {
	return r.height * r.width
}

func main () {
	// new instance of rect
	r := rect{width: 200, height: 100}

	fmt.Println(area(r))
}