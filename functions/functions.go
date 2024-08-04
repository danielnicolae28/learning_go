package functions

import (
	"fmt"
	"math"
)

/*
functions are values too. They can be passed around just like other values.
function values may be used as function arguments and return values.

go functions may be closures. A closure is a function value that references variables form outside its body. The function may access and assign to the referenced variables. In this sense the function is bound to the variable.
*/

func Functions() {
	fmt.Println("Lesson functions")

	hypot := func(x, y float64) float64 { return math.Sqrt(x*x + y*y) }
	fmt.Println(hypot(3, 6))

	fmt.Println(compute(hypot))

}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)

}
