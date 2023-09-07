/**Karena Golang bukan bahasa OOP maka tidak ada data puncak(Parents/Superclass)...
...contoh super class di Java-->java.lang.Object
Maka dari itu, di GO ada interface kosong
Interface0 --> interface yang tidak ada deklarasi fungsi/method, ...
... artinya semua tipe data adalah tipe data itu secara fungsi/implementasi
Biasa digunakan untuk proses yang ingin mengacuhkan tipe data

contoh:
	fmt.Println(a ...interface{})
	panic(v interface{})
	recover() interface{}
	dan lain-lain

*/

package main

import "fmt"

func Data(i int) interface{} {
	if i%2 == 0 {
		return false
	} else {
		return "Ganjil"
	}
}

func main() {
	//var data int = Data(1) //EROR = func data harus tak bertipe data jika dibungkus
	var data interface{} = Data(12)

	fmt.Println(data)
	fmt.Println(Data(1))
}
