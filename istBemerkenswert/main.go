package main

import (
	"fmt"
)

func main() {

	fmt.Println(bemerkenswert(2))
}

func bemerkenswert(start int) int {
	result := 0
	number := 1

	for i := start + 1; result != number*number; i++ {
		result = 0
		number = i

		for j := 1; result < number*number; j++ {
			result += j
		}
	}
	return number
}
