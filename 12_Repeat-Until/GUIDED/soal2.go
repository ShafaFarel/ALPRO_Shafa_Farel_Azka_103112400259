package main

import "fmt"

func main() {

	var bilangan int

	for selesai := false; !selesai;{
		fmt.Print("Masukkan bilangan: ")
		fmt.Scan(&bilangan)

		selesai = (bilangan > 0)
	}

	fmt.Println(bilangan,"adalah bilangan bulat positif")
}