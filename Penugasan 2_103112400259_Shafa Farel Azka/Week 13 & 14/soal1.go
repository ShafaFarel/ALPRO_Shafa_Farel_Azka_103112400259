package main

import "fmt"

func main() {

	var angka int
	var kebenaran bool

	kebenaran = true

	fmt.Scan(&angka)

	if angka <= 1 {
		kebenaran = false
	} else {
		for i := 2; i < angka; i++ {
			if angka%i == 0 {
				kebenaran = false
				break
			}
		}
	}

	fmt.Println(kebenaran)

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