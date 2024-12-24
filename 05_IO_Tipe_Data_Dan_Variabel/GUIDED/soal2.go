package main

import "fmt"

func main() {

	var n int
	var Alas int
	var Tinggi int

	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)
	fmt.Println("------------------------------")

	for i := 1; i <= n; i++ {
	fmt.Print("Masukkan alas: ")
	fmt.Scan(&Alas)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&Tinggi)

	luas := Alas * Tinggi/2
	fmt.Println("------------------------------")
	fmt.Println("Luas segitiga: ", luas)
	}
}