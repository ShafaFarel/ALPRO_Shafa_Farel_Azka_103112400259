package main

import "fmt"

func main() {

	var angka, baris, kolom int

	fmt.Scan(&angka)

	for baris = 1; baris <= angka; baris++ {

		for kolom = 1; kolom <= angka; kolom++ {
			if baris == kolom || baris+kolom == angka+1 {
				fmt.Print(baris, " ")
			} else {
				fmt.Print("  ")
			}
		}

		fmt.Println()
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