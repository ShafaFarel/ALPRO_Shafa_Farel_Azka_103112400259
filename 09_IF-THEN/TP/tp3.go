package main

import (
	"fmt"
)

func main() {

	// Membuat variabel umur dan kewarganegaran
	var (
		umur           int
		kewarganegaran string
	)

	// Input umur
	fmt.Print("Masukan umur: ")
	fmt.Scan(&umur)

	// Input kewarganegaran, yaitu disediakan pilihan a sebuah wni dan b adalah wna
	fmt.Println("Masukan a atau b")
	fmt.Println("a. WNI")
	fmt.Println("b. WNA")
	fmt.Print("Masukan pilihan kamu: ")
	fmt.Scan(&kewarganegaran)

	// Percabangan yang mengatur kondisi ketika umur dan kewarganegaran dimasukan
	if umur >= 17 && kewarganegaran == "a" {
		fmt.Println("Kamu bisa mengikuti pemilu")
	} else if umur >= 17 && kewarganegaran == "b" {
		fmt.Println("Kamu tidak bisa mengikuti pemilu karena kamu bukan WNI")
	} else if umur < 17 && kewarganegaran == "a" {
		fmt.Println("Kamu tidak bisa mengikuti pemilu karena umur kamu kurang dari 17 tahun")
	} else {
		fmt.Println("Kamu tidak bisa mengikuti pemilu karena umur kamu kurang dari 17 tahun dan kamu bukan WNI")
	}
}