package main

import (
	"fmt"
)

func main() {

	var a, b, c, d, e, hasil int

	fmt.Print("Masukan bilangan Bulat: ")
	fmt.Scanln(&a, &b, &c, &d, &e)

	hasil = a + b + c + d + e //Digunakan untuk menjumlahkan semua inputan
	fmt.Println("Hasil penjumlahan: ", a, b, c, d, e, "adalah", hasil)

}