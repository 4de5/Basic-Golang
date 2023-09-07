//String adalah kumpulan karakter(bisa nol dan  spasi di hitung)
//ditulis --> string --> <"...">

// function string
//len("string") --> menghiung panjang/jumlah karakter di string
//"string"[index] --> mangambil karakter di index, hasil di tulis kode

package main

import "fmt"

func main() {
	fmt.Println("Asep")
	fmt.Println(len("Asep")) //panjang 4
	fmt.Println("Asep"[0])   //kode byte/uint8 huruf A

	kata := "cat;dog"
	a := kata[:4] // ambil 4 index pertama
	b := kata[:3]
	c := kata[4:] // tutup 4 index pertama
	d := kata[3:]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
