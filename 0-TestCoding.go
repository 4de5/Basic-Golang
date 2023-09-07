// package main

// import "fmt"

// func GetNilaiTambah(array []int, target int) (int, int) {
// 	for nilai1 := 0; nilai1 < len(array); nilai1++ {
// 		for nilai2 := 0; nilai2 < len(array); nilai2++ {
// 			switch {
// 			case array[nilai1]+array[nilai2] == target && array[nilai1] == array[nilai2]:
// 				return -1, -1
// 			case array[nilai1]+array[nilai2] == target:
// 				return array[nilai1], array[nilai2]
// 			default:
// 				return -1, -1 // Jika semua hasil false
// 			}
// 			// if array[nilai1]+array[nilai2] == target && array[nilai1]==array[nilai2]{
// 			// 	return -1, -1
// 			// } else {
// 			// 	array[nilai1]+array[nilai2] == target{
// 			// 	return array[nilai1], array[nilai2]
// 			// 	}
// 			// }
// 		}
// 	}
// 	return -1, -1 // Jika semua hasil false
// }

// func main() {
// 	iniSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	iniTarget := 8

// 	angka1, angka2 := GetNilaiTambah(iniSlice, iniTarget)
// 	fmt.Println(angka1, angka2)

// 	iniSlice[angka1-1] = 11
// 	iniSlice[angka2-1] = 11

// 	angka3, angka4 := GetNilaiTambah(iniSlice, iniTarget)
// 	fmt.Println(angka3, angka4)

// 	iniSlice[angka3-1] = 11
// 	iniSlice[angka4-1] = 11

// 	angka5, angka6 := GetNilaiTambah(iniSlice, iniTarget)
// 	fmt.Println(angka5, angka6)

// 	iniSlice[angka5-1] = 11
// 	iniSlice[angka6-1] = 11

//		angka7, angka8 := GetNilaiTambah(iniSlice, iniTarget)
//		fmt.Println(angka7, angka8)
//	}
//
// //////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	// Balikan Input Cara 1
	// xStr := strconv.Itoa(x) // Jadi string
	// balikanStr := ""
	// for i := len(xStr) - 1; i >= 0; i-- {
	// 	balikanStr += string(xStr[i])
	// }
	// balikanInt, _ := strconv.Atoi(balikanStr)
	// if x == balikanInt {
	// 	return true
	// } else {
	// 	return false
	// }

	// Cara Jenius
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	beginn := 0
	endd := len(str) - 1
	for beginn < endd {
		if !(str[beginn] == str[endd]) {
			return false
		}
		beginn++
		endd--
	}
	return true
}

func main() {
	x := 121
	fmt.Println(isPalindrome(x))
}
