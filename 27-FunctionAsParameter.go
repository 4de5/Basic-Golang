//Menggunakan fungsi as value menjadi parameter di fungsi

package main

import "fmt"

func filter(name string) string {

	switch name {
	case "Anjing":
		return "..."
	case "Babi":
		return "..."
	case "Goblok":
		return "..."
	default:
		return name
	}

}

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello ", filter(name))
} // <filter func(string) string> bisa disingkat

//CARA MENYINGKAT PARAMETER FUNGSI dengan <type>
// Seperti pada modul 9 kita bisa menginisialisasi
// Baiknya ditulis sebelum fungsi agar ga ribet cari atau dibawah import
type Filter func(string) string

func sayHelloWithFilter2(name string, filter Filter) {
	fmt.Println("Hai ", filter(name))
}

func main() {
	sayHelloWithFilter("Babi", filter)
	sayHelloWithFilter2("Anjingku", filter)

}
