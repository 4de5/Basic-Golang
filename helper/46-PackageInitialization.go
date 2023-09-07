/*
Digunakan bila kita ingin menjalankan suatu program secara otimatis
ketika Package-nya di impor

caranya --> buat <func init(){...}>
*/
package helper

import "fmt"

var run string

func init() {
	fmt.Println("Sukses! Mengimpor data helper | edit:46-...")
	run = "File 46 dijalankan"
}

func Jalankan() string {
	return run
}
