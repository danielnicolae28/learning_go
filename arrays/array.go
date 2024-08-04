package arrays

import "fmt"

/*
the type [n]T is an array on n values of type T

an array length is part of its type, so arrays cannot be resized.
this seems limiting, but don't worry; Go provides a convenient way of working with arrays

*/

func Array() {

	lesson := "Arrays"

	fmt.Printf("Lessons about %v\n", lesson)

	var a [10]int // array of ten integers

	fmt.Println(a)
	a[5] = 15
	fmt.Println(a)

}
