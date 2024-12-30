package main

import "fmt"

func main() {

	var gula, kopi, jumlahGula, jumlahKopi, terbuat int

	fmt.Scan(&gula, &kopi, &jumlahGula, &jumlahKopi)

	for kopi >= jumlahKopi && gula >= jumlahGula {
		kopi -= jumlahKopi
		gula -= jumlahGula

		terbuat++
	}

	fmt.Println(terbuat)

}

//PSEUDO CODE
/*
Input gula, kopi, jumlah gula, jumlah kopi
Jika kopi >= jumlah kopi && gula >= jumlah gula
Output terbuat
*/