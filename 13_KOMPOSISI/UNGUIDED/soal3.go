package main

import "fmt"

func main() {

	var a,b,c,d string
	var berhasil bool
	
	berhasil = true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan", i, ": ")
		fmt.Scan(&a,&b,&c,&d)
		if !(a == "merah" && b == "kuning" && c == "hijau" && d == "ungu") {
			berhasil = false
		}
	}

	fmt.Print("BERHASIL:", berhasil)
}