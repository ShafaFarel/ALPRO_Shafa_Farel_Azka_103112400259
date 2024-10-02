package main

import (
	"fmt"
)

func main() {

	//Membuat variable yang dimana akan menyimpan dari bilangan yang dimasukan pertama dan kedua, hasil operasi dan pilih untuk memilih opsi kalkulator.
	var (
		bilSatu, bilDua, hasil float32
		pilih   int8
	)

	// Menu yang dapat dipilih
	fmt.Println("Menu Pilihan")
	fmt.Println("-------------")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")

	//Inputan untuk memasukan pilihan
	fmt.Print("Ketikan Pilihanmu: ")
	fmt.Scan(&pilih)

	//Percabangan yang dimana ketika memilih satu, maka akan masuk ke menu penjumlahan, 2 untuk pengurangan, 3 untuk perkalian dan 4 untuk pembagian.
	if pilih == 1 {

		fmt.Println("Penjumlahan")
		fmt.Println("-------------")

		fmt.Print("Masukan angka pertama: ")
		fmt.Scan(&bilSatu)

		fmt.Print("Masukan angka pertama: ")
		fmt.Scan(&bilDua)

		hasil = bilSatu + bilDua //Operasi yang digunakan untuk menjumlahkan dua bilangan.

		fmt.Println("Hasil dari penjumlahan adalah", hasil)
	}
	if pilih == 2 {

		fmt.Println("Pengurangan")
		fmt.Println("-------------")

		fmt.Print("Masukan angka pertama: ")
		fmt.Scan(&bilSatu)

		fmt.Print("Masukan angka pertama: ")
		fmt.Scan(&bilDua)

		hasil = bilSatu - bilDua //Operasi untuk melakukan pengurangan.

		fmt.Println("Hasil dari pengurangan adalah", hasil)

	}
	if pilih == 3 {

		fmt.Println("Perkalian")
		fmt.Println("-------------")

		fmt.Print("Masukan angka pertama: ")
		fmt.Scan(&bilSatu)

		fmt.Print("Masukan angka pertama: ")
		fmt.Scan(&bilDua)

		hasil = bilSatu * bilDua //Operasi untuk mengeksekusi perkalian

		fmt.Println("Hasil dari perkalian adalah", hasil)

	}
	if pilih == 4 {

		fmt.Println("Pembagian")
		fmt.Println("-------------")

		fmt.Print("Masukan angka pertama: ")
		fmt.Scan(&bilSatu)

		fmt.Print("Masukan angka pertama: ")
		fmt.Scan(&bilDua)

		hasil = bilSatu / bilDua //Operasi untuk mengeksekusi pembagian

		fmt.Println("Hasil dari pembagian adalah", hasil)

	} else {
		fmt.Println("Menu Tidak Tersedia") //Pesan yang muncul ketika memasukan angka yang tidak ada pada pilihan.
	}
}