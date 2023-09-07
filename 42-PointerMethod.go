package main

import "fmt"

type Orang struct {
	Nama string
}

func (id Orang) OffPointer() {
	fmt.Println("Ini Non-aktif: ")
	id.Nama = "Hai " + id.Nama
}

func (id *Orang) OnPointer() {
	fmt.Println("Ini Aktif: ")
	id.Nama = "Hai " + id.Nama
}

func main() {
	var asep Orang = Orang{"Asep"}

	asep.OffPointer() // Saat data asep dipanggil dengan func == di copy
	fmt.Println(asep.Nama)

	asep.OnPointer() // jika ingin PvR, tipe data struct harus *Orang
	fmt.Println(asep.Nama)
}
