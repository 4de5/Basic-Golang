//percabangan yang hampir sama denga ifElse, namun lebih ringkas

package main

import "fmt"

func main() {

	var nama = "Asep Dwi Saputra"

	//Cara 1
	switch nama {
	case "Asep Dwi Saputra":
		fmt.Println("Ini Asep") //ga pake break kaya di java, JS
	default:
		fmt.Println("Ini Bukan Asep")
	}

	//Cara Short
	switch karakterNama := len(nama); karakterNama > 16 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sesuai")
	}
}
