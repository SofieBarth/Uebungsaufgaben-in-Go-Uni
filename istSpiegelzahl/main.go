package main

import "fmt"

var neueZahl int

func main() {
	fmt.Println(spiegelzahl(51679))

}
func spiegelzahl(zahl int) int {

	if zahl >= 1 {
		neueZahl = neueZahl*10 + zahl%10
		return spiegelzahl(zahl / 10)
	}
	return neueZahl
}
