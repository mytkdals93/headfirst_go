package main

import (
	"fmt"

	"github.com/mytkdals93/headfirst_go/struct/6/magazine"
)

func main() {
	address := magazine.Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "68111",
	}

	s1 := magazine.DefaultSubscriber("Aman Sigh")
	s1.Address = address
	magazine.ApplyDiscount(s1)
	magazine.Printlnfo(s1)

	s2 := magazine.DefaultSubscriber("Beth Ryan")
	s2.Address = address
	magazine.Printlnfo(s2)
	//마치 Address가 없는 것처럼 사용
	fmt.Println(s2.City)

}
