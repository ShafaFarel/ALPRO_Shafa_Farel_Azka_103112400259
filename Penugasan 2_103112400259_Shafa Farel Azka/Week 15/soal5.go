package main

import "fmt"

func main() {

	var kTidur, tKebun, tMain int
	var kondisi1, kondisi2, kondisi3 bool

	fmt.Scan(&kTidur, &tMain, &tKebun)

	kondisi1 = kTidur >= 60 && tMain >= 75 && tKebun >= 60
	kondisi2 = kTidur >= 80 && tKebun >= 80
	kondisi3 = kTidur >= 100

	if kondisi1 || kondisi2 || kondisi3 {
		fmt.Print("Ice cream")
	} else {
		fmt.Print("Tidak")
	}

}

//PSEUDO CODE
/*
Input kTidur, tMain, tKebun
Jika kTidur >= 60 && tMain >= 75 && tKebun >= 60
Output "Ice cream"
Jika kTidur >= 80 && tKebun >= 80
Output "Ice cream"
Jika kTidur >= 100
Output "Ice cream"
*/