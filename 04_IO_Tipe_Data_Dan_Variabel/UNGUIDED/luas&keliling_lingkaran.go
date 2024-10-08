package main

import "fmt"

func main() {
	//panjang sisi alun-alun
	var jejari float64
	pi := 3.14

	//rumus keliling dan luas dari lingkaran
	keliling:= 2*pi*jejari
	luas:= pi*(jejari*jejari)

	fmt.Print("Masukkan jejari : ")
	fmt.Scan(&jejari)

	fmt.Println("Keliling : ", keliling,"cm")
	fmt.Println("Luas : ", luas,"cm²")
}