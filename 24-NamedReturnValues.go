//memeberikan  nama  pada return value
//biasanya tidak ada di bahasa pemrogaman  lalinnya

package main

import "fmt"

func nama() (panggilan, tengah, marga string /*tipe data bisa ditulis gini, di ahir jika emangsama semua, kalo ngga yang sisa sama di akhir*/) {

	panggilan = "Asep"
	tengah = "Dwi"
	marga = "Saputra"

	return panggilan, tengah, marga //bisa ditulis <return> aja

}
func alamat() (noRumah byte, desa, kecamatan string) {

	noRumah = 10
	desa = "Bojanegara"
	kecamatan = "Padamara"

	return

}

func main() {
	fmt.Println("Nama")
	fmt.Println(nama())
	fmt.Println(alamat())

}
