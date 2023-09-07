/**sebuah program yang akan menghentikan runing function dan bisa memberikan message eror

Ditulis --> <pamic("bisa gini")> atau
		--> <pamic(1)> atau
		--> <pamic(true)>

*/

package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Selesai")
}

func runAppError(eror bool) {
	defer endApp()
	if eror {
		panic("Apliaksi Eror")
	} // Eror di tengah run
	fmt.Println("Aplikasi Berjalan Baik") //Kode ini tidak sempat dijalankan
}

func main() {

	// STUDI KASUS KETIKA ADA EROR!!!!!!!
	runAppError(true)
}
