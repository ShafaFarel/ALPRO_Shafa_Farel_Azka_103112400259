package main

import "fmt"

func main(){
	var token string
	maxPercobaan := 3
	percobaan := 0

	for percobaan < maxPercobaan {
		fmt.Print("Masukan password: ")
		fmt.Scan(&token)

		if token == "golang" {
			fmt.Println("Selamat anda telah login")
		}

	fmt.Print("Masukan password: ")
	fmt.Scan(&token)

	for token == "golang" {
		fmt.Println("Password salah!, coba lagi")
		fmt.Print("Masukan password: ")
		fmt.Scan(&token)
	}
	fmt.Println("Password anda benar")
}

}