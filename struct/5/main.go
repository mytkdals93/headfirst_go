package main

import (
	"github.com/mytkdals93/headfirst_go/struct/5/magazine"
)

func main() {
	address := magazine.Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "68111",
	}

	s1 := magazine.DefaultSubscriber("Aman Sigh")
	s1.HomeAddress = address
	magazine.ApplyDiscount(s1)
	magazine.Printlnfo(s1)

	s2 := magazine.DefaultSubscriber("Beth Ryan")
	s2.HomeAddress = address
	magazine.Printlnfo(s2)

}
