package main

import "fmt"

type subscriber struct {
	name  string
	price float32
}

func applyPrice(s *subscriber) {
	s.price = 10000
	s.name = "Park"
}

func main() {
	var s1 subscriber
	var p *subscriber = &s1

	applyPrice(&s1)
	fmt.Println(s1.name, s1.price)
	// fmt.Println(*p.price) 에러
	fmt.Println((*p).price)
	fmt.Println(p.price)
}
