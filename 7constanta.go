//constanta --> const --> sama kaya di JS, variabel yang isinya ga bisa di ubah
//saat di buat isi HARUS langsung di deklarasi

package main

import "fmt"

func main() {

	//Cara 1
	const nama string = "Asep"
	// nama = "Putra" //mencoba merubah artinya EROR
	fmt.Println(nama)

	//Cara 2
	const name = "Asep"
	fmt.Println(name)

	//Multiple Const
	const (
		firstName  = "Asep"
		secondName = "Putra"
	)
	fmt.Println(firstName)
	fmt.Println(secondName)
}
