//TANPA INTERFACE

// Interface digunakan agar program lebih singkat
// Kata mas Rizky, interface digunakan agar resep restoran tidak dilihat secara publik

package main

import "fmt"

type PegawaiTetap struct {
	Nama      string
	Upah      int
	Tunjangan int
}

type PegawaiKontrak struct {
	Nama string
	Upah int
}

type PegawaiLepas struct {
	Nama     string
	Rate     int
	JamKerja int
}

func (pt PegawaiTetap) GetGaji() int {
	return pt.Upah + pt.Tunjangan
}

func (pk PegawaiKontrak) GetGaji() int {
	return pk.Upah
}

func (pl PegawaiLepas) GetGaji() int {
	return pl.JamKerja * pl.Rate
}

func main() {

	pt1 := PegawaiTetap{Nama: "A", Upah: 3_000_000, Tunjangan: 3_000_000} // 6_000_000
	pt2 := PegawaiTetap{Nama: "B", Upah: 3_000_000, Tunjangan: 3_000_000} // 6_000_000
	pk1 := PegawaiKontrak{Nama: "C", Upah: 3_000_000}                     // 3_000_000
	pk2 := PegawaiKontrak{Nama: "D", Upah: 2_000_000}                     // 2_000_000
	pl1 := PegawaiLepas{Nama: "E", Rate: 50_000, JamKerja: 21}            // 1_050_000
	pl2 := PegawaiLepas{Nama: "F", Rate: 250_000, JamKerja: 8}            // 2_000_000

	//Berapa pengeluaran gaji/upah perusahaan?
	var Pengeluaran int

	Pengeluaran += pt1.GetGaji()
	Pengeluaran += pt2.GetGaji()
	Pengeluaran += pk1.GetGaji()
	Pengeluaran += pk2.GetGaji()
	Pengeluaran += pl1.GetGaji()
	Pengeluaran += pl2.GetGaji()

	fmt.Println(Pengeluaran) // 20_050_000

}
