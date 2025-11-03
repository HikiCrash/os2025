// package main

// import "greeting"

// func main() {
// 	greeting.Hello()
// 	greeting.Hi()
// }

// func hi()가 소문자로 시작해서 패키지 외부에서 접근 불가
// 외부에서 가져오는 패키지의 함수는 대문자로 시작해야 접근 가능

package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("점수 입력 : ")
	score, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	if score >= 60 {
		fmt.Printf("%.1f점은 합격!", score)
	} else {
		fmt.Printf("%.1f점은 불합격...", score)
	}
}
