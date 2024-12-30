package main

import "fmt"

func main() {

	var angka, terbesar, digit int

	fmt.Scan(&angka)

	terbesar = 0

	for angka > 0 {

		digit = angka % 10

		if digit > terbesar {
			terbesar = digit
		}

		angka /= 10

	}

	fmt.Println(terbesar)

}
//PSEUDOCODE
/*
program CariDigitTerbesar
variabel
    angka: integer
    terbesar: integer
    digit: integer

algoritma
    input(angka)
    terbesar = 0

    while angka > 0 do
        digit = angka mod 10
        if digit > terbesar then
            terbesar = digit
        end if
        angka = angka div 10
    end while

    output(terbesar)
end program
*/