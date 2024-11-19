package main

import "fmt"

func main() {

	// Membuat variabel umur dan kartuKeluarga
	var (
		umur          int
		kartuKeluarga bool
	)

	// Inputan umur dan kartu keluarga
	fmt.Print("Masukan umur: ")
	fmt.Scan(&umur)

	fmt.Print("Masukan 1 jika punya kartu keluarga, 0 jika tidak punya kartu keluarga: ")
	fmt.Scan(&kartuKeluarga)

	// Percabangan untuk menentukan bisa membuat KTP
	if umur >= 17 && kartuKeluarga == true {
		fmt.Println("Bisa membuat KTP")
	} else {
		fmt.Println("Belum Bisa Membuat KTP")
	}

}