package main

import "fmt"

func main() {

	var x, y, z int
	var belanja, diskon float32

	fmt.Scan(&x, &y, &z)
	diskon = float32(z) * float32(x) / 100

	if diskon > float32(y) {
		diskon = float32(y)
	}

	belanja = float32(z) - diskon
	fmt.Printf("%.0f Rupiah\n", belanja)

}

//PSEUDO CODE
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