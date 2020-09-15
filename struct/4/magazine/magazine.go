package magazine

import "fmt"

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func Printlnfo(s *Subscriber) {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Rate: ", s.Rate)
	fmt.Println("Active: ", s.Active)
}
func DefaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func ApplyDiscount(s *Subscriber) {
	s.Rate = 4.99
}
