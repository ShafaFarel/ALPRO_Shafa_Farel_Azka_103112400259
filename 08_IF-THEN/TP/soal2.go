package main

import "fmt"

func main() {

	var angka int

	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&angka)

	if angka % 2 == 0 {
		fmt.Println("Angka adalah Genap")
	} else {
		fmt.Println("Angka adalah Ganjil")
	}
}