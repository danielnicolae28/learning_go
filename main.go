package main

import (
	"fmt"
	"learningtour/arrays"
	"learningtour/errors"
	"learningtour/functions"
	"learningtour/generics"
	"learningtour/goroutines"
	"learningtour/intefaces"
	"learningtour/maps"
	"learningtour/methods"
	"learningtour/pointers"
	"learningtour/readers"
	"learningtour/slices"
	"learningtour/stringer"
	"learningtour/structs"
)

func main() {
	fmt.Println("A tour of go")
	pointers.Pointer()
	structs.StructsFunction()
	arrays.Array()
	slices.Slice()
	maps.Maps()
	functions.Functions()
	methods.Methods()
	intefaces.Interfaces()
	stringer.Stringer()
	errors.Errors()
	readers.Readers()
	generics.Generics()
	goroutines.Gorutines()
}
