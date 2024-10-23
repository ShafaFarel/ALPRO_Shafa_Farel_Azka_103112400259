package main

import (
	"fmt"
)

func main() {

	var x  int
	var y int
	var hasil int

	fmt.Print("Masukkan dua bilangan:")
	fmt.Scan(&x, &y)

	if x > y {
		fmt.Println("angka pertama harus lebih kecil dari angka kedua")
		return
	} 

	for i := x; i <= y; i++ {
		hasil += i
	}
	fmt.Println("Hasil penjumlahan:", hasil)
}