package main

import "fmt"

func main() {

	var bulan1, bulan2, profit float32

	fmt.Scan(&bulan1, &bulan2)

	profit = bulan2 - bulan1

	if profit > 0 {
		fmt.Println("Peningkatan Sebesar", profit)
	} else if profit == 0 {
		fmt.Println("Tetap")
	} else {
		fmt.Println("Penurunan Sebesar", -profit)
	}
}

//PSEUDO CODE
/*
Input bulan 1, bulan 2
Jika bulan 2 > bulan 1
Output "Peningkatan Sebesar bulan 2 - bulan 1"
Jika bulan 2 == bulan 1
Output "Tetap"
Jika bulan 2 < bulan 1
Output "Penurunan Sebesar bulan 1 - bulan 2"
*/