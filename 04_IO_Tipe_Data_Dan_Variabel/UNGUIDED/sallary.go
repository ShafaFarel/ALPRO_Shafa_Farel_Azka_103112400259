package main

import "fmt"

func hitungGaji(jamKerja float64, upahPerJam float64) float64 {
	var jamNormal, jamLembur float64

	if jamKerja > 40 {
		jamNormal = 40
		jamLembur = jamKerja - 40
	} else {
		jamNormal = jamKerja
		jamLembur = 0
	}

	totalGaji := (jamNormal * upahPerJam) + (jamLembur * 1.5 * upahPerJam)

	// Menghitung gaji bulanan (4 minggu)
	return totalGaji * 4
}

func main() {
	var jamKerjaPerMinggu float64
	var upahPerJam float64

	// Meminta input dari pengguna
	fmt.Print("Masukkan jumlah jam kerja dalam seminggu: ")
	fmt.Scan(&jamKerjaPerMinggu)

	fmt.Print("Masukkan upah per jam: ")
	fmt.Scan(&upahPerJam)

	// Menghitung gaji
	totalGaji := hitungGaji(jamKerjaPerMinggu, upahPerJam)

	// Menampilkan total gaji bulanan
	fmt.Printf("Total gaji bulanan: %.2f\n", totalGaji)
}