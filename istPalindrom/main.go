package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(istPalindromiter("Otto"))
	fmt.Println(istPalindrom("sofieeifoks"))
}
func istPalindrom(zeichenfolge string) bool {
	lowZeichenfolge := strings.ToLower(zeichenfolge)

	if len(lowZeichenfolge) < 2 {
		return true
	} else if lowZeichenfolge[0:1] == lowZeichenfolge[len(lowZeichenfolge)-1:len(lowZeichenfolge)] {
		lowZeichenfolge = lowZeichenfolge[1 : len(lowZeichenfolge)-1]

		return istPalindrom(lowZeichenfolge)
	} else {
		return false
	}
}
