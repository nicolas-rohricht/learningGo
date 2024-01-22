package main

import "fmt"

var x int = 10

const y int = 20

// Course's solution
var a = 40

const b = 41

func main() {
	x = 12
	result := x + y

	fmt.Println("The result is: ", result)

	// Course's solution
	fmt.Printf("the value of a is %v and the type of a is %T\n", a, a)
	fmt.Printf("the value of b is %v and the type of b is %T\n", b, b)

	z := 42
	fmt.Printf("the value of z is %v and the type of z is %T\n", z, z)
}
