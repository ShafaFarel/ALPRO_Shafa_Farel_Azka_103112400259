package main

import "fmt"

func main() {

	var angka int

	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&angka)

	digitCount := 0

	for angka > 0 {
		angka /= 10
		digitCount++
	}

	fmt.Println(digitCount)
}