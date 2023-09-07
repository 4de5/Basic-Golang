/**
Di bahasa Java kalo tipe data non primitif tidak ada input nilai maka = Null
Nil = Null
Di GO primitif semua --> langsung ada nilainya ...
... int --> 0
... bool --> false
... string --> "<kosong>"
# Nil bisa digunakan untuk:
	1. interface
	2. function
	3. map
	4. slice
	5. pointer
	6. channel
*/

package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil // Bernilai Kosong
	} else {
		return map[string]string{
			"Name": name,
		}
	}
}
func main() {
	newMap := NewMap("")
	newMap2 := NewMap("Oke")
	fmt.Println(newMap)
	fmt.Println(newMap2)

}
