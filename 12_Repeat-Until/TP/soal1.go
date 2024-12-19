package main

import "fmt"

func main() {

	var angka int

	tebakanBenar := 7
	for {
		fmt.Print("Tebak Angka (1-10): ")
		fmt.Scan(&angka)

		if angka == tebakanBenar{
			fmt.Println("Selamat, tebakan Anda benar!")
			break
		} else{
			fmt.Println("Tebakan Anda salah, coba lagi.")
		}
	}
}