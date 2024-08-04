package structs

import "fmt"

// a struct is a collection of fields
// struct fields are accesed using a dot
// struct fields can be accesed througha a struct pointer

type Vertex struct {
	X int
	Y int

	//X,Y int //struct literal
}

/// to access the field X of a struct when we have the struct pointer P we could write (*p).X. That notation is cumbersome, so the language permis us instead to write just p.X, without the explicit dereference.

type human struct {
	firstName string
	lastName  string
	age       int
	job       string
}

func StructsFunction() {
	fmt.Println("Srtucts lesson")

	v := Vertex{1, 2}
	p := &v
	p.X = 129

	fmt.Println(v)

	person := human{"Daniel", "Ioana", 24, "Engineer"}

	fmt.Println(person.firstName, person.lastName)

	person2 := &person

	person2.age = 20
	fmt.Printf("Value address: %p\n", &person)
	fmt.Printf("Pointers address: %v\n", &person2)

}

// 0xc00005a030 this add
