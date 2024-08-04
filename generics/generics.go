package generics

import "fmt"

/*
go functions can be written to work on multiple types using type parameters. type parameters appear between brackets, before the function arguments.


comparable is a useful constraint that makes it possible to use the == and != operators on values of type
generic funcitons also supports generic types. A type can be parameterized with a type parameter which can be useful for implementing generic data structures
*/

type List[T int] struct {
	val  T
	next *List[T]
}

func Generics() {

	fmt.Println(Index([]int{1, 2, 34}, 1))
	genericTypes()
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func genericTypes() {
	b := List[int]{2, &List[int]{}}
	fmt.Println(b)
}
