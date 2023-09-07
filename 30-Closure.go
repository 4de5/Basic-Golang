// Sama seperti scope di Java
// Program anak(child) bisa mengakses program induk(parents)

// HATI_HATI --> program induk bisa terubah apabila ceroboh dalam program anak

package main

import "fmt"

func main() {
	name := "Asep"
	angka := 1

	increment := func() {
		name = "Dwi"
		//AWAS!!!!!!!!!!!!!!!!!!!!!program akan merubah var name di Induk
		name := "Putra"
		// program deklarasi var baru, tidak merubah induk
		fmt.Println("Increment", name)
		angka++
		//AWAS!!!!!!!!!!!!!!!!!!!!!program akan merubah nilai data induk
	}

	increment() // Dalam kondisi ini program anak akan dijalankan
	increment() // AWASI PROGRAM INDUK!!!!!!!!!!!!!!!!!!!!!!!!!!!

	fmt.Println(name)
	fmt.Println(angka)
}
