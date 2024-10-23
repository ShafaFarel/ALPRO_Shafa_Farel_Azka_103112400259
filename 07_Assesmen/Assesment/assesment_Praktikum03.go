package main

import (
	"fmt"
)

func main() {
	var dinar, dirham, fals, qirat int

	fmt.Print("Masukkan qirat: ")
	fmt.Scan(&qirat)

	dinar = qirat / 600 % 10
	dirham = qirat / 60 % 10
	fals = qirat / 6 % 10
	qirat= qirat % 6

	fmt.Print(dinar, dirham, fals, qirat)


}