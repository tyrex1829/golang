package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	var a int = 40
	var b float64 = 3.14
	var c bool = true
	var d string = "hello, Go devs!"

	fmt.Println("hello world " + "Random Number", rand.Int31n(10))
	fmt.Println("Integer: ", a)
	fmt.Println("Float: ", b)
	fmt.Println("Boolean: ", c)
	fmt.Println("String: ", d)
}