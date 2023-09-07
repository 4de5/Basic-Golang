// untuk menggunakan fungsi di dalam fungsi itu sendiri
// sama seperti bahasa JAVA, juga ada untuk factorial

package main

import "fmt"

// Dengan Recursive
func faktorial(nilai int) int {
	if nilai == 1 {
		return 1
	} else {
		return nilai * faktorial(nilai-1)
	}
}

//  Dengan Loop
func perkalian(nilai int) int {
	total := 1
	for i := nilai; i > 0; i-- {
		total *= i
	}
	return total
}

func main() {
	fmt.Println(faktorial(3))
	fmt.Println(perkalian(3))

}
