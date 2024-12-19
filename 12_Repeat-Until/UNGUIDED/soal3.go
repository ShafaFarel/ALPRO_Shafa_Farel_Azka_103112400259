package main

import "fmt"

func main() {
	var target, donasi, totalDonasi, jumlahDonatur int

	// Meminta input target donasi
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	// Inisialisasi variabel untuk total donasi dan jumlah donatur
	totalDonasi = 0
	jumlahDonatur = 0

	// Loop untuk meminta donasi dari setiap donatur
	for totalDonasi < target {
		fmt.Print("Masukkan jumlah donasi: ")
		fmt.Scan(&donasi)

		// Menambahkan donasi ke total donasi
		totalDonasi += donasi
		jumlahDonatur++

		// Menampilkan informasi untuk setiap donatur
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, totalDonasi)
	}

	// Menampilkan hasil akhir setelah target tercapai
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, jumlahDonatur)
}