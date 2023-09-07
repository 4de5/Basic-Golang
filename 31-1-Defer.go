/**sebuah program yang akan menjalankan sebuah fungsi-lain
apabila telah selesai/dipaksa selesai oleh eror

Ditulis --> <defer namaFungsi()>

*/

package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Selesai")
}

func runApp() {
	defer endApp() // Fungsi defer--> ...
	// ...akan menjalankan endApp saat runApp SELESAI/BERHENTI
	fmt.Println("Aplikasi Berjalan Baik")
	// Setelah di SINI(akhir), program memantul ke atas menjalankan <defer>
	// Bahkan walaupun eror di tengah(selesai paksa)
}

func runAppError(eror bool) {
	defer endApp() // Akan dijalankan apabila selesai/berhenti
	if eror {
		panic("Aplikasi Eror") // Dijelaskan di Next
	} // Eror di tengah run
	fmt.Println("Aplikasi Berjalan Baik") //Kode ini tidak sempat dijalankan
}

func main() {
	// SEKARANG FOKUS DISINI!!!!!!
	runApp()

	// STUDI KASUS KETIKA ADA EROR!!!!!!!
	// runAppError(true)
}
