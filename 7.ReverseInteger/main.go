package main

import (
	"fmt"
)

func main() {
	//num := 1534236469
	num := -1534236469
	fmt.Println(reverseInt(num))

}

func reverseInt(x int) int {
	var reversed int
	var sign int = 1
	if x < 0 {
		sign = -1
	}
	for (sign * x) > 0 {

		remainder := x % 10
		reversed = reversed*10 + remainder
		x = x / 10
	}

	if (sign * reversed) > 2147483648 {
		return 0
	}
	return reversed
}
