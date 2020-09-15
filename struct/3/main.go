package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printlnfo(s *subscriber) {
	fmt.Println("Name: ", s.name)
	fmt.Println("Rate: ", s.rate)
	fmt.Println("Active: ", s.active)
}
func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}
func main() {
	s1 := defaultSubscriber("Aman Sigh")
	applyDiscount(s1)
	printlnfo(s1)

	s2 := defaultSubscriber("Beth Ryan")
	printlnfo(s2)

}
