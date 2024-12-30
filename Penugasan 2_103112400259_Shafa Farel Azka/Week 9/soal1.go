package main

import "fmt"

func main() {

	var orang int

	fmt.Print("Masukkan jumlah orang : ")
	fmt.Scan(&orang)

	mobil := orang / 7
	sisaKursi := orang % 7

	if sisaKursi == 0 {
		mobil = mobil
	} else {
		mobil = mobil + 1
	}

	fmt.Print(mobil)

}

//PSEUDOCODE

/* 
program JumlahMobil
variabel
    orang, mobil, sisaKursi: integer
algoritma
    output("Masukkan jumlah orang :")
    input(orang)
    mobil = orang / 7
    sisaKursi = orang % 7
    if sisaKursi == 0 then
        mobil = mobil
    else
        mobil = mobil + 1
    end if
    output(mobil)
end program

*/