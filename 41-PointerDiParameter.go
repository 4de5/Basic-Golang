package main

import "fmt"

type Address struct {
	City, Country string
}

func AddressChanger(alamat Address) {
	alamat.Country = "Indonesia"
	fmt.Println(alamat) // Mencetak alamat yang sudah di edit
}

func AddressChanger2(alamat *Address) { // alamat ber-hulu Address
	alamat.Country = "Indonesia"
	fmt.Println(alamat)
}

func main() {
	address := Address{
		City:    "Purbalingga",
		Country: "aa",
	}

	AddressChanger(address)
	fmt.Println(address.City)
	fmt.Println(address.Country)
	fmt.Println()

	AddressChanger2(&address)
	// Apabila pemanggilan &address. func2 mencetak hulu
	// Apabila pemanggilan address. func2 mencetak hulu apabila address sudah mengHulu
	fmt.Println(address.City)
	fmt.Println(address.Country)
}
