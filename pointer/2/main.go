package main

import "fmt"

// func double(number int) {
// 	//값이 복사되어 전달되기 때문에 원래 값은 변하지 않음
// 	number *= 2
// }
func double(number *int) {
	//포인터 형 변수도 물론 복사되지만 복사된 값도 역시 같은 변수 주소를 가지고 있기 때문에 값이 변함
	*number *= 2
}

func main() {
	number := 6
	double(&number)
	fmt.Println(number)
}
