//variabel --> kotak pembungkus value
//1 variabel=1 tipe data, sama kaya C, Java
//1 nama value -->bila udah bikin <int value>, ga bisa bikin <string value>
//bisa diakses dimanapun?

package main

import "fmt"

func main() {

	//CARA 1
	var nama string //setelah dideklarasi variabel harus dipake, agar ga eror
	nama = "Asep Dwi Saputra"
	fmt.Println(nama)
	nama = "Menulis Perubahan 1 nama"
	fmt.Println(nama)

	//CARA2
	var name = "Asep Dwi Saputra"
	fmt.Println(name)
	name = "Menulis Perubahan 1 name"
	fmt.Println(name)

	//CARA 3
	aran := "Asep Dwi Saputra"
	fmt.Println(aran)
	aran = "Menulis Perubahan 1 aran"
	fmt.Println(aran)

	//Multiple Var
	var (
		firstName  = "Asep"
		secondName = "Putra"
	)
	fmt.Println(firstName)
	fmt.Println(secondName)

}
