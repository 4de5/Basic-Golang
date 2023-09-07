//perulangan

package main

import "fmt"

func main() {

	//Cara 1
	angka := 1
	for angka <= 3 {
		fmt.Println(angka)
		angka++
	}
	fmt.Println("")

	//Cara 2
	for nomer := 1; nomer <= 3; nomer++ {
		fmt.Println(nomer)
	}
	fmt.Println("")

	//Mencetak array, slice, map dengan For RANGE

	//Cara 1
	inislice := []string{"Asep", "Dwi", "Saputra"}

	//OFF
	fmt.Println("RANGE OFF")
	for i := 0; i < len(inislice); i++ { //i=index
		fmt.Println(i, inislice[i])
	}

	//ON
	fmt.Println("RANGE ON")
	for index, value := range inislice {
		fmt.Println(index, value) //bila var index tak ingin digunakan bisa digabti <_>
	}

	//Menonaktifkan variabel yang tak ingin dipakai, fanti dengan <_>
	fmt.Println("RANGE ON !INDEX")
	for _, value := range inislice {
		fmt.Println(value) //bila var index tak ingin digunakan bisa digabti <_>
	}
}
