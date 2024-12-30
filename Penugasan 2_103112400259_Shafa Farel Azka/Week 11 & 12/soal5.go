package main

import "fmt"

func main() {

	var angka, pertama, kedua int
	var kebenaran bool

	kebenaran = true

	fmt.Scan(&angka)

	pertama = angka % 10
	angka /= 10

	for angka > 0 {
		kedua = angka % 10
		if pertama-kedua != 1 && kedua-pertama != 1 {
			kebenaran = false
			break
		}

		pertama = kedua
		angka /= 10

	}

	fmt.Println(kebenaran)

}

//PSEUDOCODE
/*
Input angka
Pertama = angka % 10
angka = angka / 10
Kedua = angka % 10
Jika pertama - kedua != 1 && kedua - pertama != 1
Output "false"
Output "true"
*/