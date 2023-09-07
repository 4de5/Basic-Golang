//Operasi yang hanya digunakan untuk tipe data BOOLEAN
//Contohnya --> AND --> && --> Bila T-T = T, lainnya F
//			--> OR  --> || --> Bila F-F = F, lainnya T
//			--> NOT -->  ! --> Bila T jadi F, atau sebaliknya

package main

import "fmt"

func main() {

	var nilaiUAS = 90
	var nilaiAsen = 80

	var syaratUAS = nilaiUAS >= 80
	var syaratAbsen = nilaiAsen >= 80

	var NaikKelas = syaratAbsen && syaratUAS

	fmt.Println(NaikKelas)

	if NaikKelas == true {
		fmt.Println("Selamat Kamu Naik Kelas")
	}
}
