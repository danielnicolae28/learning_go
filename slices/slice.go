package slices

import "fmt"

/*

if an  array has fixed size. A slice is dynamically-sized, flexible view into the elements of an array. IN PRACTICE SLICES ARE MORE COMMON THAN ARRAYS.

The type []T is a slice with elements if type T.

A slice is formed by specifying two indices, low and a high separated by a colon
A slice does not store any data, it just describes a section of an underlying array
changing the elements of a slice modifies the corresponding elements of its  underlying array

other slices that share the same underlyig array will see those changes

a slice has both  LENGTH AND CAPACITY :

length :  the number of elements it contains
capacity : the number of elements in the undelying array,counting from the first element in the slice

length and capacity can be obtained using the expressions len(s) and cap(s)


the zero value of a slice is nil
nil length  and capacity is 0 and has no undelying array

*/

func Slice() {
	fmt.Println("Slice lessons")

	primes := [6]int{2, 3, 5, 7, 11, 14}
	numsArr := [10]int{10, 22, 32, 53, 129, 232, 78, 9, 3, 10}
	var s []int = primes[1:4] // a slice from an array whith low and high index
	p := numsArr[:5]
	a := numsArr[:] // get all data from array
	fmt.Println(s)

	s[2] = 555 // if the slice is changed than the original array is changed and the other slices that use that value like a pointer that changes a value
	fmt.Println(primes)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(p) // when slicing high and lows can be omited and use their defaults
	sliceTwo(p)
	makeSlice()
	appendToSlice(s)
}

func sliceTwo(s []int) {

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}

/*

slices can be created with the make function; this is how you create dynamically-sized arrays

the make function allocates a zeroed array and returns a slice that referes to that array
 slices can contain any type, including other slices


 the range from the for loop iterates over a slice or map
 when ranging two values are returned.First is the index and the second is a copy of the element at that index

 reasons to use a pointer receiver:
 - method can modify the value that its receiver points to
 - not copying the value on each method call(more efficient if the receiver is a large struct)
 all methods on a given type should have either value or poionter receivers, not a mixture of both.

*/

func makeSlice() {

	//ex:  datatype, len, cap
	a := make([]int, 5, 10) // len(a)=5, cap(a) = 10

	fmt.Println(a)
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

}

func appendToSlice(s []int) {
	fmt.Println(s)
	s = append(s, 5)
	fmt.Printf("%v\n", s)

	for i, v := range s {
		fmt.Printf("i=%d,v=%d \n", i, v)
	}
	//we can skip the index or the value using _ or if you want only the index ou can omit the value
	// for _, v := range s {
	// 	fmt.Printf("i=%d,v=%d \n", i, v)
	// }

}
