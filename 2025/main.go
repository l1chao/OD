package main

import (
	"fmt"
	"strings"
)

func Ans4() {
	var x string
	fmt.Scanln(&x)
	list := strings.Split(x, ",")
	fmt.Println(list)

	fmt.Println(x)
	fmt.Scanln(&x)
	fmt.Println(x)
	fmt.Scanln(&x)
	fmt.Println(x)
}

func main() {
	// x, y := 0, 0
	// // 12k33
	// fmt.Scanf("%dc%d\n", &x, &y)
	// fmt.Println(x, y)

	// //k12c33
	// fmt.Scanf("b%dc%d", &x, &y)
	// fmt.Println(x, y)

	Ans6()
}
