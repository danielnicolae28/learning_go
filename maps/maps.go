package maps

import (
	"fmt"
	"strings"
)

/*
A map maps keys to values
the zero value of a map is nil. A nil map han no keys, nor can keys be added

*/

type Vertex struct {
	lat, long float64
}

var m map[string]Vertex

func Maps() {

	fmt.Println("lesson maps")
	m = map[string]Vertex{"bell labs": Vertex{40.5234532, -65.2342},
		"google": Vertex{65.653453, -55.245},
	}
	// m["bell Labs"] = Vertex{
	// 	40.0953443, -53.432525,
	// }
	dataGoogle := m["google"]

	fmt.Println(m)
	fmt.Println(dataGoogle)

	delete(m, "google")
	fmt.Println(m)
	dataGoogle, ok := m["google"] // checkk if data is available in the map ok is bool type and return true or false
	fmt.Println(ok, dataGoogle)
	WordCount()
}

// "I am learning Go!"
func WordCount() {
	text := "I am learning Go!"

	fmt.Println(strings.Split(text, " "))
}
