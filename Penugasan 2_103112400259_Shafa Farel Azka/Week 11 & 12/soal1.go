package main

import "fmt"

func main() {

	var (
		jumlah             int
		username, password string
		selesai            bool
	)

	jumlah = -1

	for selesai = false; !selesai; {
		fmt.Scan(&username, &password)

		jumlah++
		selesai = username == "admin" && password == "admin"
	}

	fmt.Println(jumlah)
	fmt.Print("Login berhasil")

}

//PSEUDOCODE

/*	
	program Login
	variabel
		username, password: string
		jumlah: integer
		selesai: boolean
	algoritma
		jumlah = -1
		selesai = false
		while selesai == false do
			input(username, password)
			jumlah = jumlah + 1
			selesai = username == "admin" && password == "admin"
		end while
		output(jumlah)
		output("Login berhasil")
	end program
*/