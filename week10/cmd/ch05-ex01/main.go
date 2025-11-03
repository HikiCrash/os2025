package main

import (
	"fmt"
)

func main() {
	arrayBool := [3]bool{true, false}
	arrayInt := [3]int{-9, 11, 7}
	for i := 0; i < len(arrayInt); i++ {
		fmt.Println(i, arrayInt[i])
		// fmt.Println(i, arrayBool[i]) // runtime error
	}

	fmt.Printf("%#v\n", arrayBool)
	fmt.Printf("%#v\n", arrayInt)
}
