package main

import (
	"fmt"
	"math"
)

const anzOrte = 4

func main() {

	entfernungsmatrix := [anzOrte][anzOrte]int{
		{0, 7, 4},
		{7, 0, 6},
		{4, 6, 0},
	}
	fmt.Println(rundreise(entfernungsmatrix))
}

func rundreise(entfernung [anzOrte][anzOrte]int) ([anzOrte + 1]int, int) {
	var reiseroute [anzOrte + 1]int
	reiseroute[0] = 1
	var indexReturnArray = 1
	var nextOrt int
	var nextRow int
	var strecke int
	var boolMatrix [anzOrte][anzOrte]bool

	for i := range boolMatrix {
		for j := range boolMatrix[i] {
			if j == 0 {
				boolMatrix[i][j] = true
			} else {
				boolMatrix[i][j] = false
			}
		}
	}
	fmt.Println(boolMatrix)

	for indexReturnArray < len(entfernung) {
		i := nextRow
		//nextOrt = entfernung[i][0]
		shortestDistance := math.MaxInt

		for j := range entfernung[i] {
			//fmt.Printf("iteration %d, entfernung: %d; boolmatrix: %t; nextOrt: %d; true?: %t \n", j, entfernung[i][j], boolMatrix[i][j], nextOrt, (entfernung[i][j] < shortestDistance))
			if (boolMatrix[i][j] == false) && (entfernung[i][j] < shortestDistance) {
				//	fmt.Printf("true \n")
				shortestDistance = entfernung[i][j]
				nextOrt = j
			}
		}
		for k := range boolMatrix {
			boolMatrix[k][nextOrt] = true
		}
		fmt.Println(boolMatrix)
		reiseroute[indexReturnArray] = nextOrt + 1
		indexReturnArray++
		strecke += entfernung[i][nextOrt]
		nextRow = nextOrt
	}
	reiseroute[indexReturnArray] = 1
	strecke += entfernung[nextOrt][0]

	return reiseroute, strecke
}
