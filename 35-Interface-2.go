//DENGAN INTERFACE

// Interface digunakan agar program lebih singkat
// Kata mas Rizky, interface digunakan agar resep restoran tidak dilihat secara publik

package main

import "fmt"

type Karyawan interface {
	GetGaji() int
}

type KaryawanTetap struct {
	Nama      string
	Upah      int
	Tunjangan int
}

type KaryawanKontrak struct {
	Nama string
	Upah int
}

type KaryawanLepas struct {
	Nama     string
	Rate     int
	JamKerja int
}

func (pt KaryawanTetap) GetGaji() int {
	return pt.Upah + pt.Tunjangan
}

func (pk KaryawanKontrak) GetGaji() int {
	return pk.Upah
}

func (pl KaryawanLepas) GetGaji() int {
	return pl.JamKerja * pl.Rate
}

func TotalGaji(gaji []Karyawan) int { // tipe slice sesuai interface
	total := 0
	for _, gajiPerOrang := range gaji {
		total = total + gajiPerOrang.GetGaji() // HARUS PANGGIL FUNC
	}
	return total
}

func main() {

	pt1 := KaryawanTetap{Nama: "A", Upah: 3_000_000, Tunjangan: 3_000_000} // 6_000_000
	pt2 := KaryawanTetap{Nama: "B", Upah: 3_000_000, Tunjangan: 3_000_000} // 6_000_000
	pk1 := KaryawanKontrak{Nama: "C", Upah: 3_000_000}                     // 3_000_000
	pk2 := KaryawanKontrak{Nama: "D", Upah: 2_000_000}                     // 2_000_000
	pl1 := KaryawanLepas{Nama: "E", Rate: 50_000, JamKerja: 21}            // 1_050_000
	pl2 := KaryawanLepas{Nama: "F", Rate: 250_000, JamKerja: 8}            // 2_000_000

	//Berapa pengeluaran gaji/upah perusahaan?

	Pengeluaran := TotalGaji([]Karyawan{pt1, pt2, pk1, pk2, pl1, pl2}) // Slice Interface

	fmt.Println(Pengeluaran) // 20_050_000

}
