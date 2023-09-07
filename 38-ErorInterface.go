// Interface untuk kontrak membuat eror
//Sudah ada library jadi ga perlu buat manual

package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	// Kembalian int(untuk data benar), eror(untuk data eror)
	if pembagi == 0 {
		return 0, errors.New("Pembagi dilarang 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	// RUN
	hasil, err := Pembagian(2, 0)
	//hasil, err := Pembagian(2, 1)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Eror", err.Error())
	}
}
