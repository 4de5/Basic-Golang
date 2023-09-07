//SLICE adalah bagian dari array
//Data slice bisa diubah
//terdiri dari --> Pointer(petunjuk data awal), Length(panjang), Capacity(Kapasitas)

package main

import "fmt"

func main() {

	//Misalnya Terdapat array hari-hari
	var days = [...]string{
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "mingggu",
	}
	fmt.Println(days)

	//KITA AKAN MEMBUAT SLICE!!

	//Cara 1
	var weekday = days[0:5] //slice weekday diambil dari array days index 0 s/d (5-1)
	//bisa saja ditulis [:5]
	fmt.Println(weekday)
	fmt.Println(len(weekday)) //length
	fmt.Println(cap(weekday)) //Capacity

	//Cara 2
	weekend := days[5:] //dari index 5 s/d terbesar
	//dalam hal ini seluruh capacity telah terpakai
	fmt.Println(weekend)
	fmt.Println(len(weekend)) //length
	fmt.Println(cap(weekend)) //Capacity

	//dari 2 cara diatas didapat pola
	//POLA PEMBUATAN SLICE DARI ARRAY!!!!!!!!
	//array[low:high] --> diambil dari array(array), index low(yg dipilih) s/d index(high-1)
	//array[:high] --> diambil dari array(array), index terkecil(0) s/d index(high - 1)
	//array[low:] --> diambil dari array(array), index low(yg dipilih) s/d index terakhir/terbesar
	//array[:] --> diambil dari array(array), semua index

	//Cara 3 --> membuat sliceBARU(dari lama) dan menambah dan/atau rubah data baru di belakang
	lemburDay := append(weekend, "cutiDay") //membuat array baru, karena capacity weekend sudah penuh
	lemburDay[0] = "sabtuLembur"            //merubah data dalam slice lemburDay index[0]
	fmt.Println(lemburDay)
	fmt.Println(weekend) //data di slice(weekend) tidak ikut berubah karena lemburDay membuat array baru
	fmt.Println(days)    //data di array tidak berubah karena data slice weekEnd tidak berubah

	//Cara 4 --> membuat slice tanpa membuat array dulu, disini array disembunyikan
	//namaArray := make([]tipeData, totalLength, totalCapacity)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Asep"
	newSlice[1] = "Putra"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//COPY SLICE
	//BUAT SLICE TANPA ARRAY DULU, Total length - Total Capacity DIPERHATIKAN!
	//copy(namaTujuan, namaSumber)
	namaKu := make([]string, 2, 5) //length-capacity SAMA, agar ga kepotong/eror
	copy(namaKu, newSlice)         //copy slice newSlice ke slice kosong(namaKu)
	fmt.Println(namaKu)

	//HATI-HATI!!
	iniArray := [...]int{1, 2, 3}
	iniJugaArray := [3]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}

	fmt.Println(iniArray)
	fmt.Println(iniJugaArray)
	fmt.Println(iniSlice)

}
