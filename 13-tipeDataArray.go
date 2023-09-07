//array di Go tidak bisa menampung tipe data yang berbeda
//array tidak bisa ditambah daya tampungnya

package main

import "fmt"

func main() {

	//Cara 1
	var names [3]string
	names[0] = "Asep"
	names[1] = "Dwi"
	names[2] = "Saputra"

	fmt.Println(names)

	//Cara 2
	var values = [3]int{
		90,
		80,
		100,
	}
	fmt.Println(values)

	//Cara 3
	numbers := [...]int{ //disini array tidak langsung ditentukan panjangnya
		1, 2, 3, 4, 5, //dari sini Go tau panjang array 5length
	}
	fmt.Println(numbers)
}
