package main

import (
	"github.com/mytkdals93/headfirst_go/struct/4/magazine"
)

func main() {
	s1 := magazine.DefaultSubscriber("Aman Sigh")
	magazine.ApplyDiscount(s1)
	magazine.Printlnfo(s1)

	s2 := magazine.DefaultSubscriber("Beth Ryan")
	magazine.Printlnfo(s2)

}
