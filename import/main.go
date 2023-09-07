package main

import (
	"belajar-go-with-pzn/helper"
	_ "belajar-go-with-pzn/helper"
	"fmt"
	_ "fmt"
)

func main() {
	helper.Jalankan()              // hanya run init karena otomatis
	fmt.Println(helper.Jalankan()) // run init&run Jalankan

	// Apa jadinya, bila kita ingin impor tapi ingin menjalankan init
	// Padahal secara default, deklarasi harus di gunakan
	// CARANYA!!!
	/*
		import (
			_ "belajar-go-with-pzn/helper"
	*/
	// Dengan <_> saat import kita tak perlu pemanggilan...
	// ...usahakan ada fmt biar di cetak(nampak jelas)
}
