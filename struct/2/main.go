package main

import "fmt"

type subscriber struct {
	name   string
	rate   float32
	active bool
}

func changeName(s subscriber) {
	//복사된 값이 인자로 전달되기 때문에 원본 값이 변하지않음.
	s.name = "이상민"
}
func applyDiscount(s *subscriber) {
	//Go에서 *s.rate는 (s.rate)를 포인터로 인식하며 에러를 발생시킨다.
	//그래서 (*s).rate로 접근해야한다.
	(*s).rate = 4.99
	//하지만 이건 너무 번거롭기 때문에 생략하고 그냥 아래와 같이 입력해도 정상작동한다.
	s.rate = 4.99
}

func showInfo(s subscriber) {
	fmt.Println(s)
}

func main() {
	var s1 subscriber
	s1.name = "상민"
	s1.rate = 5.0
	s1.active = true

	changeName(s1)
	applyDiscount(&s1)

	showInfo(s1)
}
