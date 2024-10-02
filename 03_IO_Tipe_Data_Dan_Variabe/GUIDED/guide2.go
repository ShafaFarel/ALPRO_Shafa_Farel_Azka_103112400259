package main

import (
	"fmt"
)

func main() {
	var alas, tinggi, hasil float32

	fmt.Print("Masukan Panjang alas segitiga: ")
	fmt.Scan(&alas)
	fmt.Print("Masukan Panjang alas segitiga: ")
	fmt.Scan(&tinggi)


	hasil = alas * tinggi / 0.5 //Rumus luas segitiga
	fmt.Println("Luas Segitiga Adalah = ", hasil)
}
