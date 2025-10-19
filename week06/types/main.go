package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// var name string,int,float64 > 0
	// var a bool > False, (zero value.)

	// var name string
	// name = "Kim"

	// var name = "Kim"

	name := "Kim"
	fmt.Println(math.Round(2.71))
	fmt.Println(strings.Title("gohead first go"))
	fmt.Println(name, reflect.TypeOf(name))

	// *변수 이름은 숫자로 시작하지 못함.

}
