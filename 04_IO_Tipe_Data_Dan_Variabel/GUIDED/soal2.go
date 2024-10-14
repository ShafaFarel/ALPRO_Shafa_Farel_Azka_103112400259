package main

import (
	"fmt"
)

func main() {

	var (
		masukan, angka1, angka2, angka3 int
	)

	fmt.Print("Masukkan 3 digit angka: ")
	fmt.Scan(&masukan)

	angka1 = masukan / 100
	angka2 = (masukan % 100) / 10
	angka3 = masukan % 10

	fmt.Println(angka1 <= angka2 && angka2 <= angka3)

}