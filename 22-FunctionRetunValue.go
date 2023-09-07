//Function bisa mengembalikan data
//Sebelum itu kita perlu menulis tipe data apa yang akan dikembalikan
//Jika tipe data sudah dideklarasi kita wajib untuk mengembakikan(return)
//return <nilai kemabli> --> untuk mengembalikan data

package main

import "fmt"

func penjumlahan(angka1 int, angka2 int) int /*ini tipe data return*/ {
	return angka1 + angka2
	//apabila ada coding dibawah <return> maka tidak dianggap, dan EROR
}

func main() {
	fmt.Println(penjumlahan(1, 2))

	//Harga Apel adalah harga Pisang + harga Mangga
	//harga pisang = 1000, mangga = 2000, apel?
	hargaPisang := 1000
	hargaMangga := 2000
	hargaApel := penjumlahan(hargaPisang, hargaMangga)
	fmt.Println("Harga Apel =", hargaApel)
}
