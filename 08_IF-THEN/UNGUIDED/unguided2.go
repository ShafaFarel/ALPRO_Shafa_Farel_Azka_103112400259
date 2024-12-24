package main

import "fmt"

func main () {

	var bilanganBulat int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilanganBulat)

	if bilanganBulat < 0 && bilanganBulat%2 == 0 {
		fmt.Print("genap negatif")
	} else {
		fmt.Print("bukan genap negatif")
	}
}