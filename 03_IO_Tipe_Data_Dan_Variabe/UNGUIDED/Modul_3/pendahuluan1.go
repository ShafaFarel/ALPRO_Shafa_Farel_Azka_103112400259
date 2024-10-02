package main

import (
	"fmt"
)

func main() {

	//Masukan sebuah panjang sisi persegi.
	sisi := 27

	keliling := 4 * sisi //Rumus yang digunakan untuk menemukan keliling persegi
	luas := sisi * sisi  //Rumus yang digunakan untuk menghitung luas dari sebuah persegi

	// Output yang digunakan untuk menampilkan sebuah hasil perhitungan keliling dan luas persegi.
	fmt.Print("Keliling dari sebuah persegi adalah: ", keliling, " Meter")
	fmt.Print(" dan luas dari sebuah persegi adalah: ", luas, " Meter Persegi")
}