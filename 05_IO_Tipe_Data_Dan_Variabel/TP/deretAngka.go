package main

import "fmt"

func main() {
	
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	jumlah := 0
	for i := 1; i <= n; i++ {
		jumlah += i
		fmt.Print(i, " ")
	}

	fmt.Println("Jumlah deret angka: ", jumlah)
}