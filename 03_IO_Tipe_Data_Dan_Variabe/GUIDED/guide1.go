package main

import (
	"fmt"
)

func main() {
	var sisi, hasil int

	fmt.Print("Masukan Panjang Sisi: ")
	fmt.Scan(&sisi)

	hasil = sisi * sisi * sisi //Rumus kubus
	fmt.Println("Volume kubus adalah = ", hasil)
}
