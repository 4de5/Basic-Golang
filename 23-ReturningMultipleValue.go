//Mengembalikan fungsi di dalam fungsi == multiple
//bisa mengembalikan lebih dari 1 fungsi == multiple
//jika mengembalikan 2 data string-->maka ditulis 2x

package main

import "fmt"

func tampilkanNama() (string, string /*ini tipe data untung return*/) {
	return "Asep", "Putra" //pake <,> yang artinya pemisah
	//SEBELUMNYA, aku pake <+> itu diitung 1 return
}

func main() { //dalam 1 package sebenarnya tidak boleh ada funnction yang sama,
	//bisa dilihat <main> bertanda merah karena di duplicate tiap file

	fmt.Println(tampilkanNama())

	namaPanggilan, _ /*disini saya mengABAIkan return ke 2, karena tidak dipakai*/ := tampilkanNama()
	fmt.Println(namaPanggilan)
	// fmt.Println(namaLain) //data namLain saya ignore<_>, karena tidak dipake

}
