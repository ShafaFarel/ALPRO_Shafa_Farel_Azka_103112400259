package main

import (
	"fmt"
)

func main() {

	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

// Codingan di atas mengharuskan user agar memasukan 3 kata berbeda dan akan disimpan dengan menggunakan tipe data string.
// Setelah user memasukan 3 kata berbeda, maka sistem akan otomatis mengacak kata yang telah dimasukan. Kata kedua akan pada posisi pertama, kata ketiga akan menjadi posisi kedua dan kata pertama akan menjadi di posisi ketiga.