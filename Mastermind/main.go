package main

import (
	"fmt"
)

const anzStellen int = 4

func main() {
	codearr := [anzStellen]string{"B", "A", "B", "B"}
	versuch1 := [anzStellen]string{"A", "B", "C", "A"}

	fmt.Println(bewertung(versuch1, codearr))

}
func bewertung(versuch, code [anzStellen]string) (int, int) {
	var schwarzeStifte int
	var weisseStifte int
	var index0 int
	var index1 int
	var index2 int
	var index3 int

	for i := 0; i < len(versuch); i++ {
		switch {
		case versuch[i] == code[i]:
			schwarzeStifte++
		case versuch[i] == code[0] && versuch[0] != code[0] && index0 == 0:
			weisseStifte++
			index0 = 1
		case versuch[i] == code[1] && versuch[1] != code[1] && index1 == 0:
			weisseStifte++
			index1 = 1
		case versuch[i] == code[2] && versuch[2] != code[2] && index2 == 0:
			weisseStifte++
			index2 = 1
		case versuch[i] == code[3] && versuch[3] != code[3] && index3 == 0:
			weisseStifte++
			index3 = 1
		}

	}
	return schwarzeStifte, weisseStifte
}

