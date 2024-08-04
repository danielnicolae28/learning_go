package stringer

import "fmt"

type IPAddr [4]byte

// {127, 0, 0, 1}, input
// "127.0.0.1",  output

// TODO: Add a "String() string" method to IPAddr.

func (addr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", addr[0], addr[1], addr[2], addr[3])
}

func Stringer() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for _, ip := range hosts {
		fmt.Printf("%v\n", ip)
	}
}
