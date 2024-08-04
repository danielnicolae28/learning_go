package intefaces

import "fmt"

/*

a interface type is defined as a set of method signatures. A value of interface type can hold any value that implements those methods.
if the concreate value inside the interface is nil, the method will be called with a nil receiver
the interface type that specifies zero methods is known as the empty interface{}


type assertion
provides acces to an interface value underlying concrete value
to test whether an interface value holds a specific type, a type assertion can return two value that reports whether the assertion succeeded

t, ok := i.(T)



type switch is a construct that permits several type assertions in series
is tha same as a normal switch but the cases in a type switch specify types not values

Stringers

one of the most umbiquitous interface is Stringer defined by the fmt package

*/
type Stringer interface {
	/// Is a type that can be described itself as a string. The fmt package look for this interface to print values.
}

type I interface {
	M()
}

type T struct {
	s string
}

func (t T) M() {
	fmt.Println(t.s)
}

func Interfaces() {

	var a interface{}
	fmt.Println(a)
	var i I = T{"hello"}
	t, ok := i.(T)

	fmt.Println(t, ok)

	describe(2)
	i.M()

}

func describe(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Printf("%v has int type\n", v)
	case string:
		fmt.Printf("%v has string type", v)
	default:
		fmt.Println("don't know the type")
	}
	fmt.Printf("(%v, %T)\n", i, i)
}
