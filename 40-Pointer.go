package main

import "fmt"

type Address struct {
	City, Province, Country string
}

type Siswa struct {
	Nama        string
	Umur, NoAbs int
}

func main() {
	// Normal - PASS BY VALUE
	var alamat = Address{"Purbalingga", "Jawa Tengah", "Indonesia"}

	var alamat1 = alamat
	var alamat2 Address = alamat1

	alamat2.City = "Bekasi"     // Hanya alamat2 yg berubah karena Pass by Value
	alamat1.Country = "Wakanda" // alamat2 tidak berubah karena di copy sebelum edit

	// fmt.Println(alamat)
	// fmt.Println(alamat1)
	// fmt.Println(alamat2)

	//############################################

	// Pointer - PASS BY REFERENCE
	siswa := Siswa{Nama: "Asep", Umur: 17, NoAbs: 4}
	siswa1 := &siswa // Memasang Pointer di hulu
	/*
		CARA LAIN!!!!!!!
		var siswa Siswa = Siswa{Nama: "Asep", Umur: 17, NoAbs: 4}
		var siswa1 *Siswa = &siswa
	*/

	//siswa2 := &siswa1 // Memasang Pointer di hulu
	//siswa2.Umur = 12 // Hanya siswa yg berubah karena Pass by Value
	/*
		ternyata tidak bisa meng-pointer sebuah pointer
	*/

	siswa1.Umur = 11 // siswa1 diubah == siswa diubah == Siswa{...} diubah
	fmt.Println(siswa)
	fmt.Println(&siswa)  // &{Asep 11 4}
	fmt.Println(*siswa1) //karena var siswa 1 menghulu lewat siswa tak langsung ke Siswa, maka pake *siswa1
	fmt.Println(&siswa1) //  0xc0000ca018 => alamat data disimpan di memory
	//fmt.Println(siswa2) bisa tapi aneh run.e
}

/*
PASS BY VALUE --> data akan dicopy jadi tidak mempengaruhi data sumber
PASS BY REFERENCE --> data akan dihubungkan dengan sumber, jadi tidak mengcopy

komputer akan mengCopy dan menggunakan RAM untuk menyimpan sementara,  jadi itulah
kenapa PbR lebih bagus untuk kebutuhan bisnis(hemat)

saat menggunakan func/method data parameter juga PbV
*/
