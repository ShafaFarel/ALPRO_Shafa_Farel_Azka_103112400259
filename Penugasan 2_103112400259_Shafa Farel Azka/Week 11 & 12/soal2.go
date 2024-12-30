package main

import "fmt"

func main() {

	var (
		uang, jumlah int
		selesai      bool
	)

	for selesai = false; !selesai; {

		fmt.Scan(&uang)

		jumlah += uang

		if uang <= 0 {
			selesai = true
		}
	}

	fmt.Print(jumlah)

}

//Pseudocode
/*
program JumlahUang
variabel
	uang, jumlah: integer
	selesai: boolean
algoritma
	selesai = false
	while selesai == false do
		input(uang)
		jumlah = jumlah + uang
		if uang <= 0 then
			selesai = true
		end if
	end while
	output(jumlah)
end program
*/