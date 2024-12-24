package main

import "fmt"

func main(){
	var nama_tanaman string
	fmt.Scan(&nama_tanaman)
	switch nama_tanaman{
	case "nepenthes", "drosera":
		fmt.Println("Termasuk tanaman karnivora")
		fmt.Println("Asli Indonesia")
	case "venus", "sarracenia":
		fmt.Println("Termasuk tanaman karnivora")
		fmt.Println("tidak asli Indonesia")
	default:
		fmt.Println("tidak termasuk tanaman karnivora")
	}
}