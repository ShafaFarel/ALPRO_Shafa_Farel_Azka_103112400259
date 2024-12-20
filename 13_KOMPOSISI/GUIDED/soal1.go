package main

import "fmt"

func main() {

	var kata string
	var perulangan int

		fmt.Print("Masukkan kata: ")
        fmt.Scan(&kata)
		fmt.Print("Masukkan jumlah perulangan: ")
		fmt.Scan(&perulangan)
		counter := 0

		for selesai := false; !selesai;{
			fmt.Println(kata)
			counter++
			selesai = (counter >= perulangan)
		}
}