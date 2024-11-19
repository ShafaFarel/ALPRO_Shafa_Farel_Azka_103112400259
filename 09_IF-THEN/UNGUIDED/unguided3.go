package main

import "fmt"

func main() {

	// Membuat variabel
	var (
		b     int
		prima int
		hasil bool
	)

	// Inputan untuk memasukan bilangan
	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&b)

	// Pecabangan yang pertama adalah untuk memeriksa apakah inputan lebih dari 0, jika tidak maka program tidak berjalan.
	if b > 0 {

		// Kemudian terdapat percabangan yang digunakan untuk menampilkan deret angka dimulai dari 1 sampai dengan bilangan yang dimasukan
		for i := 1; i <= b; i++ {
			// Perulangan kedua digunakan untuk menampilkan, angka tertentu, yaitu angka yang bisa habis membagi bilangan yang dimasukan
			if b%i == 0 {
				fmt.Print(i, " ")
				// Prima++ digunakan untuk menjumlah ada berapa angka yang bisa habis membagi bilangan yang dimasukan
				prima++
			}
		}
		// Kemudian terdapat boolean, yaitu menentukan prima atau tidak. karena prima hanya memiliki dua hasil, yaitu ketika hanya bisa dibagi dengan 1 dan angka itu sendiri, maka prima di definisikan dengan 2
		hasil = prima == 2
		fmt.Print(hasil)
		// Output yang ditampilkan ketika inputan kurang dari 0
	} else {
		fmt.Print("Masukan bilangan bulat lebih dari 0!")
	}

}