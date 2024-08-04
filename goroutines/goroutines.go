package goroutines

import (
	"fmt"
	"time"
)

/*
a goruotine  ia a lightweight thread managed by the go runtime
channels are typesd conduit through which you can send and receive values wiyh the channel operator <-

channels can be buffered.provide the buffer length as the second argument to make to initialize a buffered channel

		a sender can close the channel tom check we need to check byu addding a second parameter

	 a select statement lets a goroutine wait on multiple communication operations. a select blocks until one of its cases can run.then it executes that case.
*/
func Gorutines() {
	c := make(chan string)
	go say("hello", c)
	x := <-c

	// select {
	// case x, ok := <-c:

	// 	if ok {
	// 		fmt.Println(ok, x)
	// 	}
	// case y := <-c:
	// 	fmt.Println(y)

	// default:
	// }
	// if ok {
	// 	fmt.Println(ok)
	// }
	fmt.Println(x)
	fmt.Println("goroutine")
}

func say(s string, c chan string) {
	time.Sleep(2000)
	for i := 0; i < 5; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(i)
	}
	fmt.Println(s)

	c <- s
}
