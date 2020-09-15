package main

import (
	"fmt"
	"log"

	"github.com/mytkdals93/headfirst_go/package/keyboard"
)

func main() {
	fmt.Println("Enter a temrature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsious := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius \n", celsious)
}
