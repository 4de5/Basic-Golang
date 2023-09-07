// sebuah Expresi --> kode  program yang menghasilkan nilai
// digunakan untuk membuat percabangan "Jika ... maka ..." dan ragam jenisnya
package main

import "fmt"

func main() {

	//Kita buat Variabel untuk dibandingkan
	//Jika perbandingan bernilai True, maka dijalankan
	nama := "Asep"
	fmt.Println(nama)

	if nama == "Asep" { //bila <nama == "Asep"> adalah True, maka dijalankan
		fmt.Println("Halo, sep!")
	}

	if nama == "Asep" { //bila <nama == "Asep"> adalah True, maka dijalankan
		fmt.Println("Apa kabar?")
	} else { //bila <nama == "Asep"> adalah False, maka dijalankan
		fmt.Println("Kamu Asep?")
	}

	if nama == "Asep" { //bila <nama == "Asep"> adalah True, maka dijalankan
		fmt.Println("Apa kabar?")
	} else if nama == "Oke" { //bila <nama == "Oke"> adalah True, maka dijalankan
		fmt.Println("Kamu Oke?")
	} else { //bila kedua perbandingan di atas tidak ada True, maka dijalankan
		fmt.Println("Kamu siapa?")
	}

	//Membuat short ifElse
	if totalNama := len(nama); totalNama > 16 { //Ini shortnya if ... := len(...);
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah sesuai")
	}
}
