// 1
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	// casting
// 	var length float64 = 3.2
// 	var width int = 2
// 	fmt.Println("면적은", int(length)*width)
// 	fmt.Println("길이 > 너비?", int(length) > width)
// 	fmt.Println("형변환", reflect.TypeOf(int(length)))
// 	fmt.Println("원본", reflect.TypeOf(length))
// }

// 2
// package main

// import (
// 	"fmt"
// 	"strings"
// 	"time"
// )

// func main() {
// 	var now time.Time = time.Now()
// 	var day int = now.Day()
// 	fmt.Println(day)
// 	univ := "Go$ Code$"
// 	changer := strings.NewReplacer("$", "!")
// 	changed := changer.Replace(univ)
// 	fmt.Println(changed)
// }

// 3
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	r := bufio.NewReader(os.Stdin)
// 	i, _ := r.ReadString('\n') // ignore error
// 	fmt.Println(i)
// }

// 4
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"
// )

// func main() {
// 	var now time.Time = time.Now()
// 	var month time.Month = now.Month() // month := now.Month()
// 	fmt.Println(month)

// 	r := bufio.NewReader(os.Stdin)
// 	i, err := r.ReadString('\n')
// 	//fmt.Println(err)
// 	log.Fatal(err) // report the error and exit the program
// 	fmt.Println(i)
// }

// 5
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// shadowing
	// var fmt string = "code"
	// var int int = 7
	// var k int = 11
	// fmt.Println(int)

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	//fmt.Println(err)
	if err != nil {
		log.Fatal(err) // report the error and exit the program
	}

	i = strings.TrimSpace(i)
	score, err := strconv.ParseFloat(i, 64)
	if err != nil {
		log.Fatal(err)
	}

	if score >= 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	fmt.Println(score)
}
