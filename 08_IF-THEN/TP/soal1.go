package main

import "fmt"

func main() {

	var nilai int

	fmt.Print("Masukkan Nilai Ujian: ")
	fmt.Scan(&nilai)

	if nilai >= 70 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak lulus")
	}
}