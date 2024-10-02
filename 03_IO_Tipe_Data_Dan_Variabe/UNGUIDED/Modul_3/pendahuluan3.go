
package main

import "fmt"

func main() {

	var (
		suhuFahrenheit float32
		suhuKelvin     float32
	)

	// Inputan untuk memasukan suhu dalam fahrenheit.
	fmt.Print("Masukan suhu dalam Fahrenheit: ")
	fmt.Scanln(&suhuFahrenheit)

	//Operasi untuk mengkonfersi dari fahrenheit ke suhu kelvin
	suhuKelvin = (suhuFahrenheit-32)*5/9 + 273
	fmt.Println("Suhu Kelvin nya adalah: ", suhuKelvin, "Kelvin")

}
