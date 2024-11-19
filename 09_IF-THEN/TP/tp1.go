package main

import (
	"fmt"
)

func main() {

	// Variabel untuk menyimpan nilai
	var nilai int

	// Inputan untuk memasukan nilai
	fmt.Print("Masukan nilamu: ")
	fmt.Scan(&nilai)

	// Percabangan yang digunakan untuk menentukan index dari nilai yang dimasukan
	if nilai > 90 {
		fmt.Println("Nilai anda A")
	} else if nilai >= 80 && nilai <= 90 {
		fmt.Println("Nilai anda AB")
	} else if nilai >= 70 && nilai < 80 {
		fmt.Println("Nilai anda B")
	} else if nilai < 70 {
		fmt.Println("Nilai anda C")
	}
}