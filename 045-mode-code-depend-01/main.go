package main

import (
	"fmt"

	"github.com/nicolas-rohricht/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

	//also like this
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
}
