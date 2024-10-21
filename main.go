package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	const a int = 40
	const b float64 = 3.14
	const c bool = true
	const d string = "hello, Go devs!"

	fmt.Println("hello world " + "Random Number", rand.Int31n(10))
	fmt.Println("Integer: ", a)
	fmt.Println("Float: ", b)
	fmt.Println("Boolean: ", c)
	fmt.Println("String: ", d)
}