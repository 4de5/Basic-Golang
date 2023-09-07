// Fungsi tak bernama
// Alih-alih membuat func secara manual, biasanya anonyFunc digunakan
// biasanya juga langsung di bungkus dengan var

package main

import "fmt"

type Blacklist func(name string) bool // AnonymusFunc

func daftarApp(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Kamu di block", name)
	} else {
		fmt.Println("Selamat Datang", name)
	}
}

func main() {

	// Masukkan dulu daftar blokir, isi AnonyFunc
	blacklist := func(name string) bool {
		return name == "Anjing" // kembalian bool sesuai func
	}

	daftarApp("Anjing", blacklist)

}
