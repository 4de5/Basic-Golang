/**
Kemampuan merubah tipe data

Sering digunakan bila ingin mengubah data dari interface(abstrak)...
...menjadi nilai tipedata(int, string, bool)
*/

package main

import (
	"fmt"
)

func Es() interface{} {
	return "Es Milo"
}

func main() {
	result := Es()
	//Cara 1
	var value string = result.(string) //Bila salah konfersi=panic
	fmt.Println(value)

	//Cara 2
	switch result.(type) {
	case string:
		fmt.Println("Result", result, "is String")
	case int:
		fmt.Println("Result", result, "is Int")
	default:
		fmt.Println("Result", result, "is Boolean")
	}

}
