package main

import "fmt"

func main() {

	var x, n, digit int
	var kebenaran bool

	fmt.Scan(&x, &n)

	for n > 0 {

		digit = n % 10

		if digit == x {
			kebenaran = true
			break
		}

		n /= 10

	}

	fmt.Print(kebenaran)

}

//PSEUDO CODE
/*
program CekDigit
variabel
    x: integer
    n: integer
    digit: integer
    kebenaran: boolean

algoritma
    input(x, n)
    kebenaran = false

    while n > 0 do
        digit = n mod 10
        if digit == x then
            kebenaran = true
            break
        end if
        n = n div 10
    end while

    output(kebenaran)
end programs
*/