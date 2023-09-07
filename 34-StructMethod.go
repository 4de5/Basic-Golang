// Struct Method/Function
// Struc sama saja dengan tipe data lain(int, string, bool) ...
// ... dan bisa  dimsukkan ke dalam parameter fungsi
//sebenarnya struc disini menyamar sebagai fungsi

/**
Kenapa harus struct method --> ada baiknya bila struct terhubung dengan fungsi
maka dibuat struct function/struct method
(kebiasaan programer)
*/

package main

import "fmt"

type Customer struct {
	Nama, Alamat string
}

//NORMAL
func sayHi(nama string, customer Customer) {
	fmt.Println("Hello", nama, ",Saya", customer.Nama, "dari", customer.Alamat)
}

//STRUCT METHOD
func (customer Customer) sayHi(nama string) {
	fmt.Println("Hello", nama, ",Saya", customer.Nama, "dari", customer.Alamat)
}

func main() {
	asep := Customer{Nama: "Asep", Alamat: "Indonesia"}

	// Normal
	sayHi("Beta", asep)

	//STRUCT FUNC
	asep.sayHi("Beta")
}
