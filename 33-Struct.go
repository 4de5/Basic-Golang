/**
STRUC --> adalah sebuah prototipe data/objek
--> Hampir sama dengan Array, Slice, Map. Namun bisa berbagai tipe data
--> Tidak bisa langsung mendeklarasi isinya

*/

package main

import "fmt"

// Deklarasi Struct sebagai Prototipe
type Customer struct { // Nama berawal besar = UperCase = Kebiasaan Programer
	Field1, Field2 string // Nama berawal besar = UperCase = Kebiasaan Programer
	Field3         int
	Field4         bool
}

func main() {
	//Pengisian data/objek pada struct
	var asep Customer
	asep.Field1 = "Asep Dwi Saputra"
	asep.Field2 = "Indonesia"
	asep.Field3 = 20
	asep.Field4 = true

	fmt.Println(asep)

	//STRUCT LITERAL 1
	oki := Customer{
		Field1: "Oki Budianto",
		Field2: "Indonesia",
		Field3: 24,
		Field4: true,
	}
	fmt.Println(oki.Field1)

	//STRUCT LITERAL 2
	bose := Customer{"Bose", "Indonesia", 23, false}
	//Kelemahan Struct 2 ini !!!!!!!!!!!!!!!!!!!
	//1. Pengisian data harus sesuai dengan Fieldconst
	//2. Apabila Field berubah tempat, tipe , ataubahkan dihapus akan Eror
	//3. Perlu TELITI jika ada eror

	fmt.Println(bose.Field3)

	// STRUCT TAG DI MODUL 56
}
