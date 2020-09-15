package magazine

import "fmt"

type Subscriber struct {
	Name    string
	Rate    float64
	Active  bool
	Address //익명구조체
}
type Employee struct {
	Name   string
	Salray string
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

func Printlnfo(s *Subscriber) {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Rate: ", s.Rate)
	fmt.Println("Active: ", s.Active)
	fmt.Println("Address: ", s.Address)
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
