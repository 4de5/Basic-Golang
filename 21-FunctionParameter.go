//Parameter --> Data yang diambil dari luar function

package main

import "fmt"

func main() {
	sapaan("Asep", "belajar") //proses memasukan parameter
	firstName := "Putra"
	sapaan(firstName, "belajar")
}

func sapaan(firstName string, lastName string) { //firstName, lastName == parameter
	fmt.Println("Hai", firstName, lastName)
}
