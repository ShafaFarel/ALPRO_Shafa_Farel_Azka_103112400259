package main

import "fmt"

func main() {

	var angka1, angka2, angka3, angka4, angka5 float32

	fmt.Scan(&angka1, &angka2, &angka3, &angka4, &angka5)

	if angka1 > angka2 && angka2 > angka3 && angka3 > angka4 && angka4 > angka5 || angka1 < angka2 && angka2 < angka3 && angka3 < angka4 && angka4 < angka5 {
		fmt.Println("Stabil naik/turun")
	} else {
		fmt.Println("Tidak stabil")
	}

}
//Pseudo Code
/*
Input angka 1, angka 2, angka 3, angka 4, angka 5
Jika angka 1 > angka 2 && angka 2 > angka 3 && angka 3 > angka 4 && angka 4 > angka 5 || angka 1 < angka 2 && angka 2 < angka 3 && angka 3 < angka 4 && angka 4 < angka 5
Output "Stabil naik/turun"
Output "Tidak stabil"
*/