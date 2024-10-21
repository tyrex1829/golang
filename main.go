package main

func main() {
	// const a int = 40
	// const b float64 = 3.14
	// const c bool = true
	// const d string = "hello, Go devs!"

	// fmt.Println("hello world " + "Random Number", rand.Int31n(10))
	// fmt.Println("Integer: ", a)
	// fmt.Println("Float: ", b)
	// fmt.Println("Boolean: ", c)
	// fmt.Println("String: ", d)

	// sum := 0
	// for i := range 15 {
	// 	sum += i
	// }

	// println(sum)

	// a := 3
	// if a%2 == 0 {
	// 	println("Even")
	// } else {
	// 	println("Odd")
	// }

	// gender := "male"

	// switch gender {
	// case "male":
	// 	println("Male")
	// case "female":
	// 	println("Female")
	// default:
	// 	println("You are either male or female")
	// }

	var a [5]int = [5]int{1, 2, 3, 4, 5}
	sum := 0

	for i := range a {
		sum += a[i]
	}
	println(sum)
}