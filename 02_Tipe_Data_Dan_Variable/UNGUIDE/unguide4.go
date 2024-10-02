package main

import "fmt"

func main() {

	var (
		suhuFahrenheit int16
		suhuCelcius    int16
	)

	fmt.Print("Masukan suhu dalam Fahrenheit: ")
	fmt.Scanln(&suhuFahrenheit)

	suhuCelcius = (suhuFahrenheit - 32) * 5 / 9 //Rumus untuk konfersi suhu fahrenheit ke celcius 
	fmt.Println("Suhu Celcius nya adalah: ", suhuCelcius)

}