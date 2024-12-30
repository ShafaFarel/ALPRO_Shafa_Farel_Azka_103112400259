package main

import "fmt"

func main() {

	var jam12, jam24 int

	fmt.Scan(&jam24)

	if jam24%12 == 0 {
		jam12 = 12
	} else {
		jam12 = jam24 % 12
	}

	if jam24 < 12 {
		fmt.Print(jam12, " AM")
	} else {
		fmt.Print(jam12, " PM")
	}

}

//PSEUDOCODE
/*
program CariLebarTerbesar
variabel
    jumlahDaun: integer
    lebar: integer
    lebarTerbesar: integer

algoritma
    input(jumlahDaun)
    lebarTerbesar = 0

    for i = 1 to jumlahDaun do
        input(lebar)
        if lebar > lebarTerbesar then
            lebarTerbesar = lebar
        end if
    end for

    output(lebarTerbesar)
end program
*/