package main

import "fmt"

func main() {

	var jumlahBarang int
	var totalPoin int

	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)

	for i := 1; i <= jumlahBarang; i++ {
		if i <= 5 {
			totalPoin += 10
		} else {
			totalPoin += 15
		}
	}

	fmt.Println("Total poin yang didapat:",totalPoin)
}