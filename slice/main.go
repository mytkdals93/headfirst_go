package main

import "fmt"

func main() {
	//slice 변수 선언
	var notes []string
	//슬라이스 생성후 대입
	notes = make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "m1"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	//선언없이 할당, 타입은 추론
	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])

	fruits := []string{"apple", "banana", "orange", "grape"}
	for _, v := range fruits {
		fmt.Println(v)
	}

}
