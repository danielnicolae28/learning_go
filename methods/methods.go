package methods

import (
	"fmt"
	"math"
)

/*

go does not have calsses but we can define methods of types

a method is a function with a special receiover argument

the receiver appears in its own argument list between the func keyword and the method name
a method is just a function with a receiver argument
we can declare a method on non struct types too.
methods and pointer indirection. methods with pointer receivers take either a value or a pointer as the receiver when they are called. Functions with a pointer argument must take a pointer
functions that take a value argument must take a value of that specific type
*/

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) Abs2() float64 {

	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y + v.Y)
}

func (v *Vertex) Scale(f float64) {

	v.X = v.X * f
	v.Y = v.Y * f
}

func Methods() {
	fmt.Println("Methods")
	v := Vertex{3, 4}
	f := MyFloat(-math.Sqrt2)
	v.Scale(10)
	fmt.Println(v.Abs())
	fmt.Println(f.Abs2())
}
