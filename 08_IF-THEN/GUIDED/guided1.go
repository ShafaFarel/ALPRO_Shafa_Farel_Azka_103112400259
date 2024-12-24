package main

import "fmt"

func main() {

	var bilanganBulat int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilanganBulat)

	if bilanganBulat < 0 {
		bilanganBulat = -bilanganBulat
	}

	fmt.Println("Nilai mutlaknya adalah: ",bilanganBulat)
}