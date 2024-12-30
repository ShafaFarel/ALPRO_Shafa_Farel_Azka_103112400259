package main

import "fmt"

func main() {

	var masukan string
	var jumlahA, jumlahB, jumlahC int

	for {
		fmt.Scan(&masukan)

		if masukan == "A" {
			jumlahA++
		} else if masukan == "B" {
			jumlahB++
		} else if masukan == "C" {
			jumlahC++
		} else {
			break
		}
	}

	fmt.Println("Tipe A:", jumlahA)
	fmt.Println("Tipe B:", jumlahB)
	fmt.Println("Tipe C:", jumlahC)
}

//PSEUDO CODE
/*
program HitungTipe
variabel
    masukan: string
    jumlahA, jumlahB, jumlahC: integer

algoritma
    jumlahA = 0
    jumlahB = 0
    jumlahC = 0

    while true do
        input(masukan)

        if masukan == "A" then
            jumlahA = jumlahA + 1
        else if masukan == "B" then
            jumlahB = jumlahB + 1
        else if masukan == "C" then
            jumlahC = jumlahC + 1
        else
            break
        end if
    end while

    output("Tipe A:", jumlahA)
    output("Tipe B:", jumlahB)
    output("Tipe C:", jumlahC)
end program
*/