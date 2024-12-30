package main

import "fmt"

func main() {

	var angka int

	fmt.Scan(&angka)

	biner := ""

	for angka > 0 {
		biner = fmt.Sprintf("%d%s", angka%2, biner)
		angka /= 2
	}

	fmt.Println(biner)

}

//PSEUDOCODE
/*
program KonversiKeBiner
variabel
    angka: integer
    biner: string

algoritma
    input(angka)
    biner = ""

    while angka > 0 do
        biner = angka mod 2 + biner
        angka = angka / 2
    end while

    output(biner)
end program
*/