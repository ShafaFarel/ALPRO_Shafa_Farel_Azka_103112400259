package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1

	var tebakan int
	var percobaan int
	const maxPercobaan = 5

	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Selamat datang di permainan Tebak Angka!")
	fmt.Println("Tebak angka antara 1 sampai 100")
	fmt.Println("Kamu punya kesempatan 5 kali untuk menebak angka yang benar")
	fmt.Println("----------------------------------------------------------------")

	for percobaan = 1; percobaan <= maxPercobaan; percobaan++ {
		fmt.Print("Masukkan tebakanmu: ")
		fmt.Scanln(&tebakan)
	

		if tebakan < target {
			fmt.Println("Tebakanmu terlalu kecil, coba lagi!")
			fmt.Println("----------------------------------------------------------------")
		} else if tebakan > target {
			fmt.Println("Tebakanmu terlalu besar, coba lagi!")
			fmt.Println("----------------------------------------------------------------")
		} else {
			fmt.Println("Selamat, tebakanmu benar dalam", percobaan, "percobaan!")
			break
	}
	}

	if percobaan > maxPercobaan {
		fmt.Println("Maaf, kesempatan habis, angka yang benar adalah", target)
	}
}