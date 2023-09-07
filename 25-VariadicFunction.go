//atau variabel argumen
//Khusus untuk data parameter akhir
//VarArgs --> data bisa menrima input lebih dari 1 ataupun 0
//		--> seperti array/slice
//		--> 1 parameter bisa 0 atau 1 atau lebih data

package main

import "fmt"

func sumAll(angka ...int) int { //kuncinya di ...int
	// <...int> parameter ini harus di akhir, jika lebih dari 1

	total := 0
	for _, value := range angka {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10) //data parameter bebas mau isi berapa
	fmt.Println("Totalnya =", total)
	fmt.Println("")

	//STUDI KASUS KITA PUNYA DATA SLICE
	iniSliceNilai := []int{100, 100, 100, 90, 80, 75}
	totalNilai := sumAll(iniSliceNilai...)
	nilaiAkhir := totalNilai / len(iniSliceNilai)
	fmt.Println("Nilai Akhirnya =", nilaiAkhir)
}
