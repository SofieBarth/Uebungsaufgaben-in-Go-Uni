package main

import (
	"fmt"
)

const anzZeilen int = 6
const anzSpalten int = 5

func main() {

	matrix := [anzZeilen][anzSpalten]int{
		{15, 8, 1, 24, 17},
		{16, 14, 7, 5, 23},
		{22, 20, 13, 6, 4},
		{3, 21, 19, 12, 10},
		{9, 2, 25, 18, 11},
		{27, -6, 13, 31, -1},
	}
	fmt.Println(matrix)
	fmt.Println(alternierendeSpaltensumme(&matrix))
	fmt.Println(matrix)

}
func alternierendeSpaltensumme(arr *[anzZeilen][anzSpalten]int) int {
	var alternSumme int
	var FehlerAnzahl int

	for j := 0; j < anzSpalten; j++ {
		for i := 0; i < anzZeilen-1; i++ {
			if (i % 2) == 0 {
				alternSumme += arr[i][j]
			} else {
				alternSumme -= arr[i][j]
			}
		}
		if alternSumme != arr[anzZeilen-1][j] {
			arr[anzZeilen-1][j] = alternSumme
			FehlerAnzahl += 1
		}
		alternSumme = 0
	}
	return FehlerAnzahl
}
