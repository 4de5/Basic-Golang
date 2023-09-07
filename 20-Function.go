//Function atau Method di beberapa bahasa pemrograman lainnya
//Caranya hampir sama dengan Java yang kita pelajari kemarin

package main

import "fmt"

func sayHello() { //ini Fungsi sayHello
	fmt.Println("Halo Mbakk")
}

func main() { //ini function MAIN

	//kita coba panggil fungsi sayHello
	sayHello()
	fmt.Println()
	//Kita coba dengan Loop
	for i := 0; i < 3; i++ {
		sayHello()
	}
}
