//TYPE DECLARATIONS --> meng-Alias tipe data yang sudah ada dengan nama baru yg kita mau
//caranya --> type <nama alias> <tipe data aslinya>

package main

import "fmt"

func main() {

	//CONTOH
	type NIK string
	type statusMenikah bool

	var NIK_Saya NIK = "00899299"
	var status statusMenikah = false

	fmt.Println(NIK_Saya)
	fmt.Println(status)

	//di modul 27 kita akan meng-type sebuah function
}
