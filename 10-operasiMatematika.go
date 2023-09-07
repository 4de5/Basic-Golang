//Operasi sama dengan bahasa Java, JS, dan lainnya

package main

import "fmt"

func main() {

	//Operasi  Dasar (+ - * / %)
	var result = 10 + 10
	var a = 10
	var b = 10
	var c = a % b
	fmt.Println(result)
	fmt.Println(c)

	//Augmanted Assignment
	var angka = 1
	angka += 1
	fmt.Println(angka)

	//Unary Operator (++, --, -, +, !)
	//operasi yang dilakukan pada 1 variabel
	var nilai = true
	fmt.Println(!nilai)
	var x int = 1
	x--
	fmt.Println(x)

}
