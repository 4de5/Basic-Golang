/**sebuah program yang akan menangkap pesan panic(eror) dan ...
...menjalankan program setelah eror

recover bernilai == pesan eror di program panic

BAIKNYA ditulis di dalam fungsi <defer> karena defer dijalankan di akhir.
Kalo, program recover tidak aktif karena dibawah panic, yasalammm

Ditulis --> <recover()> atau
		BISA MEMBUNGKUS RECOVER DALAM VAR
		contoh --> <message := recover()
					fmt.Println("Pesan eror =", message)>
*/

package main

import "fmt"

func endApp() {
	recover()
	// Ditangkap dan program dibawahnya fungsi pemanggilan akan dijalankan

	//Opsi NIL --> dibahas di #37
	// message := recover()
	// if message != nil {
	// 	fmt.Println("Penanda EROR dengan pesan", message)
	// }

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
	fmt.Println("Ini tidak berjalan ketika Eror tidak di Recover")
}
