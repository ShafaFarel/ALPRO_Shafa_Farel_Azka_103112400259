package main

import (
	"fmt"
)

func main() {

	var x  int
	var y int

	fmt.Print("Masukkan dua bilangan:")
	fmt.Scan(&x, &y)

	fmt.Println("Hasil penjumlahan:", x+y)
}