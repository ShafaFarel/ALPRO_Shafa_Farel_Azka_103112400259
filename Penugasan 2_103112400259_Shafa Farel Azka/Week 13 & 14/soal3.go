package main

import "fmt"

func main() {

	var jumlahDaun, lebar, lebarTerbesar int

	fmt.Scan(&jumlahDaun)

	lebarTerbesar = 0

	for i := 1; i <= jumlahDaun; i++ {
		fmt.Scan(&lebar)

		if lebar > lebarTerbesar {
			lebarTerbesar = lebar
		}
	}

	fmt.Print(lebarTerbesar)
}

//PSEUDO CODE
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