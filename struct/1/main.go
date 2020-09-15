package main

import "fmt"

//구조체 정의
type part struct {
	description string
	count       int
}
type cart struct {
	name     string
	topSpeed float64
}
type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	// 구조체 타입 정의
	var subscriber1 struct {
		name   string
		rate   float64
		active bool
	}
	//=> 매번 타입을 다시 정의해 줘야함
	var subscriber2 struct {
		name   string
		rate   float64
		active bool
	}

	subscriber1.name = "상민"
	subscriber1.rate = 4.99
	subscriber1.active = true
	fmt.Printf("%#v", subscriber1)
	subscriber2.name = "상민2"
	subscriber2.rate = 3.99
	subscriber2.active = false
	fmt.Printf("%#v", subscriber2)
	//패키지 레벨에 정의된 subscriber타입 사용
	var s3 subscriber
	s3.name = "상민3"
	s3.rate = 4.99
	s3.active = true
	fmt.Printf("%#v", s3)
}
