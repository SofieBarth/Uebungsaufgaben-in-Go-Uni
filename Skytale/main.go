package main

import (
	"fmt"
)

func main() {

	fmt.Println(encrypt("EINEGEHEIMEBOTSCHAFT", 5))
	fmt.Println(decrypt("EGIOHIEMTANHESFEEBCT", 5))

	//fmt.Println(encrypt("LUCASIOISTEINESCHLAFMUETZE", 5))
	//fmt.Println(decrypt("LIECMUOIHUCINLEASEATSTSFZ", 5))
}

func encrypt(message string, skytale int) string {

	var cryptomatrix [100][100]string
	var crypticmessage string
	var rowlength int = len(message) / skytale
	var messageIndex int = -1

	for i := 0; i < len(cryptomatrix); i++ {
		if i < skytale {
			for j := 0; j < rowlength; j++ {
				messageIndex++
				cryptomatrix[i][j] = string(message[messageIndex])
			}
		}
	}
	for j := 0; j < len(cryptomatrix[0]); j++ {
		if j < rowlength {
			for i := 0; i < skytale; i++ {
				crypticmessage += cryptomatrix[i][j]
			}
		}
	}
	fmt.Println(cryptomatrix)
	return crypticmessage
}
func decrypt(crypticmessage string, skytale int) string {

	var cryptomatrix [100][100]string
	var decryptmessage string
	var rowlength int = len(crypticmessage) / skytale
	var messageIndex int = -1

	for j := 0; j < len(cryptomatrix[0]); j++ {
		if j < rowlength {
			for i := 0; i < skytale; i++ {
				messageIndex++
				cryptomatrix[i][j] = string(crypticmessage[messageIndex])
			}
		}
	}
	for i := 0; i < len(cryptomatrix); i++ {
		if i < skytale {
			for j := 0; j < rowlength; j++ {
				messageIndex++
				decryptmessage += cryptomatrix[i][j]
			}
		}
	}
	fmt.Println(cryptomatrix)
	return decryptmessage
}
