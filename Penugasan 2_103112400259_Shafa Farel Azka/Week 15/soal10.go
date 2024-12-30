package main

import "fmt"

func main() {

	var j, y, bilangan, total int

	fmt.Scan(&y)

	total = 0

	for j = 1; j <= 9; j++ {
		fmt.Scan(&bilangan)
		total += bilangan
	}

	if total >= y*5 {
		fmt.Print("Median Bernilai ", y)
	} else {
		fmt.Print("Median Bernilai 0")
	}

}

//PSEUDO CODE
/*
program MatriksDiagonal
variabel
    angka, baris, kolom: integer

algoritma
    input(angka)

    for baris = 1 to angka do
        output(newline)   // Pindah ke baris berikutnya

        for kolom = 1 to angka do
            if baris == kolom atau baris + kolom == angka + 1 then
                output(baris, " ") 
            else
                output("  ") 
            end if
        end for
    end for
end program
*/