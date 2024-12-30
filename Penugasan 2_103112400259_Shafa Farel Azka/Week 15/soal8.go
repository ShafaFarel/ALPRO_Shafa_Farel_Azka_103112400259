package main

import "fmt"

func main() {

	var s, minggu int

	fmt.Scan(&s)

	minggu = s / 7

	if s%7 == 0 {
		fmt.Print("Minggu ke-", minggu)
	} else {
		fmt.Print("Minggu ke-", minggu+1)
	}

}

// PSEUDO CODE
/*
program MatriksAngka
variabel
    angka, i, j: integer

algoritma
    input(angka)

    for i = 1 to angka do
        for j = 1 to angka do
            output(j, " ")
        end for
        output(newline)
    end for
end program
*/