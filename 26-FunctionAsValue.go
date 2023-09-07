//First Class Citizen --> bisa dirubah, bisa jadi variabel, independen, ga perlu Class
// Fungsi bisa di inisialkan dengan value atau variabel
// Tidak seperti di Java yang harus bergantung pada Class, di Go tidak ada func anak tiri

// Manfaatnya --> nanti misal butuh func yang butuh parameter func juga,
//					nanti parameter func.nya dibungkus value, value di bungkus var,
//					var digunakan untuk func barunya.

package main

import "fmt"

func ucapan(name string) string {
	return "Hai " + name
}

func main() {

	sapaOrang := ucapan("Kamu") //kita bungkus func ucap ke var sapaOrang
	fmt.Println(sapaOrang)

	var haiAsep = ucapan("Asep")
	fmt.Println(haiAsep)
}
