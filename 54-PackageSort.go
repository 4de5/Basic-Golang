/*
Sort = meng-Urut-kan

agar data bisa diurut kita perlu, mengimplementasi kontrak interface...
... di sort.Interface

https://golang.org/pkg/sort/

// ISI sort.Interface
type Interface interface {
	Len() int           // Mengembalikan jumlah data
	Less(i, j int) bool // Cek apakah index i < index j
	Swap(i, j int)      // Menukar posisi i dan j
}

*/

package main

import (
	"fmt"
	"sort"
)

type User struct {
	Nama string
	Umur int
}

type UserSlice []User

func (user UserSlice) Len() int {
	return len(user)
}

func (user UserSlice) Less(i, j int) bool {
	return user[i].Umur < user[j].Umur
}

func (user UserSlice) Swap(i, j int) {
	user[i], user[j] = user[j], user[i]
}

func main() {
	data := []User{
		{"Asep", 19},
		{"Dwi", 25},
		{"Putra", 21},
	}
	fmt.Println(data)
	fmt.Println(UserSlice(data))
	sort.Sort(UserSlice(data)) // Perintah urutkan data
	fmt.Println(data)

	/////////////////////

	// Inisialisasi array string
	angka := []int{-1, 2, 1, 5, 3, 6, 4}
	names := []string{"John", "Mary", "Peter", "Jane"}
	names2 := []string{"John", "Mary", "Peter", "Jane"}

	// Urutkan array
	sort.Ints(angka)

	sort.Slice(names, func(a, b int) bool {
		return names[a] < names[b]
	})

	sort.Strings(names2)

	// Cetak array yang telah diurutkan
	fmt.Println(angka)
	fmt.Println(names)
	fmt.Println(names2)
}
