package main

import "fmt"

func main() {
	var bilanganPertama, bilanganKedua, hasil int

	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scan(&bilanganPertama)

	fmt.Print("Masukkan bilangan kedua: ")
	fmt.Scan(&bilanganKedua)

	// Menghitung perkalian menggunakan penjumlahan berulang
	for i := 0; i < bilanganKedua; i++ {
		hasil += bilanganPertama
	}

	fmt.Println("Hasil perkalian adalah:", hasil)
}