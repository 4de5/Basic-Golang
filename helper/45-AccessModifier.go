/*
Akses yang diperbolehkan untuk di impor/digunakan di file lain

Apabila var dan/atau func berawalan kecil <sayHello> --> tidak bisa diakses
Apabila var dan/atau func berawalan besar <SayHello> --> bisa diakses
*/
package helper

import "fmt"

func SayHi(name string) { // bisa diakses
	fmt.Println("Hai", name)
}

func sayHi(name string) { // tak bisa
	fmt.Println("Hai", name)
}

// Berlaku untuk var
