package main

import (
	"fmt"
	"log"

	"github.com/mytkdals93/headfirst_go/array/average"

	"github.com/mytkdals93/headfirst_go/array/readfile"
)

func main() {
	var myArr [3]int
	myArr = [3]int{1, 2, 3}
	for i := 0; i < len(myArr); i++ {
		fmt.Printf("%d: %d\n", i, myArr[i])
	}
	fmt.Println("-----------------------------")
	for i, v := range myArr {
		fmt.Printf("%d: %d\n", i, v)
	}
	fmt.Println("-----------------------------")
	myArr2 := [2]string{"Hello", "world"}
	fmt.Printf("%#v\n", myArr2)
	fmt.Println(myArr2)
	fmt.Println("-----------------------------")
	numbers, err := readfile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(numbers)
	result := average.Calc(numbers)
	fmt.Println(result)

}
