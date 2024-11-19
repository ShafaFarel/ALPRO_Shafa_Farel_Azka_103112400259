package main

import (
	"fmt"
)

func main() {

	// Variabel untuk menyimpan huruf
	var huruf string

	// Inputan untuk memasukan huruf
	fmt.Print("Masukan huruf: ")
	fmt.Scan(&huruf)

	// Percabangan menentukan dimana ketika memasukan AIUEO, maka akan muncul vokal, ketika memasukan huruf selain itu akna muncul hurufkonsonan
	if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" || huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Println("Huruf Vokal")
	} else {
		fmt.Println("Huruf Konsonan")
	}
}