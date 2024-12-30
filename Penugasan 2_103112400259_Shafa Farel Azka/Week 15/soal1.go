package main

import "fmt"

func main() {

	var x1, x2 int
	var status bool

	fmt.Scan(&x1, &x2)
	status = (x1%2 == 0 && x2%2 != 0) || (x1%2 != 0 && x2%2 == 0)

	if status == true {
		fmt.Print("Berhasil")
	}

}
//PSEUDOCODE
/*
program CekBilanganPrima
variabel
    angka: integer
    kebenaran: boolean

algoritma
    input(angka)

    if angka <= 1 then
        kebenaran = false
    else
        for i = 2 to angka-1 do
            if angka mod i == 0 then
                kebenaran = false
                break
            end if
        end for
    end if

    output(kebenaran)
end program
*/