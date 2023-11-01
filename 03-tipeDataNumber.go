//tipe data number terdiri dari integer dan floating point

//biasa --> int8(-128 s/d 127), int16(=/- 32k), int32(+/- 2B), int64(paling la intine)
//unsign --> tanpa minus --> nilai 2x data di atas(uint8, uint16, uint32, uint64)

//float --> float32(1,18 x 10^-38 sampai 3,4 x 10^38) --> normal
//		--> float64(2,23 x 10^-308 sampai 1,80 x 10^308)--> banyak
//complex --> untuk perhitungan aritmatika kompleks --> complex32, complex64

//ALIAS
//int --> bisa ditulis begini, nilai tergantung bit laptop(32 bit=int32)
//uint --> bisa ditulis begini, nilai tergantung bit laptop(32 bit=int32)
//rune --> bisa ditulis begini, nilai = int32
//byte --> bisa ditulis begini, nilai = uint8

// Contoh kode program
package main

import "fmt"

func main() {
	fmt.Println("Satu =", 1)
	fmt.Println("Dua =", 2)
}
