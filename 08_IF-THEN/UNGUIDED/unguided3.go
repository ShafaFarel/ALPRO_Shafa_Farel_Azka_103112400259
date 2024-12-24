package main

import (
	"fmt"
)

func main() {
	var x, y int
	var hasilX, hasilY bool

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan pertama (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan kedua (y): ")
	fmt.Scan(&y)

	// Mengecek apakah x adalah faktor dari y
	
	if y%x == 0 {
		hasilX = true
	} 

	if x%y == 0 {
		hasilY = true
	} 

	fmt.Println(hasilX)
	fmt.Println(hasilY)
}