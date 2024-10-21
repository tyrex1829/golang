package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	a := 40
	b := 3.14
	c := true
	d := "hello, Go devs!"

	fmt.Println("hello world " + "Random Number", rand.Int31n(10))
	fmt.Println("Integer: ", a)
	fmt.Println("Float: ", b)
	fmt.Println("Boolean: ", c)
	fmt.Println("String: ", d)
}