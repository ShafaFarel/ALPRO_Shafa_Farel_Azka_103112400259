package main

import "fmt"

func main() {

	var (
		beratBadan, tinggiBadan, hasil float32
	)

	fmt.Print("Masukan berat badan: ")
	fmt.Scan(&beratBadan)

	fmt.Print("Masukan tinggi badan: ")
	fmt.Scan(&tinggiBadan)

	hasil = beratBadan / (tinggiBadan * tinggiBadan)

	fmt.Printf("BMI %.2f", hasil)
}