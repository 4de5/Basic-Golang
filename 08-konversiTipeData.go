package main

import (
	"fmt"
)

func main() {

	//konversi NUMBER
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32) //merubah nilai32(int32) ke nilai64(int64)
	var nilai8 int8 = int8(nilai32)    //mengecilkan HATI-HATI

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) //terjadi overFlow>>setelah data terisi penuh direset

	//konversi STRING
	var nama string = "Asep"
	var A = nama[0]                      //output bernilai byte/uint8
	var konversiAjadi string = string(A) //biar uint8 jadi karakter string

	fmt.Println(nama)
	fmt.Println(A)
	fmt.Println(konversiAjadi)
}
