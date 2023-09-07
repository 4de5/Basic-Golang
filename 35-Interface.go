// Adalah sebuah SOP/Standar dalam membuat struct
// Struct akan tehubung otomatis, tidak seperti di Java, C++

package main

import "fmt"

type PemilikRumah interface {
	GetName() string //PAKE TIPE DATA!!!!!!!
}

func Welcome(pemilikRumah PemilikRumah) {
	fmt.Println("Selamat datang di Rumah, Mr", pemilikRumah.GetName())
} // Belum cukup disini kita perlu membuat Struct+GetName()

type Person struct {
	Name string
	Umur int
}

func (person Person) GetName() string {
	return person.Name // Kembalian sesuai Welcome()
}

func main() {
	//var asep PemilikRumah
	//Welcome(asep) // EROR, perlu ada Struct+GetName()

	asep := Person{Name: "Putra", Umur: 20}
	Welcome(asep)

	/**
	Setiap program yang menggunakan Func GetName() string,
	WAJIB mengikuti aturan interface PemilikRumah+func Welcome
	*/
}
