package main

import "fmt"

func main() {

	var angka, jumlah, digit int

	fmt.Scan(&angka)

	for angka > 0 {
		digit = angka % 10

		fmt.Print(digit, " ")

		jumlah += digit
		angka /= 10
	}

	fmt.Println("")
	fmt.Println(jumlah)

}

//PSEUDOCODE
/*
Input angka
Jika angka > 0
Output angka
Jumlah = jumlah + angka
angka = angka / 10
end if
Output jumlah
*/