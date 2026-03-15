package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(buchstabenSortieren("mnbvcxylkjhgfdsapoiuztrewq"))

}
func buchstabenSortieren(word string) string {
	word = strings.ToLower(word)
	result := word[0:1]

	for i := 1; i < len(word); i++ {
		nextLetter := word[i : i+1]
		lastLetterIndex := len(result)

		if nextLetter <= result[0:1] {
			result = nextLetter + result
		} else if nextLetter >= result[lastLetterIndex-1:lastLetterIndex] {
			result = result + nextLetter
		} else {
			index := 0

			for j := 0; nextLetter >= result[j:j+1]; j++ {
				index = j + 1
			}
			resultpart1 := result[:index]
			resultpart2 := result[index:]

			result = resultpart1 + nextLetter + resultpart2
		}
	}
	return result
}
