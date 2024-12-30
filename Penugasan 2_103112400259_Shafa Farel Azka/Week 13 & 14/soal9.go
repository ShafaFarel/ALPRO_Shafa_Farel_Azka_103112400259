package main

import "fmt"

func main() {

	var angka int

	fmt.Scan(&angka)

	for i := 1; i <= angka; i++ {
		fmt.Println()

		for j := 1; j <= angka; j++ {
			if i == 1 || i == angka || j == 1 || j == angka {
				fmt.Print(i, " ")
			} else {
				fmt.Print("  ")
			}
		}
	}

}

//pseudocode
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