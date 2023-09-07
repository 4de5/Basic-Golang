//Map hampir sama dengan array ataupun slice, NAMUN
//index(key) di MAP bebas kita buat tidak harus 0-1-2-3...

package main

import (
	"fmt"
)

func main() {

	//CARA 1
	//Langsung di deklarasi
	var rumah map[string]string = map[string]string{
		//BISA DI SINGKAT JADI :
		//var rumah = map[string]string{...} atau
		//rumah := map[string]string{...}
		"pemilik": "Asep",
		"alamat":  "Indonesia",
	}
	fmt.Println(rumah)            //pemanggilan
	fmt.Println(rumah["pemilik"]) //pemanggilan key
	fmt.Println(rumah["alamat"])
	fmt.Println(len(rumah)) //melihat panjang length
	delete(rumah, "alamat") //meghapus key(alamat) dari map(rumah)
	fmt.Println(rumah)

	//CARA 2
	//Tidak Langsung
	var house map[string]string = make(map[string]string)
	//BISA DI SINGKAT JADI :
	//var house = make(map[string]string) atau
	//house := make(map[string]string)
	house["owner"] = "Asep" //tanpa koma, karena ini 1 statement
	house["address"] = "Indonesia"

	fmt.Println(house)
}
