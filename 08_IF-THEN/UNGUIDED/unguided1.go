package main

import "fmt"

func main() {

	var jumlahOrang int

	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scan(&jumlahOrang)

	if jumlahOrang%2 == 0{

		jumlahOrang = jumlahOrang / 2
	} else {
		jumlahOrang = (jumlahOrang + 1) / 2
	}

	fmt.Print("Jumlah motor yang dibutuhkan: ", jumlahOrang)
	
}