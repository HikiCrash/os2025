package main

import "week10/pkg/greeting"

func main() {
	greeting.Hello()
	greeting.Hi()
}

// func hi()가 소문자로 시작해서 패키지 외부에서 접근 불가
// 외부에서 가져오는 패키지의 함수는 대문자로 시작해야 접근 가능
