package main

import "fmt"

func main() {

	var jenisKendaraan string
	var durasiParkir int
	var tarifPerJam int
	var totalBiaya int

	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scanln(&jenisKendaraan)

	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scanln(&durasiParkir)

	switch {
	case durasiParkir < 1:
		durasiParkir = 1
	}

	switch jenisKendaraan {
		case "motor":
            tarifPerJam = 2000
			totalBiaya = tarifPerJam * durasiParkir
			fmt.Printf("%s %d jam Rp %d\n", jenisKendaraan, durasiParkir, totalBiaya)
        case "mobil":
            tarifPerJam = 5000
			totalBiaya = tarifPerJam * durasiParkir
			fmt.Printf("%s %d jam Rp %d\n", jenisKendaraan, durasiParkir, totalBiaya)
        case "truk":
            tarifPerJam = 8000
			totalBiaya = tarifPerJam * durasiParkir
			fmt.Printf("%s %d jam Rp %d\n", jenisKendaraan, durasiParkir, totalBiaya)
        default:
            fmt.Println("Jenis kendaraan yang dimasukkan salah!")
	}
}