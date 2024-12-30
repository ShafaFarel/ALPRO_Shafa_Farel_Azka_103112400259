package main

import "fmt"

func main() {

	var datang, jdatang, mdatang, durasi, jselesai, mselesai int

	fmt.Scan(&jdatang, &mdatang, &durasi)

	datang = jdatang*60 + mdatang + durasi
	jselesai = datang / 60
	mselesai = datang % 60

	if jselesai < 20 || (jselesai == 20 && mselesai == 0) {
		fmt.Print(jselesai, mselesai)
	} else {
		fmt.Print("Silahkan datang kembali besok")
	}

}

//PSEUDOCODE
/*
program CekDurasi
variabel
	datang, jdatang, mdatang, durasi, jselesai, mselesai: integer
algoritma
	input(jdatang, mdatang, durasi)
	datang = jdatang * 60 + mdatang + durasi
	jselesai = datang / 60
	mselesai = datang % 60
	if jselesai < 20 || (jselesai == 20 && mselesai == 0) then
		output(jselesai, mselesai)
	else
		output("Silahkan datang kembali besok")
	end if
end program
*/