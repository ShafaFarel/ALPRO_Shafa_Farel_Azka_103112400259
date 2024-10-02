package main

import (
	"fmt"
	"math"
)

func main() {

	var r float64

	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scan(&r)

	var operasi float64 = 3.14 * r * r
	hasil := math.Round(operasi*10) / 10
	fmt.Println("Hasil dari luas lingkaran dengan jari-jari:", r, "adalah", hasil)
}