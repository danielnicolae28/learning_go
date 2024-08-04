package pointers

import "fmt"

/*

 */

func Pointer() {

	fmt.Println("Pointer lesson")
	var p *int
	i := 40
	p = &i // read i through the pointer p

	fmt.Println(*p)
	*p = 21 // set i through the pointer p

	fmt.Println(i, *p)

}
