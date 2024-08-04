package readers

import (
	"fmt"
	"io"
	"strings"
)

/*

the io package specifies the io.Reader interface, which represents the read end of a  stream data
the io.Reader interface has a Read method

*/

func Readers() {

	r := strings.NewReader("hello, world")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err=%v b=%v\n", n, err, b)
		if err == io.EOF {
			break
		}
	}
}
