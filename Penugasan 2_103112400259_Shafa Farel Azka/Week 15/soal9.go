package main

import "fmt"

func main() {

	var p, r, hari int

	fmt.Scan(&p, &r)
	hari = p / r

	if p%r > 0 {
		hari++
	}

	fmt.Print(hari, " Hari")

}

//PSEUDO CODE
/*
program MatriksBingkai
variabel
    angka, i, j: integer

algoritma
    input(angka)

    for i = 1 to angka do
        output(newline)  

        for j = 1 to angka do
            if i == 1 atau i == angka atau j == 1 atau j == angka then
                output(i, " ")
            else
                output("  ") 
            end if
        end for
    end for
end program
*/