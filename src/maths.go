package main

import (
	"fmt"
	//"math"
)

func maths() {
	var num int = 1
	fmt.Println(num)

	for i := 1; i <= 8; i++ {
		fmt.Println(i)
	}

	num2 := 7
	if num2%2 == 0 {
		fmt.Println(num2, " is even")
	} else {
		fmt.Println(num2, " is off")
	}
}
