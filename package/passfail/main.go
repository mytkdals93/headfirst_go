package main

import (
	"fmt"
	"log"

	"github.com/mytkdals93/headfirst_go/package/keyboard"
)

func main() {
	fmt.Println("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Printf("Agrade of %v, is %v. \n", status, grade)
}
