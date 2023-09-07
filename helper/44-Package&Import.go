package helper

import "fmt"

func SayHello(nama string) {
	fmt.Println("Hello", nama)
}

/*
	Dalam 1 folder package baiknya berisi 1 fungsi
	yang tidak ada kesamaan dengan file lain,
	agar mesin tidak bingung dan menghindari eror
*/
