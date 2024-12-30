package main

import "fmt"

func main() {

	var angka int

	fmt.Scan(&angka)

	for i := 1; i <= angka; i++ {

		for j := 1; j <= angka; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()

	}

}

//PSEUDOCODE
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