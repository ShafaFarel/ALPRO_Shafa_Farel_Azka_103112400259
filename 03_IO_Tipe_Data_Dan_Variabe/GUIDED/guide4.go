package main

import (
	"fmt"
)

func main() {
	var rupiah int
	var hasil int

	
	fmt.Scan(&rupiah)

	hasil = rupiah / 150000
	fmt.Print("Konversi Dolar = ", hasil)
}
