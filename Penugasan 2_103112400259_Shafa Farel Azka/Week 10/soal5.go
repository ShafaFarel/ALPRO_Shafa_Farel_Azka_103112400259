package main

import "fmt"

func main() {

	var totalBelanja, hargaAkhir int
	var kartu, diskon, cashback bool

	fmt.Scan(&totalBelanja, &kartu)

	if totalBelanja >= 200000 && kartu == true {
		diskon = true
		cashback = true
	} else if totalBelanja >= 100000 {
		diskon = true
		cashback = false
	}

	if diskon == true && cashback == true {
		hargaAkhir = (totalBelanja - (totalBelanja * 10 / 100)) - 75000
	} else if diskon == true && cashback == false {
		hargaAkhir = totalBelanja - (totalBelanja * 10 / 100)
	} else {
		hargaAkhir = totalBelanja
	}

	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon", diskon)
	fmt.Println("Cashback?", cashback)
	fmt.Println("Rp", hargaAkhir)

}

//pseudo code
/*
Input total belanja, kartu
Jika total belanja >= 200000 && kartu == true
Output "Kartu?", true
Output "Diskon", true
Output "Cashback?", true
Output "Rp", total belanja - (total belanja * 10 / 100) - 75000
Jika total belanja >= 100000
Output "Kartu?", false
Output "Diskon", true
Output "Cashback?", false
Output "Rp", total belanja - (total belanja * 10 / 100)
*/